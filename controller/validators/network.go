package validators

import (
	"github.com/gin-gonic/gin"
	"nubeio-rubix-lib-rest-go/helpers"
	"nubeio-rubix-lib-rest-go/model"
	"gopkg.in/validator.v2"
	"net/http"
)


func CheckNetworkJson(c *gin.Context, data *model.Network)  (*model.Network, error) {
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return  nil, err
	}
	return data, nil
}


func CheckAddNetwork(data *model.Network) (*model.Network, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}
	data.Uuid, _ = helpers.MakeUUID()
	return data, nil
}


func CheckDeviceJson(c *gin.Context, data *model.Device)  (*model.Device, error) {
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return  nil, err
	}
	return data, nil
}


func CheckAddDevice(data *model.Device) (*model.Device, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}
	data.Uuid, _ = helpers.MakeUUID()
	return data, nil
}
