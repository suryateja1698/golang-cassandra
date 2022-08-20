package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/suryateja1698/golang-cassandra/service"
)

func Routes(r *gin.Engine) {
	r.GET("/healthcheck", service.HealthCheck)
	r.POST("/player-details", service.AddPlayerDetails)
}
