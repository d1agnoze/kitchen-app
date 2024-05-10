package v1

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/d1agnoze/kitchen/services/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	resp, err := handlers.Ping()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}

func Puck(c *gin.Context) {

	resp, err := handlers.Puck()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}

