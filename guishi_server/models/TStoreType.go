package models

import (
// "github.com/astaxie/beego/orm"
// _ "github.com/go-sql-driver/mysql"
)

//商铺分类表
type TStoreType struct {
	// 商品分类编号
	Id int64
	// 商品分类名称
	Name string
	// 父分类编号
	PartnerId int64
	// 小区编号
	CommunityId int64
	//官方
	IsOffice bool
}
