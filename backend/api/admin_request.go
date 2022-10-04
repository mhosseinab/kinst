package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/r3labs/diff"
	"github.com/spf13/cast"
	"connections"
	"es"
	"models"
)

func adminUpdateRequest(c *gin.Context) {
	id := c.Param("id")

	var m models.Request
	if err := db.First(&m, "id=?", id).Error; err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	beforeChange := m

	curStatus := m.Status
	curExpertStatus := m.ExpertStatus

	if err := c.ShouldBind(&m); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error_msg": err.Error()})
		return
	}

	var autoAcceptMaxAmount int64 = 15_000_000
	if m.Province == "خوزستان" {
		autoAcceptMaxAmount = 50_000_000
	}

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch {
			if curStatus != m.Status {
				if m.ExpertAcceptedAmount > autoAcceptMaxAmount && m.Status == models.RequestStatusPayed {
					c.JSON(400, gin.H{"error_msg": "برای این مبالغ امکان تغییر وضعیت توسط کاربران شعب وجود ندارد."})
					return
				}
			}
		}
	}

	if curStatus != m.Status {
		if m.Status == models.RequestStatusReadyToPay && m.Sheba == "" {
			c.JSON(400, gin.H{"error_msg": "باید شبا وارد شده باشد"})
			return
		}

		if !(m.DamageType == models.RequestDamageTypeDeath && m.Status == models.RequestStatusPayed) {
			sendSmsChaneStatus(m.MobileNumber)
		}

	}

	// expert status changed
	if curExpertStatus != m.ExpertStatus {
		m.ExpertUpdatedAt = time.Now()
		if m.ExpertStatus == models.RequestExpertStatusAccepted {
			if m.DamageType == models.RequestDamageTypeInstrument {
				if m.ExpertAcceptedAmount <= autoAcceptMaxAmount {
					m.Status = models.RequestStatusReadyToPay
					m.AcceptedAmount = m.ExpertAcceptedAmount
				}
			}
		}
	}

	if err := db.Save(&m).Error; err != nil {
		log.Println(err.Error())
	}

	changelog, _ := diff.Diff(beforeChange, m)
	if changelog != nil && len(changelog) > 1 {
		changelogByte, _ := json.Marshal(changelog)
		cl := models.RequestChangelog{
			Changelogs: string(changelogByte),
			CreatedAt:  time.Now(),
			RequestID:  m.ID,
		}
		cl.UserID = getUserFromSession(c).UserID
		db.Create(&cl)
	}

	if client, err := connections.GetElasticsearch(); err == nil {
		es.StoreRequestItem(client, m)
	}

	c.JSON(http.StatusOK, m)
}

// adminRequestItems returns request items
func adminRequestItems(c *gin.Context) {
	var reqs models.RequestSlice
	var count int
	offset, limit, order := getRequestFilters(c)
	d := db.Model(&models.Request{}).Where("reference_code!=?", "")

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")
			d = d.Where("company_id in (?)", p)
		}
	}

	filterKeys := []string{
		"status",
		"province",
		"city",
		"sum_damage_amount",
		"casuality_date",
		"id",
		"damage_type",
		"national_code",
		"reference_code",
		"bill_identifier",
		"expert_status",
	}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)
		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	if q, ok := c.Request.URL.Query()["query"]; ok {
		if isValidRawQuery(q[0]) {
			d = d.Where(q[0])
		}
	}

	err := d.Count(&count).
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&reqs).Error

	c.JSON(http.StatusOK, gin.H{
		"objects": reqs,
		"meta": gin.H{
			"total_count": count,
			"err":         err,
		},
	})
}

func isValidRawQuery(q string) bool {
	log.Println(q)
	re, _ := regexp.Compile("^sum_damage_amount = (\\d+)$")
	if re.MatchString(q) {
		return true
	}

	re, _ = regexp.Compile("^sum_damage_amount BETWEEN (\\d+) AND (\\d+)$")
	if re.MatchString(q) {
		return true
	}
	log.Println(q)
	log.Println("not a valid advanced query")
	return false
}

func getMysqlQueryFromQueryString(q string, d *gorm.DB) string {
	f := strings.Split(q, "+")
	field := f[0]
	operator := f[1]

	qi := ""

	switch field {
	case "sum_damage_amount":
		switch operator {
		case "=", "%3D":
			qi = fmt.Sprintf("%s = ?", field)
			if d != nil {
				d.Where(qi, f[2])
			}
			return qi
		case "BETWEEN":
			qi = fmt.Sprintf("%s BETWEEN ? AND ?", field)
			if d != nil {
				d.Where(qi, f[2], f[4])
			}
			return qi
		}

	}

	return qi
}

// adminRequestChangelog returns request items with changelog
func adminRequestChangelog(c *gin.Context) {
	var reqs models.RequestChangelogSlice
	var count int
	offset, limit, order := getRequestFilters(c)
	d := db.Model(&models.RequestChangelog{})

	filterKeys := []string{"request_id"}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)
		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	err := d.Count(&count).
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&reqs).Error

	for index := range reqs {
		var cd diff.Changelog
		err := json.Unmarshal([]byte(reqs[index].Changelogs), &cd)
		if err != nil {
			log.Println(err.Error())
		}
		reqs[index].ChangelogData = cd
	}

	c.JSON(http.StatusOK, gin.H{
		"objects": reqs,
		"meta": gin.H{
			"total_count": count,
			"err":         err,
		},
	})
}

// adminRequestItem returns request item
func adminRequestItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Request

	d := db.Model(&models.Request{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")
			d = d.Where("company_id in (?)", p)
		}
	}

	if err := d.
		Find(&item, "id=?", id).Error; err != nil {
		log.Println(err.Error())
		c.JSON(404, "404 page not found")
		return
	}

	c.JSON(http.StatusOK, item)
}

func getRequestFilters(c *gin.Context) (offset int, limit int, sort string) {
	reqVals := c.Request.URL.Query()
	offset = 0
	limit = 10
	sortField := "id"
	sortOrder := "desc"

	if start := reqVals.Get("page"); start != "" {
		offset = cast.ToInt(start)
		if end := reqVals.Get("limit"); end != "" {
			limit = cast.ToInt(end) - cast.ToInt(start)
		}
	}
	if sortFieldReq := reqVals.Get("sortField"); sortFieldReq != "" {
		sortField = sortFieldReq
		if sortOrderReq := reqVals.Get("sortOrder"); sortOrderReq != "" {
			switch sortOrderReq {
			case "descend":
				sortOrder = "DESC"
			case "ascend":
				sortOrder = "ASC"
			}

		}
	}
	return offset, limit, sortField + " " + sortOrder
}
