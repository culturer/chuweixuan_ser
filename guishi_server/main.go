package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"guishi_server/models"
	_ "guishi_server/routers"
)

func init() {
	models.RegiesterDB()
}

func main() {
	// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
