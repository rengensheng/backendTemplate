package models

type Dept struct {
	Id string `form:"id" json:"id" xorm:"pk notnull"`

	Revision int64 `form:"revision" json:"revision"`

	CreatedBy string `form:"created_by" json:"created_by"`

	CreatedTime string `form:"created_time" json:"created_time"`

	UpdatedBy string `form:"updated_by" json:"updated_by"`

	UpdatedTime string `form:"updated_time" json:"updated_time"`

	DeptName string `form:"dept_name" json:"dept_name"`

	ParentDept string `form:"parent_dept" json:"parent_dept"`

	OrderNo int64 `form:"order_no" json:"order_no"`

	Remark string `form:"remark" json:"remark"`

	Status string `form:"status" json:"status"`
}
