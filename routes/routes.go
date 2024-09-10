package routes

import (
	"time"

	"backend/controllers/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
)

// StartGin function
func StartGin() {

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
			if origin == "https://app.tobiaswaagefeldballe.dk" {
				return true
			}
			if origin == "https://osandweb.dk" {
				return true
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(ginkeycloak.RequestLogger([]string{"uid"}, "data"))
	router.Use(gin.Recovery())

	api := router.Group("/api")

	api.Static("/assets", "./assets")
	api.POST("/runfortran", exec.RunFortranSpeedtest)
	api.POST("/rungolang", exec.RunGolangSpeedtest)

	router.Run(":8001")
}
