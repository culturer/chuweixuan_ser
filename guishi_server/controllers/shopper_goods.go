package controller

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	"guishi_server/models"
	"time"
)

//添加商品分类
func (this *ShopperController) addProductType() {

	productType := &models.TProductType{}
	beego.Info(productType)
	err := this.ParseForm(productType)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	productTypeId, err := models.AddProductType(productType)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	productType.Id = productTypeId
	this.Data["json"] = map[string]interface{}{"status": 200, "productType": productType, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
}

func (this *ShopperController) addGoods() {

	product := &models.TProduct{}
	beego.Info(product)
	err := this.ParseForm(product)
	productId, err := models.AddProduct(product)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "productId": productId, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()

}
