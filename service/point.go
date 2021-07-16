package service

import (
	"github.com/gin-gonic/gin"
	"rubix-lib-rest-go/model"

	"gorm.io/gorm"
	"log"
)

func GetPoint(db *gorm.DB, id string) (*model.Point, error) {
	var err error
	point := new(model.Point)

	if err := db.Where("id = ? ", id).Preload("Tags").First(&point).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return point, err
}

func GetPoints(c *gin.Context, db *gorm.DB, args model.Args) ([]model.Point, int64, int64, error) {
	var points []model.Point
	var filteredData, totalData int64

	table := "points"
	query := db.Select(table + ".*")
	query = query.Offset(Offset(args.Offset))
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))

	if err := query.Find(&points).Error; err != nil {
		log.Println(err)
		return points, filteredData, totalData, err
	}

	// // Count filtered table
	// // We are resetting offset to 0 to return total number.
	// // This is a fix for Gorm offset issue
	query = query.Offset(0)
	query.Table(table).Count(&filteredData)

	// // Count total table
	db.Table(table).Count(&totalData)

	return points, filteredData, totalData, nil
}

// SavePoint both creates and updates network according to if ID field is empty or not
func SavePoint(db *gorm.DB, point *model.Point) (*model.Point, error) {
	if err := db.Save(&point).Error; err != nil {
		return point, err
	}
	return point, nil
}


//AddPoint both creates and updates according to the ID
func AddPoint(db *gorm.DB, item *model.Point, DeviceUUID string) (*model.Point, error) {
	device := new(model.Device)
	if err := db.Where("uuid = ? ", DeviceUUID).First(&device).Error; err != nil {
		return nil, &NotFoundError{deviceName, DeviceUUID, nil, "not found"}
	}
	if err := db.Create(&item).Error; err != nil {
		return item, &NotFoundError{deviceName, "na", err, "not found"}
	}
	return item, nil
}




// DeletePoint soft deletes all records.
func DeletePoint(db *gorm.DB, id string) error {
	point := new(model.Point)
	if err := db.Where("id = ? ", id).Delete(&point).Error; err != nil {
		log.Println(err)
		return err
	}

	//tag := new(model.Tag)
	//if err := db.Where("post_id = ? ", id).Delete(&tag).Error; err != nil {
	//	log.Println(err)
	//}

	return nil
}
