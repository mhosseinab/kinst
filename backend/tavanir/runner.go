package tavanir

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/r3labs/diff"

	"github.com/spf13/cast"
	ptime "github.com/yaa110/go-persian-calendar"
	"models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDb(d *gorm.DB) {
	db = d
}

func handleMessage(m Message) error {
	switch m.Message {

	case StatusMessageAmountCompliant,
		StatusMessageCheckoutCompliant,
		StatusMessageTavanirAmountCompliant,
		StatusMessageTavanirCoverCompliant:

		var cnt int
		db.Model(&Case{}).Where(Case{TavanirID: m.CaseID}).Count(&cnt)
		if cnt > 0 {
			// updateCase(c)
		} else {

			db.Model(&StatusUpdateQueue{}).Save(&StatusUpdateQueue{
				CaseID:    cast.ToUint64(m.CaseID),
				NewStatus: m.Message,
				RefID:     cast.ToUint64(m.RefID),
			})
		}
		break

	case StatusMessageAssigned:
		/*
			var cnt int
			db.Model(&Case{}).Where(Case{TavanirID: m.CaseID}).Count(&cnt)
			if cnt > 0 {
				// updateCase(c)
			} else {
				c, err := GetCaseString(m.CaseID)
				if err != nil {
					log.Println(err.Error())
					return err
				}
				c.Status = CaseStatusNew
				saveCase(c)
			}
		*/
		break

	case StatusMessageDocumentUpdate:
		var c Case
		var cnt int
		db.Model(&Case{}).Where(Case{TavanirID: m.CaseID}).Find(&c).Count(&cnt)
		if cnt == 0 {
			saveCase(c)
		}
		db.Model(&StatusUpdateQueue{}).Save(&StatusUpdateQueue{
			CaseID:    cast.ToUint64(m.CaseID),
			NewStatus: MethodGetDocumentUpdate,
			RefID:     cast.ToUint64(m.RefID),
		})
		break
	case StatusMessageUpdateBankInfo:
		var c Case
		var cnt int
		db.Model(&Case{}).Where(Case{TavanirID: m.CaseID}).Find(&c).Count(&cnt)
		if cnt == 0 {
			saveCase(c)
		}

		db.Model(&StatusUpdateQueue{}).Save(&StatusUpdateQueue{
			CaseID:    cast.ToUint64(m.CaseID),
			NewStatus: MethodGetUpdateBankInfo,
			RefID:     cast.ToUint64(m.RefID),
		})
		break

	default:
		return fmt.Errorf("unhandled %s", m.Message)
	}

	return nil
}

func syncWithTavanir() {
	for {
		var tasks StatusUpdateQueueSlice
		db.Where("is_done = 0").Find(&tasks)
		for _, t := range tasks {
			var c Case
			db.Model(&Case{}).Where(Case{TavanirID: cast.ToString(t.CaseID)}).Find(&c)
			switch t.NewStatus {
			case models.RequestStatusIncomplete:
				// Missing Document 2000 is INTERNAL related to SHEBA
				c.MissingDocuments = strings.Replace(c.MissingDocuments, "2000,", "", -1)
				c.MissingDocuments = strings.Replace(c.MissingDocuments, ",2000", "", -1)
				refID, err := SetDocumentRequired(
					cast.ToUint64(c.TavanirID),
					cast.ToIntSlice(strings.Split(c.MissingDocuments, ",")),
					c.ExpertDescription)
				updateStatusQueue(t, err, refID)
				break
			case models.RequestStatusRejected:
				refID, err := SetNotCovered(cast.ToUint64(c.TavanirID), c.NotCoveredReason, c.ExpertDescription)
				updateStatusQueue(t, err, refID)
				break
			case models.RequestStatusNeedToVisitInPerson:
				refID, err := SetPhysicalRequired(cast.ToUint64(c.TavanirID), c.ExpertDescription)
				updateStatusQueue(t, err, refID)
				break
			case MethodPaymentInfoUpdateRequest:
				refID, err := PaymentInfoUpdateRequest(cast.ToUint64(c.TavanirID), c.ExpertDescription)
				updateStatusQueue(t, err, refID)
				break
			case models.RequestStatusReadyToPay:
				if c.AcceptedAmount != 0 {
					refID, err := SetCovered(cast.ToUint64(c.TavanirID), c.ExpertDescription)
					//fmt.Println("SetCovered", refID, err)
					if err != nil {
						updateStatusQueue(t, err, refID)
						continue
					}
					refID, err = SetAmountEstimate(cast.ToUint64(c.TavanirID), cast.ToUint64(c.AcceptedAmount), c.ExpertDescription)
					updateStatusQueue(t, err, refID)
				} else {
					t.IsDone = true
					updateStatusQueue(t, fmt.Errorf("invalid AcceptedAmount"), 0)
				}
				break
			case models.RequestStatusPayed:
				if c.AcceptedAmount != 0 {
					pt := ptime.Now(ptime.Iran())
					s := strings.Split(c.Sheba, " ")
					refID, err := Checkout(cast.ToUint64(c.TavanirID), c.ExpertDescription, uint64(c.AcceptedAmount),
						pt.Format("yyyy-MM-dd HH:MM"), "000000", strings.Join(s[1:], " "), s[0])
					updateStatusQueue(t, err, refID)
					c.Status = models.RequestStatusPayed
					c.ExpertStatus = models.RequestStatusPayed
					db.Model(Case{}).Where("tavanir_id = ?", c.TavanirID).Update(&c)
				} else {
					t.IsDone = true
					updateStatusQueue(t, fmt.Errorf("invalid AcceptedAmount"), 0)
				}
				break
			case MethodGetAmountCompliant:
				b, err := GetAmountCompliant(cast.ToUint64(c.TavanirID), cast.ToUint64(t.RefID))
				if err != nil {
					updateStatusQueue(t, err, t.RefID)
					continue
				}

				beforeChange := c
				c.Status = models.RequestCompliant
				c.Descr = b
				c.ExpertStatus = models.RequestStatusIncompleteChange

				db.Model(Case{}).Where("tavanir_id = ?", t.CaseID).Update(&c)
				updateStatusQueue(t, err, t.RefID)
				saveChange(beforeChange, c)
				break
			case MethodGetCheckoutCompliant:
				b, err := GetCheckoutCompliant(cast.ToUint64(c.TavanirID), cast.ToUint64(t.RefID))
				if err != nil {
					updateStatusQueue(t, err, t.RefID)
					continue
				}
				beforeChange := c
				c.Status = models.RequestCompliant
				c.Descr = b
				c.ExpertStatus = models.RequestStatusIncompleteChange

				db.Model(Case{}).Where("tavanir_id = ?", t.CaseID).Update(&c)
				updateStatusQueue(t, err, t.RefID)
				saveChange(beforeChange, c)
				break
			case MethodGetTavanirAmountCompliant:
				b, err := GetTavanirAmountCompliant(cast.ToUint64(c.TavanirID), cast.ToUint64(t.RefID))
				if err != nil {
					updateStatusQueue(t, err, t.RefID)
					continue
				}
				beforeChange := c
				c.Status = models.RequestCompliant
				c.Descr = b
				c.ExpertStatus = models.RequestStatusIncompleteChange

				db.Model(Case{}).Where("tavanir_id = ?", t.CaseID).Update(&c)
				updateStatusQueue(t, err, t.RefID)
				saveChange(beforeChange, c)
				break
			case MethodGetTavanirCoverCompliant:
				b, err := GetTavanirCoverCompliant(cast.ToUint64(c.TavanirID), cast.ToUint64(t.RefID))
				if err != nil {
					updateStatusQueue(t, err, t.RefID)
					continue
				}
				beforeChange := c
				c.Status = models.RequestCompliant
				c.Descr = b
				c.ExpertStatus = models.RequestStatusIncompleteChange

				db.Model(Case{}).Where("tavanir_id = ?", t.CaseID).Update(&c)
				updateStatusQueue(t, err, t.RefID)
				saveChange(beforeChange, c)
				break
			case MethodGetUpdateBankInfo:
				b, err := GetUpdateBankInfo(cast.ToUint64(c.TavanirID), cast.ToUint64(t.RefID))
				if err != nil {
					updateStatusQueue(t, err, t.RefID)
					continue
				}
				beforeChange := c
				c.Sheba = fmt.Sprintf("%s %s %s", b.AccountShebaNumber, b.AccountBankName, b.AccountOwnerName)
				c.ExpertStatus = models.RequestStatusIncompleteChange

				db.Model(Case{}).Where("tavanir_id = ?", t.CaseID).Update(&c)
				updateStatusQueue(t, err, t.RefID)
				saveChange(beforeChange, c)
				break
			case MethodGetDocumentUpdate:
				docs, err := GetDocumentUpdate(cast.ToUint64(t.CaseID), cast.ToUint64(t.RefID))
				if err != nil {
					updateStatusQueue(t, err, t.RefID)
					continue
				}
				for _, d := range docs {
					d.CaseId = c.Id
					count := 0
					db.Model(&Document{}).
						Where(Document{DocumentTypeID: d.DocumentTypeID, FileName: d.FileName}).
						Count(&count)
					if count == 0 {
						db.Save(&d)
					}
				}

				c.ExpertStatus = models.RequestStatusIncompleteChange

				db.Model(Case{}).Where("tavanir_id = ?", t.CaseID).Update(&c)
				updateStatusQueue(t, err, t.RefID)
				break
			case FetchFailedMessage:
				//log.Println("[ i ] " + FetchFailedMessage)
				var m Message
				db.Model(&Message{}).
					Where(Message{MessageID: cast.ToString(t.RefID)}).
					First(&m)
				//log.Println(m)
				err := handleMessage(m)
				updateStatusQueue(t, err, t.RefID)
				break
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func updateStatusQueue(t StatusUpdateQueue, err error, refID uint64) {
	if err != nil {
		//log.Println("[!] Task Failed")
		//log.Println(err)
		t.Note = err.Error()
		if strings.Contains(strings.ToLower(err.Error()), "invalid state") {
			t.IsDone = true
		}
		db.Model(&StatusUpdateQueue{}).Update(&t)
		return
	}

	t.RefID = refID
	t.IsDone = true
	t.Success = true
	db.Model(&StatusUpdateQueue{}).Update(&t)
}

func updateCase(c Case) {
	docs := c.Documents
	c.Documents = []Document{}

	// fix arabic characters
	c.StateName = FixArabicCharacters(c.StateName)
	c.CityName = FixArabicCharacters(c.CityName)

	db.Model(&Case{}).Where(Case{TavanirID: c.TavanirID}).Update(&c)

	for _, d := range docs {
		d.CaseId = c.Id
		count := 0
		db.Model(&Document{}).
			Where(Document{DocumentTypeID: d.DocumentTypeID, FileName: d.FileName}).
			Count(&count)
		if count == 0 {
			db.Save(&d)
		}
		//db.FirstOrCreate(&d, Document{DocumentTypeID: d.DocumentTypeID, FileName: d.FileName})
	}
}

func saveCase(c Case) {
	docs := c.Documents
	c.Documents = []Document{}

	// fix arabic characters
	c.StateName = FixArabicCharacters(c.StateName)
	c.CityName = FixArabicCharacters(c.CityName)

	db.Model(&Case{}).Save(&c)

	for _, d := range docs {
		d.CaseId = c.Id
		count := 0
		db.Model(&Document{}).
			Where(Document{DocumentTypeID: d.DocumentTypeID, FileName: d.FileName}).
			Count(&count)
		if count == 0 {
			db.Save(&d)
		}
	}
}

func saveChange(before, current Case) {
	changelog, _ := diff.Diff(before, current)
	if changelog != nil && len(changelog) > 1 {
		changelogByte, _ := json.Marshal(changelog)
		cl := CaseChangelog{
			ChangeLogs: string(changelogByte),
			CreatedAt:  time.Now(),
			CaseID:     cast.ToUint(current.Id),
		}
		cl.UserID = 0
		db.Create(&cl)
	}
}

func syncTavanirStatus() {
	var cases CaseSlice
	db.Select("tavanir_id").Find(&cases)
	for _, c := range cases {
		fmt.Println("c.TavanirID: ", c.TavanirID)
		s, err := getCaseStatus(cast.ToUint64(c.TavanirID))
		if err != nil {
			log.Println("[!!] getCaseStatus failed: " + err.Error())
			continue
		}

		if val, ok := tavanirStatusText[s]; ok {
			db.Model(&Case{}).Where("tavanir_id = ?", c.TavanirID).Update("status", val)
		}

	}
}

func fixArabicCharacters() {
	var cases CaseSlice
	db.Find(&cases)
	for _, c := range cases {
		//log.Println("case: " + c.TavanirID)
		db.Model(&Case{}).
			Where("tavanir_id = ?", c.TavanirID).
			Updates(Case{StateName: FixArabicCharacters(c.StateName), CityName: FixArabicCharacters(c.CityName)})
	}
}

func fixTavanirExpertStatus() {
	db.Model(&Case{}).Where("status = ?", models.RequestStatusPayed).Update("expert_status", models.RequestStatusPayed)
}

func fetchMessages() {
	var start uint64 = 1
	for {
		fmt.Println("Getting messages")
		messages, err := GetNextMessages(start, justNotSeenTypeAll, recordsCountType100)
		if err != nil {
			//log.Println("[!] Fetch ERROR")
			log.Println(err.Error())
		}
		for _, m := range messages {
			//spew.Dump(m)
			start = cast.ToUint64(m.MessageID)
			count := 0
			db.Model(&Message{}).Where(Message{MessageID: m.MessageID}).Count(&count)
			if count != 0 {
				// make sure case has already been registered
				// This should fix missing cases in the past where we didn't check message handle status
				db.Model(&Case{}).Where(Case{TavanirID: m.CaseID}).Count(&count)
				if count != 0 {
					//log.Println("[!] already processed")
					continue
				}
			}

			err = handleMessage(m)
			if err != nil {
				//log.Println("[!] handler ERROR")
				log.Println(err.Error())
				db.Model(&StatusUpdateQueue{}).Save(&StatusUpdateQueue{
					CaseID:    cast.ToUint64(m.CaseID),
					NewStatus: FetchFailedMessage,
					RefID:     cast.ToUint64(m.MessageID),
				})
			}

			// we either fetched the message or will retry it in the future if failed,
			// it should not be processed again
			db.Model(&Message{}).Where(Message{MessageID: m.MessageID}).
				Assign(&m).
				Create(&m)

			fmt.Printf("MessageID: %s\tCaseID: %s\tRefID: %s\tMessage: %s\n", m.MessageID, m.CaseID, m.RefID, m.Message)

		}
		fmt.Println("Messages fetch cycle complete. Waiting...")
		time.Sleep(time.Second * 3)
	}
}

func Runner() {

	go syncWithTavanir()
	//go syncTavanirStatus()
	go fetchMessages()
	//go fixTavanirExpertStatus()
	//go fixArabicCharacters()

}
