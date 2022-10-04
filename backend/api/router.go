package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	elastic "github.com/olivere/elastic/v7"
	"models"
	"tools"
)

func reverseProxy() gin.HandlerFunc {

	// target := "localhost:3005"

	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req = c.Request
			req.URL.Scheme = "http"
			req.URL.Host = "192.168.110.137:3005"
			req.Host = "192.168.110.137:3005"
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func hd(c *gin.Context) {
	remote, err := url.Parse("http://192.168.110.137:3005")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	c.Request.Host = remote.Host
	proxy.ServeHTTP(c.Writer, c.Request)
}

func setupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(static.Serve("/gw/media/", static.LocalFile("/home/kowthar/www/media", false)))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://127.0.0.1:3000",
		"http://localhost:3030",
		"http://127.0.0.1:8080",
		"http://0.0.0.0:8080",
		"https://tavanir.example.com",
		"http://tavanir-dashboard.example.com",
		"https://tavanir-dashboard.example.com",
		"http://kins.abrbit.com",
		"http://kins-dashboard.abrbit.com",
	}
	config.AllowMethods = []string{
		http.MethodPost,
		http.MethodGet,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	}
	config.AllowHeaders = []string{
		"x-requested-with",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
	}
	config.AddExposeHeaders("X-Total-Count")

	r.Use(cors.New(config))

	r.Any("/gw/meta/", hd)

	adminAPI := r.Group("/gw/admin/api/v1/")
	{
		adminAPI.Use(authMiddleware())
		adminAPI.Any("/export/request/", adminExportRequest)
		adminAPI.Any("/export/excel/", adminExport)

		adminAPI.GET("/request/", adminRequestItems)
		adminAPI.GET("/request_change_log/", adminRequestChangelog)
		adminAPI.GET("/request/:id/", adminRequestItem)
		adminAPI.PUT("/request/:id/", adminUpdateRequest)

		adminAPI.Any("/export/tavanir/request/", adminExportTavanirRequest)
		adminAPI.Any("/export/tavanir/excel/", adminExportTavanir)
		adminAPI.GET("/tavanir_request/", adminTavanirRequestItems)
		adminAPI.GET("/tavanir/cases/", adminTavanirRequestItems)
		adminAPI.GET("/tavanir/docs/:id/", adminTavanirRequestDocuments)
		adminAPI.GET("/tavanir/case_changelog/", adminTavanirRequestChangelog)
		adminAPI.GET("/tavanir/similar/case/:id/", adminTavanirRequestSimilar)
		adminAPI.GET("/tavanir/case/:id/", adminTavanirRequestItem)
		adminAPI.GET("/tavanir/id/:id/", adminTavanirRequestItemById)
		adminAPI.PUT("/tavanir/case/:id/", adminTavanirUpdateRequest)
		adminAPI.GET("/tavanir/sync_queue/", adminTavanirSyncQueueItems)
		adminAPI.GET("/tavanir/messages/", adminTavanirMessageItems)

		adminAPI.GET("/user/", adminUserList)
		adminAPI.GET("/user/:id/", adminUserItem)
		adminAPI.PUT("/user/:id/", adminUserEdit)
		adminAPI.POST("/user/change_password/", adminUserChangePassword)
		adminAPI.POST("/user/", adminUserCreate)
		adminAPI.GET("/search/res/", func(c *gin.Context) {
			client, err := elastic.NewClient(
				elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
			)
			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			searchResult, err := client.Search().
				Index(tools.GetEnv("ES_INDEX", "request")).
				Pretty(true).
				Do(context.Background())

			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, searchResult)
		})

		adminAPI.POST("/search/res/*d", func(c *gin.Context) {
			client, err := elastic.NewClient(
				elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
			)
			rawBody, _ := c.GetRawData()
			searchResult, err := client.Search().
				Index(tools.GetEnv("ES_INDEX", "request")). // search in index "twitter"
				Source(string(rawBody)).
				Pretty(true).            // pretty print request and response JSON
				Do(context.Background()) // execute

			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, searchResult)
		})

		adminAPI.GET("/search/res_tavanir/", func(c *gin.Context) {
			client, err := elastic.NewClient(
				elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
			)
			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			searchResult, err := client.Search().
				Index(tools.GetEnv("ES_TAVANIR_INDEX", "case")).
				Pretty(true).
				Do(context.Background())

			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, searchResult)
		})

		adminAPI.POST("/search/res_tavanir/*d", func(c *gin.Context) {
			client, err := elastic.NewClient(
				elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
			)
			rawBody, _ := c.GetRawData()
			searchResult, err := client.Search().
				Index(tools.GetEnv("ES_TAVANIR_INDEX", "case")). // search in index "twitter"
				Source(string(rawBody)).
				Pretty(true).            // pretty print request and response JSON
				Do(context.Background()) // execute

			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			c.JSON(http.StatusOK, searchResult)
		})

	}

	v1 := r.Group("/gw/api/v1/")
	{
		v1.POST("/track/result/", trackResult)
		v1.PUT("/request/update/", UpdateRequest)
		v1.POST("/auth/login/", login)
		v1.POST("/damage/", DamagePost)
		v1.POST("/upload", upload)
		v1.POST("/steps/1/", step1)
		v1.POST("/steps/2/", step2)
		v1.POST("/steps/3/", step3)
		v1.POST("/steps/4/", step4)
		v1.POST("/steps/5/", step5)
		v1.POST("/steps/6/", step6)
	}

	return r
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getClaimsFromJWTHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "authentication error")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

var jwtKey = []byte("e&lyv+nviq+2ppww@fy6&yvos43rfwedkxg6pquj(jc!hs@3*sr65o^h)v")

type KinsClaims struct {
	UserID   uint   `json:"user_id"`
	Usrename string `json:"username"`
	Role     string `json:"role"`
	State    string `json:"state"`
	jwt.StandardClaims
}

// getClaimsFromJWTHeader returns user id from jwt
func getClaimsFromJWTHeader(c *gin.Context) (*KinsClaims, error) {
	t := c.GetHeader("Authorization")
	if t == "" {
		t = c.Query("token")
	}
	token, err := jwt.ParseWithClaims(t, &KinsClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*KinsClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("error in token")
}

func signJWT(u models.User) string {
	expirationTime := time.Now().Add(time.Hour * 240 * 30)

	claims := &KinsClaims{
		Role:     u.Role,
		State:    u.Province,
		UserID:   u.ID,
		Usrename: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	return tokenString
}
