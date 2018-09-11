package controller

import (
	// "fmt"
	"github.com/astaxie/beego"
	"guishi_server/models"
	"time"
)

//登录
func (this *ClientController) login() {

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
		if user.Pwd == pwd && user.IsCustomer == true && user.IsLock == false {
			user.Pwd = ""
			this.Data["json"] = map[string]interface{}{"status": 200, "user": user, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码错误 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//注册
func (this *ClientController) register() {
	tel := this.Input().Get("Tel")
	beego.Info(tel)
	if tel == "" {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	//判断该手机号是否已经注册
	user, err := models.GetUserByTel(tel)
	if err != nil {
		beego.Info(err)
		err := this.ParseForm(user)
		user.AddTime = time.Now().Unix()
		user.IsCustomer = true
		user.IsLock = false
		beego.Info(user)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		userId, err := models.AddUser(user)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "userId": userId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if user != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "该手机号已注册！", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
}
