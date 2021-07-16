package service

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/model"
	"gorm.io/gorm"

)




func GetDevice(db *gorm.DB, uuid string, args model.Args) (*model.Device, error) {
	var err error
	item := new(model.Device)
	withChildren, _ := WithChildren(args.WithChildren)
	if withChildren { // drop child to reduce json size
		query := db.Where("uuid = ? ", uuid).Preload("Point").First(&item)
		if query.Error != nil {
			return nil, &NotFoundError{deviceName, uuid, nil, "not found"}
		}
		return item, err
	} else {
		query := db.Where("uuid = ? ", uuid).First(&item)
		if query.Error != nil {
			return nil, &NotFoundError{deviceName, uuid, nil, "not found"}
		}
		return item, err
	}

}

//GetDevices get all items
func GetDevices(db *gorm.DB, args model.Args) ([]model.Device, int64, int64, error) {
	var items []model.Device
	var filteredData int64
	table := model.TableNames.Device
	query := db.Select(table + ".*")
	query = query.Offset(Offset(args.Offset))
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))
	withChildren, _ := WithChildren(args.WithChildren)

	if withChildren { // drop child to reduce json size
		query = db.Preload("Point").Find(&items)
		if query.Error != nil {
			return items, filteredData, query.RowsAffected, query.Error
		}
		return items, filteredData, query.RowsAffected,  query.Error
	} else {
		query = db.Find(&items)
		if query.Error != nil {
			return items, filteredData, query.RowsAffected,  query.Error
		}
		return items, filteredData, query.RowsAffected,  query.Error
	}

}


//AddDevice both creates and updates according to the ID
func AddDevice(db *gorm.DB, item *model.Device, NetworkID string) (*model.Device, error) {
	network := new(model.Network)
	if err := db.Where("uuid = ? ", NetworkID).First(&network).Error; err != nil {
		return nil, &NotFoundError{networkName, NetworkID, nil, "not found"}
	}
	if err := db.Create(&item).Error; err != nil {
		return item, &NotFoundError{deviceName, "na", err, "not found"}
	}
	return item, nil
}


// DeleteDevice delete an item by ID
func DeleteDevice(db *gorm.DB, uuid string) error {
	device := new(model.Device)
	point := new(model.Point)
	n := db.Where("uuid = ? ", uuid).Unscoped().Delete(&device) //delete device
	d := db.Where("device_uuid = ? ", uuid).Unscoped().Delete(&point)  //delete points
	if n.Error != nil || d.Error != nil  {
		return n.Error
	}
	r := n.RowsAffected
	if r == 0 {
		return &NotFoundError{deviceName, uuid, nil, "not found"}
	} else {
		return nil
	}
}

// DeleteByNetworkUUID delete network uuid
func DeleteByNetworkUUID(db *gorm.DB, networkUUID string) error {
	device := new(model.Device)
	//point := new(model.Point)
	//n := db.Where("network_uuid = ? ", networkUUID).Find(&device)
	//n := db.Find(&device, "network_uuid = ?", networkUUID)
	nn := db.Select(device).Where("network_uuid = ?", networkUUID)
	fmt.Printf("%+v", nn)
	//fmt.Printf("%+v", device)
	fmt.Println(4444)
	//fmt.Println(n)
	//for i := range device {
	//
	//	fmt.Println(i)
	//}

	fmt.Println(4444)
	return nil
	//d := db.Where("device_uuid = ? ", deviceUUID).Unscoped().Delete(&point)  //delete points
	//if n.Error != nil || d.Error != nil  {
	//	return n.Error
	//}
	//r := n.RowsAffected
	//if r == 0 {
	//	return &NotFoundError{deviceName, deviceUUID, nil, "not found"}
	//} else {
	//	return nil
	//}
}



