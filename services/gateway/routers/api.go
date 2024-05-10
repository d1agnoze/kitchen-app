package routers

import (
	v1 "github.com/d1agnoze/kitchen/services/gateway/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() error {
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	app.GET("/ping", v1.Ping)
	app.GET("/puck", v1.Puck)

	err := app.Run("0.0.0.0:8080")
	return err
}
