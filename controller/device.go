package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/controller/validators"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/service"
)

var logName string = "LOGS: devices"
var logError string = "ERROR: devices"


// GetDevice godoc
// @Summary Show a Device
// @Description get by ID
// @Tags devices
// @ID
// @Accept  json
// @Produce  json
// @Param id path int true "Device ID"
// @Success 200 {array} model.DeviceBody
// @Router /devices/{id} [get]
func (base *Controller) GetDevice(c *gin.Context) {
	id := c.Params.ByName("uuid")
	var args model.Args
	var at = model.ArgsType
	var ad = model.ArgsDefault
	args.WithChildren = c.DefaultQuery(at.WithChildren, ad.WithChildren)
	data, err := service.GetDevice(base.DB, id, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if data != nil {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// GetDevices godoc
// @Summary List devices
// @Description get devices
// @Tags devices
// @Param Limit query int false "/devices/?Limit=2"
// @Param Offset query int false "/devices/?Limit=0"
// @Param Sort query string false "/devices/?Sort=1"
// @Param Order query string false "/devices/?Order=DESC"
// @Param Search query string false "/devices/?Search="something"
// @Accept  json
// @Produce  json
// @Success 200 {array} model.DeviceBody
// @Router /devices/ [get]
func (base *Controller) GetDevices(c *gin.Context) {
	var args model.Args
	var at = model.ArgsType
	var ad = model.ArgsDefault
	// Define and get sorting field
	args.Sort = c.DefaultQuery(at.Sort, ad.Sort)
	args.Order = c.DefaultQuery(at.Order, ad.Order)
	args.Offset = c.DefaultQuery(at.Offset, ad.Offset)
	args.Limit = c.DefaultQuery(at.Limit, ad.Limit)
	args.Search = c.DefaultQuery(at.Search, ad.Search)
	args.WithChildren = c.DefaultQuery(at.WithChildren, ad.WithChildren)
	// Fetch results from database
	devices, filteredData, totalData, err := service.GetDevices(base.DB, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := model.DeviceData{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         devices,
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// AddDevice godoc
// @Summary Create Device
// @Description Create Device
// @Tags devices
// @Accept  json
// @Produce  json
// @Param data body model.DeviceBody true "input body"
// @Success 200 {array} model.DeviceBody
// @Router /devices/ [post]
func (base *Controller) AddDevice(c *gin.Context) {
	data := new(model.Device)
	if data, err = validators.CheckDeviceJson(c, data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data, err =  validators.CheckAddDevice(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data, err =  service.AddDevice(base.DB, data, data.NetworkUuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})

}

//func (base *Controller) AddDevice(c *gin.Context) {
//	data := new(model.Device)
//	if err := c.ShouldBindJSON(&data); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	data.Uuid, _ = helpers.MakeUUID()
//	data, err := service.AddDevice(base.DB, data, data.NetworkUuid)
//	if err != nil {
//		log.Println(logError, err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"data": data})
//}


// UpdateDevice godoc
// @Summary Update Device
// @Description Update Device
// @Tags devices
// @Accept  json
// @Produce  json
// @Param id path int true "Device ID"
// @Param Network body object true "Device"
// @Success 200 "Success"
// @Router /devices/{uuid} [patch]
func (base *Controller) UpdateDevice(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var at = model.ArgsType
	var ad = model.ArgsDefault
	var device model.Device
	if err := base.DB.Where("uuid = ?", uuid).First(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input  model.DeviceBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	base.DB.Model(&device).Updates(input)

	withChildren, all := service.WithChildren(c.DefaultQuery(at.WithChildren, ad.WithChildren))
	fmt.Println(all)
	if withChildren { // drop child to reduce json size
		device.Point = nil
	}
	c.JSON(http.StatusOK, gin.H{"data": device})
}

// DeleteDevice godoc
// @Summary Delete Device
// @Description Delete Device
// @Tags devices
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param uuid path int true "Device ID"
// @Success 200 "Success"
// @Router /devices/{uuid} [delete]
func (base *Controller) DeleteDevice(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	err := service.DeleteDevice(base.DB, uuid)
	fmt.Println(444, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid#" + uuid: "deleted"})
}
