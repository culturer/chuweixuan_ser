package models

import (
// "github.com/astaxie/beego/orm"
// _ "github.com/go-sql-driver/mysql"
)

//首页广告轮播
type TBanner struct {
	// 广告标号
	Id int64
	//产品编号
	ProductId int64
	// 图片连接
	Picture string
	// 商户编号
	StoreId int64
}
