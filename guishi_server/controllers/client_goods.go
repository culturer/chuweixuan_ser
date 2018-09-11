package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"guishi_server/models"
	"time"
)

//首页广告轮播
func (this *ClientController) getBanner() {

	communityId, err := this.GetInt64("communityId")
	beego.Info(communityId)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	banners := make([]*models.TBanner, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_banner")
	_, err = qs.Filter("community_id", communityId).All(&banners)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "banners": banners, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

type Goods struct {
	ProductType *models.TProductType
	Count       int
	Products    []*models.TProduct
}

//获取商品数据
func (this *ClientController) getGoods() {

	beego.Info("getGoods")

	storeId, err := this.GetInt64("storeId")
	beego.Info(storeId)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	size, err := this.GetInt("size")
	beego.Info(size)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	index, err := this.GetInt("index")
	beego.Info(index)
	if index == 0 {
		//index从1开始计数
		index = 1
	}
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	productTypes := make([]*models.TProductType, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_product_type")
	_, err = qs.Filter("store_id", storeId).All(&productTypes)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	datas := make([]Goods, 0)

	for i := 0; i < len(productTypes); i++ {
		where := fmt.Sprintf("where product_type_id = %v", productTypes[i].Id)
		items, count, err := models.GetProductPage(index, size, where)
		if err != nil {
			beego.Info(err)
		}
		data := Goods{
			Products:    items,
			ProductType: productTypes[i],
			Count:       count,
		}
		datas = append(datas, data)
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "datas": datas, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

type Shopers struct {
	StoreType *models.TStoreType
	Count     int
	Stores    []*models.TStore
}

//获取商铺信息
func (this *ClientController) getShopper() {

	beego.Info("getShopper")
	communityId, err := this.GetInt64("communityId")
	isOffice, err := this.GetInt("isOffice")

	if isOffice == 0 {
		beego.Info(communityId)
		if err != nil {
			beego.Info(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
		}

		size, err := this.GetInt("size")
		beego.Info(size)
		if err != nil {
			beego.Info(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
		}

		index, err := this.GetInt("index")
		beego.Info(index)
		if index == 0 {
			//index从1开始计数
			index = 1
		}
		if err != nil {
			beego.Info(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
		}

		storeTypes := make([]*models.TStoreType, 0)
		o := orm.NewOrm()
		qs := o.QueryTable("t_store_type")
		_, err = qs.Filter("community_id", communityId).All(&storeTypes)

		datas := make([]Shopers, 0)

		for i := 0; i < len(storeTypes); i++ {
			where := fmt.Sprintf("where store_type_id = %v", storeTypes[i].Id)
			items, count, err := models.GetStorePage(index, size, where)
			if err != nil {
				beego.Info(err)
			}
			data := Shopers{
				Stores:    items,
				StoreType: storeTypes[i],
				Count:     count,
			}
			datas = append(datas, data)
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "shopper": datas, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	} else {
		store, err := models.GetOfficeByCommitId(communityId)
		if err != nil {
			beego.Info(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "shopper": store, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

}

func (this *ClientController) searchGoods() {

}
