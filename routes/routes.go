package routes

import (
	"time"

	"backend/config"
	"backend/controllers/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
)

var (
	masterClientID        = "admin-cli"
	masterClientSecret    = "a923d503-0398-4d56-ab1c-9d0a4bd63bd6"
	masterClientSecretDev = "472b0199-caef-4cc4-9a37-3f44e687d5ec"
	clientID              = "lttr"
	clientSecret          = "c0cb0fe2-4b2a-4a20-b0c3-879aa0279fb6"
	clientSecretDev       = "fc8b269b-e7e5-428c-9819-5f03f39f3acc"
)

//StartGin function
func StartGin() {
	config := config.GetConfig()
	production := config.GetBool("production")
	if !production {
		masterClientSecret = masterClientSecretDev
		clientSecret = clientSecretDev
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "content-type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if origin == "https://tobiaswaagefeldballe.dk" {
				return true
			}
			return origin == config.GetString("appUrl")
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(ginkeycloak.RequestLogger([]string{"uid"}, "data"))
	router.Use(gin.Recovery())

	api := router.Group("/api")

	api.POST("/runfortran", exec.RunFortranSpeedtest)
	api.POST("/rungolang", exec.RunGolangSpeedtest)

	router.Run(":8001")
}
