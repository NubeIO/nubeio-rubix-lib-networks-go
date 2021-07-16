package model

import (

	"time"
)

// -+ networks
// --+ network
// ---+ devices
// -----+ device
// ----+ points
// -----+ point
// ------+ pointStore




var TableNames = struct {
	Network   		string
	Device   		string

}{
	Network:   	"networks",
	Device:   	"devices",

}

type Args struct {
	Sort   			string
	Order  			string
	Offset 			string
	Limit  			string
	Search 			string
	WithChildren 	string
}

var ArgsType = struct {
	Sort   				string
	Order   			string
	Offset   			string
	Limit   			string
	Search   			string
	WithChildren 		string
}{
	Sort:   			"Sort",
	Order:   			"Order",
	Offset:   			"Offset",
	Limit:   			"Limit",
	Search:   			"Search",
	WithChildren:   	"WithChildren",
}

var ArgsDefault = struct {
	Sort   			string
	Order   		string
	Offset   		string
	Limit   		string
	Search   		string
	WithChildren   	string
}{
	Sort:   			"ID",
	Order:   			"DESC",
	Offset:   			"0",
	Limit:   			"25",
	Search:   			"",
	WithChildren:   	"false",
}



type Common struct {
	Name        	string `json:"name" validate:"min=1,max=255"  gorm:"type:varchar(255);unique;not null"`
	Description 	string `json:"description"`
	Enable bool   			`json:"enable"`
	Fault bool   				`json:"fault"`
	FaultMessage string   			`json:"fault_message"`
	EnableHistory bool   				`json:"history_enable"`
	CreatedAt time.Time     `json:"created_on"`
	UpdatedAt time.Time     `json:"updated_on"`

}

type CommonDevice struct {
	Name        	string `json:"name" validate:"min=1,max=255"  gorm:"type:varchar(255);unique;not null"`
	Description 	string `json:"description"`
}

type CommonPoint struct {
	Writeable bool   			`json:"writeable"`
}

type Network struct {
	Uuid			string 		`json:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	Common
	Device 			[]Device 	 `json:"devices" gorm:"constraint:OnDelete:CASCADE;"`
}


type Device struct {
	Uuid				string `json:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	Common
	NetworkUuid     	string  `json:"network_uuid" gorm:"TYPE:varchar(255) REFERENCES networks;not null;default:null"`
	Point 				[]Point `json:"points" gorm:"constraint:OnDelete:CASCADE"`

}

type Point struct {
	Uuid			string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	Common
	DeviceUuid     	string `json:"device_uuid" gorm:"TYPE:string REFERENCES devices;not null;default:null"`
	CommonPoint
	PointStore 	PointStore `json:"point_store" gorm:"constraint:OnDelete:CASCADE"`

}

type PointStore struct {
	Uuid			string `json:"uuid" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	PointUuid     	string `gorm:"TYPE:string REFERENCES points;not null;default:null"`
}


type NetworkBody struct {
	Name string `json:"name" binding:"required"`
	Port string `json:"port" binding:"required"`
}


type DeviceBody struct {
	Name string 			`json:"name" binding:"required"`
	Description string  	`json:"description"`
	NetworkID uint 			`json:"NetworkID" binding:"required"`
}


type NetworkData struct {
	TotalData    int64
	FilteredData int64
	Data         []Network
}

type DeviceData struct {
	TotalData    int64
	FilteredData int64
	Data         []Device
}

type PointData struct {
	TotalData    int64
	FilteredData int64
	Data         []Point
}

