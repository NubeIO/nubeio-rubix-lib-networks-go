package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rubix-lib-rest-go/controller/validators"
	"rubix-lib-rest-go/model"
	"rubix-lib-rest-go/service"
)


var err error

// GetNetwork godoc
// @Summary Show a Network
// @Description get by ID
// @Tags networks
// @ID
// @Accept  json
// @Produce  json
// @Param uuid path int true "Device ID"
// @Success 200 {array} model.Network
// @Router /networks/{uuid} [get]
func (base *Controller) GetNetwork(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var args model.Args
	var at = model.ArgsType
	var ad = model.ArgsDefault
	args.WithChildren = c.DefaultQuery(at.WithChildren, ad.WithChildren)
	data, err := service.GetNetwork(base.DB, uuid, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if data != nil {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
// GetNetworks godoc
// @Summary List networks
// @Description get networks
// @Tags networks
// @Accept  json
// @Produce  json
// @Success 200 "Success"
// @Router /networks/ [get]
func (base *Controller) GetNetworks(c *gin.Context) {
	var args model.Args
	var at = model.ArgsType
	var ad = model.ArgsDefault
	args.Sort = c.DefaultQuery(at.Sort, ad.Sort)
	args.Order = c.DefaultQuery(at.Order, ad.Order)
	args.Offset = c.DefaultQuery(at.Offset, ad.Offset)
	args.Limit = c.DefaultQuery(at.Limit, ad.Limit)
	args.Search = c.DefaultQuery(at.Search, ad.Search)
	args.WithChildren = c.DefaultQuery(at.WithChildren, ad.WithChildren)
	networks, filteredData, totalData, _err := service.GetNetworks(base.DB, args); if _err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": _err.Error()})
		return
	}
	data := model.NetworkData{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         networks,
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}


// AddNetwork godoc
// @Summary Create Network
// @Description Create Network
// @Tags networks
// @Accept  json
// @Produce  json
// @Param Network body object true "Network"
// @Success 200 "Success"
// @Router /networks/ [network]
func (base *Controller) AddNetwork(c *gin.Context) {
	data := new(model.Network)
	if data, err = validators.CheckNetworkJson(c, data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data, err =  validators.CheckAddNetwork(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data, err =  service.AddNetwork(base.DB, data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})

}


// UpdateNetwork godoc
// @Summary Update Network
// @Description Update Network
// @Tags networks
// @Accept  json
// @Produce  json
// @Param uuid path int true "Network ID"
// @Param Network body object true "Network"
// @Success 200 "Success"
// @Router /networks/{uuid} [put]
func (base *Controller) UpdateNetwork(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var at = model.ArgsType
	var ad = model.ArgsDefault
	data := new(model.Network)
	if data, err = validators.CheckNetworkJson(c, data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = base.DB.Where("uuid = ?", uuid).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = base.DB.Model(&data).Updates(data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	withChildren, all := service.WithChildren(c.DefaultQuery(at.WithChildren, ad.WithChildren))
	fmt.Println(all)
	if withChildren { // drop child to reduce json size
		data.Device = nil
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}


// DeleteNetwork DeleteNetwork godoc
// @Summary Delete Network
// @Description Delete Network
// @Tags networks
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param uuid path int true "Network ID"
// @Success 200 "Success"
// @Router /networks/{uuid} [delete]
func (base *Controller) DeleteNetwork(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	if err = service.DeleteNetwork(base.DB, uuid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uuid#" + uuid: "deleted"})
}
