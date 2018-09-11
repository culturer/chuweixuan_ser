package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//商品分类表
type TProductType struct {
	// 商品分类编号
	Id int64
	// 商品分类名称
	Name string
	// 父分类编号
	PartnerId int64
	// 商铺编号
	StoreId int64
}

func AddProductType(productType *TProductType) (int64, error) {
	o := orm.NewOrm()
	productTypeId, err := o.Insert(productType)
	return productTypeId, err
}
