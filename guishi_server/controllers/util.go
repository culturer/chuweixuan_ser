package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"guishi_server/libs/alisms-go-master/SmsClient"
	// "guishi_server/models"
	"math/rand"
	"net/http"
	"time"
)

type UtilsController struct {
	beego.Controller
}

func (this *UtilsController) Get() {
	this.getMsg()
}

func (this *UtilsController) Post() {

	options := this.Input().Get("options")
	beego.Info(options)
	//检查请求的方法
	if options != "" {
		switch options {
		//登录
		case "getMsg":
			this.getMsg()
		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}
	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "options is null !", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//验证码结构---------------------------------
type registerCode struct {
	Phone string
	Code  string
}

//阿里密钥
const (
	accessKeyID     = "LTAIoJ9QOQK24bbm"
	secretAccessKey = "s2rb2ovh2hzTyBBaiX4GBiKJHWKCCx"
)

func (this *UtilsController) getMsg() {

	tel := this.Input().Get("Tel")

	//初始化短信发送端
	sc, err := SmsClient.NewSMSClient(accessKeyID, secretAccessKey)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "短信系统繁忙请稍后尝试 ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	beego.Info("生成验证码")
	//随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	beego.Info(code)
	this.SetSession("code", code)
	//发送验证码
	statusCode, text, _err := sc.SendSMS(SmsClient.Params{tel, "奇点", "SMS_139910267", fmt.Sprintf(`{"code":"%v"}`, code)}) //手机号，签名，短信模板，验证码
	if statusCode == http.StatusOK {
		this.Data["json"] = map[string]interface{}{"status": 200, "data": text, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	} else if _err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "data": _err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()

	} else {
		this.Data["json"] = map[string]interface{}{"status": 400, "data": text, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

}
