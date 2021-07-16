package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nubeio-rubix-lib-rest-go/helpers"
	"nubeio-rubix-lib-rest-go/model"
	"nubeio-rubix-lib-rest-go/service"
)

//var _err error

// GetPoint godoc
// @Summary Show a Point
// @Description get by ID
// @Tags points
// @ID
// @Accept  json
// @Produce  json
// @Param id path int true "Point ID"
// @Success 200 "Success"
// @Router /points/{id} [get]
func (base *Controller) GetPoint(c *gin.Context) {
	id := c.Params.ByName("id")
	point, err := service.GetPoint(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
	}

	c.JSON(200, point)
}

// GetPoints godoc
// @Summary List points
// @Description get points
// @Tags points
// @Accept  json
// @Produce  json
// @Success 200 "Success"
// @Router /points/ [get]
func (base *Controller) GetPoints(c *gin.Context) {
	var args model.Args

	// Define and get sorting field
	args.Sort = c.DefaultQuery("Sort", "ID")

	// Define and get sorting order field
	args.Order = c.DefaultQuery("Order", "DESC")

	// Define and get offset for pagination
	args.Offset = c.DefaultQuery("Offset", "0")

	// Define and get limit for pagination
	args.Limit = c.DefaultQuery("Limit", "25")

	// Get search keyword for Search Scope
	args.Search = c.DefaultQuery("Search", "")


	// Fetch results from database
	points, filteredData, totalData, err := service.GetPoints(c, base.DB, args)
	if err != nil {
		c.AbortWithStatus(404)
	}

	// Fill return data struct
	data := model.PointData{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         points,
	}

	c.JSON(200, data)
}

// CreatePoint godoc
// @Summary Create Point
// @Description Create Point
// @Tags points
// @Accept  json
// @Produce  json
// @Param Network body object true "Point"
// @Success 200 "Success"
// @Router /points/ [point]
func (base *Controller) CreatePoint(c *gin.Context) {
	point := new(model.Point)

	err := c.ShouldBindJSON(&point)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	point, err = service.SavePoint(base.DB, point)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, point)
}


func (base *Controller) AddPoint(c *gin.Context) {
	data := new(model.Point)
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.Uuid, _ = helpers.MakeUUID()
	data, err := service.AddPoint(base.DB, data, data.DeviceUuid)
	if err != nil {
		log.Println(logError, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}




// UpdatePoint godoc
// @Summary Update Point
// @Description Update Point
// @Tags points
// @Accept  json
// @Produce  json
// @Param id path int true "Point ID"
// @Param Network body object true "Point"
// @Success 200 "Success"
// @Router /points/{id} [put]
func (base *Controller) UpdatePoint(c *gin.Context) {
	id := c.Params.ByName("id")

	point, err := service.GetPoint(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	err = c.ShouldBindJSON(&point)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	point, err = service.SavePoint(base.DB, point)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, point)
}

// DeletePoint godoc
// @Summary Delete Point
// @Description Delete Point
// @Tags points
// @ID
// @Accept  json
// @Produce  json
// @Param id path int true "Point ID"
// @Success 200 "Success"
// @Router /points/{id} [delete]
func (base *Controller) DeletePoint(c *gin.Context) {
	id := c.Params.ByName("id")

	err := service.DeletePoint(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, gin.H{"id#" + id: "deleted"})
}
