package models

import (
// "github.com/astaxie/beego/orm"
// _ "github.com/go-sql-driver/mysql"
)

//小区
type TCommunity struct {
	// 小区编号
	Id int64
	// 小区经纬度JSON
	Position string
	// 小区地址
	Address string
}
