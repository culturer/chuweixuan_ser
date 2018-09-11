package controller

import (
	// "fmt"
	"github.com/astaxie/beego"
	"guishi_server/models"
	"time"
)

//登录
func (this *ShopperController) login() {

	//获取数据信息
	pwd := this.Input().Get("Pwd")
	tel := this.Input().Get("Tel")

	beego.Info(tel)
	beego.Info(pwd)

	if tel == "" || pwd == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//判断该手机号是否已经注册
	user, err := models.GetUserByTel(tel)

	if err != nil {
		beego.Info(err)
	}
	if user != nil {
		if user.Pwd == pwd && user.IsCustomer == true && user.IsLock == false && user.IsSeller {

			store, err := models.GetStoreById(user.Id)
			if store == nil || err != nil {
				beego.Info(err)
				this.Data["json"] = map[string]interface{}{"status": 400, "msg": "请先注册店铺后再登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			user.Pwd = ""
			this.Data["json"] = map[string]interface{}{"status": 200, "user": user, "store": store, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码错误 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//注册
func (this *ShopperController) register() {
	tel := this.Input().Get("Tel")
	beego.Info(tel)
	if tel == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	code := this.Input().Get("code")
	beego.Info(code)
	beego.Info(this.GetSession("code"))
	if code == "" || code != this.GetSession("code") {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "验证码不正确！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//判断该手机号是否已经注册
	user, err := models.GetUserByTel(tel)
	if user.Id != 0 {
		beego.Info(user)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "该手机号已注册！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if err != nil {
		beego.Info(err)
		err := this.ParseForm(user)
		user.AddTime = time.Now().Unix()
		user.IsCustomer = true
		user.IsLock = false
		user.IsSeller = true
		beego.Info(user)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		userId, err := models.AddUser(user)
		beego.Info(userId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "msg": "注册成功！", "userId": userId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
}

//创建店铺
func (this *ShopperController) createStore() {

	beego.Info("createStore")

	store := &models.TStore{}
	beego.Info(store)
	err := this.ParseForm(store)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	store.AddTime = time.Now().Unix()
	beego.Info(store)
	storeId, err := models.AddStore(store)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	store.Id = storeId
	this.Data["json"] = map[string]interface{}{"status": 200, "store": store, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()

}

//修改店铺信息
func (this *ShopperController) updateStore() {

	beego.Info("updateStore")

	store := &models.TStore{}
	beego.Info(store)
	err := this.ParseForm(store)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	store.AddTime = time.Now().Unix()
	beego.Info(store)
	storeId, err := models.UpdateStore(store)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	store.Id = storeId
	this.Data["json"] = map[string]interface{}{"status": 200, "store": store, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()

}
