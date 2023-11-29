package models

type Menu struct {
	Component string `form:"component" json:"component" xorm:"varchar(255)" binding:"max=255"`

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(255)" binding:"max=255"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"varchar(255)" binding:"max=255"`

	Icon string `form:"icon" json:"icon" xorm:"varchar(255)" binding:"max=255"`

	Id string `form:"id" json:"id" xorm:"varchar(255) pk notnull" binding:"max=255"`

	IsExt string `form:"is_ext" json:"is_ext" xorm:"varchar(255)" binding:"max=255"`

	Keepalive string `form:"keepalive" json:"keepalive" xorm:"varchar(255)" binding:"max=255"`

	MenuName string `form:"menu_name" json:"menu_name" xorm:"varchar(255)" binding:"max=255"`

	OrderNo int64 `form:"order_no" json:"order_no" xorm:"bigint"`

	ParentMenu string `form:"parent_menu" json:"parent_menu" xorm:"varchar(255)" binding:"max=255"`

	Permission string `form:"permission" json:"permission" xorm:"varchar(255)" binding:"max=255"`

	Revision int64 `form:"revision" json:"revision" xorm:"bigint"`

	RoutePath string `form:"route_path" json:"route_path" xorm:"varchar(255)" binding:"max=255"`

	Show string `form:"show" json:"show" xorm:"varchar(255)" binding:"max=255"`

	Status string `form:"status" json:"status" xorm:"varchar(255)" binding:"max=255"`

	Type string `form:"type" json:"type" xorm:"varchar(255)" binding:"max=255"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(255)" binding:"max=255"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"varchar(255)" binding:"max=255"`
}

type MenuMeta struct {
	Icon            string `json:"icon"`
	Title           string `json:"title"`
	HideMenu        bool   `json:"hideMenu"`
	IgnoreKeepAlive bool   `json:"ignoreKeepAlive"`
	ShowMenu        bool   `json:"showMenu"`
}
type MenuItem struct {
	Path       string     `json:"path"`
	Name       string     `json:"name"`
	Icon       string     `json:"icon"`
	Component  string     `json:"component"`
	Children   []MenuItem `json:"children"`
	Meta       MenuMeta   `json:"meta"`
	ParentMenu string     `json:"parentPath"`
	Id         string     `json:"id"`
	OrderNo    int64      `json:"orderNo"`
}
