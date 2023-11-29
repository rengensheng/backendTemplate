package models

type Role struct {
	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(255)" binding:"max=255"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"varchar(255)" binding:"max=255"`

	Id string `form:"id" json:"id" xorm:"varchar(255) pk notnull" binding:"max=255"`

	Menu string `form:"menu" json:"menu" xorm:"text" binding:"max=65535"`

	OrderNo string `form:"order_no" json:"order_no" xorm:"varchar(255)" binding:"max=255"`

	Remark string `form:"remark" json:"remark" xorm:"varchar(255)" binding:"max=255"`

	Revision int64 `form:"revision" json:"revision" xorm:"bigint"`

	RoleName string `form:"role_name" json:"role_name" xorm:"varchar(255)" binding:"max=255"`

	RoleValue string `form:"role_value" json:"role_value" xorm:"varchar(255)" binding:"max=255"`

	Status string `form:"status" json:"status" xorm:"varchar(255)" binding:"max=255"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(255)" binding:"max=255"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"varchar(255)" binding:"max=255"`
}

type SimpleRole struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}
