package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rubix-lib-rest-go/model"
	adapter "rubix-lib-rest-go/mqtt"
)

type HealthController struct {
}

func Ping(mqttConn *adapter.MqttConnection) func(c *gin.Context) {
	return func(c *gin.Context) {
		h := model.Health{}
		if mqttConn.IsConnected() {
			h.Status = model.HealhStatusUp
			c.JSON(http.StatusOK, h)
		} else {
			h.Status = model.HealhStatusDown
			c.JSON(http.StatusInternalServerError, h)
		}
	}
}
