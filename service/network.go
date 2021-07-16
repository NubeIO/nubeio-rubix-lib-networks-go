package service

import (
	"fmt"
	"rubix-lib-rest-go/model"
	"gorm.io/gorm"
	"log"
)



//GetNetwork get item by uuid
func GetNetwork(db *gorm.DB, uuid string, args model.Args) (*model.Network, error) {
	var err error
	item := new(model.Network)
	withChildren, all := WithChildren(args.WithChildren)
	fmt.Println(all)
	if withChildren { // drop child to reduce json size
		query := db.Where("uuid = ? ", uuid).Preload("Device").First(&item)
		if query.Error != nil {
			return nil, &NotFoundError{deviceName, uuid, nil, "not found"}
		}
		return item, err
	} else {
		query := db.Where("uuid = ? ", uuid).First(&item)
		if query.Error != nil {
			return nil, &NotFoundError{networkName, uuid, nil, "not found"}
		}
		return item, err
	}
}

//GetNetworks get all items
func GetNetworks(db *gorm.DB, args model.Args) ([]model.Network, int64, int64, error) {
	var items []model.Network
	var filteredData int64
	table := model.TableNames.Network
	query := db.Select(table + ".*")
	query = query.Offset(Offset(args.Offset))
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))
	withChildren, _ := WithChildren(args.WithChildren)

	if withChildren { // drop child to reduce json size
		query = db.Preload("Device").Find(&items)
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


//AddNetwork both creates and updates according to the ID
func AddNetwork(db *gorm.DB, item *model.Network) (*model.Network, error) {
	if err := db.Create(&item).Error; err != nil {
		return item, &NotFoundError{networkName, "na", err, "not found"}
	}
	return item, nil
}



// SaveNetwork SavePost both creates and updates network according to if ID field is empty or not
func SaveNetwork(db *gorm.DB, network *model.Network) (*model.Network, error) {

	if err := db.Save(&network).Error; err != nil {
		log.Println(err)
		return network, err
	}
	return network, nil
}


//
//// DeleteNetwork DeletePost soft deletes all records.
//func DeleteNetwork(db *gorm.DB, uuid string) error {
//	//network := new(model.Network)
//	//devices := new(model.Device)
//	//n := db.Where("uuid = ? ", uuid).Unscoped()
//	err := DeleteByNetworkUUID(db, uuid)
//	if err != nil {
//		return err
//	}
//	return nil
//}


// DeleteNetwork DeletePost soft deletes all records.
func DeleteNetwork(db *gorm.DB, uuid string) error {
	network := new(model.Network)
	//devices := new(model.Device)
	n := db.Where("uuid = ? ", uuid).Delete(&network)
	err := DeleteByNetworkUUID(db, uuid)
	if err != nil {
		return err
	}
	d := db.Where("network_uuid = ? ", uuid)
	//d := db.Where("network_uuid = ? ", uuid).Unscoped().Delete(&devices)
	if n.Error != nil || d.Error != nil  {
		return n.Error
	}
	r := n.RowsAffected
	if r == 0 {
		return &NotFoundError{networkName, uuid, nil, "not found"}
	} else {
		return nil
	}

}
