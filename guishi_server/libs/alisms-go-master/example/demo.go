package main

import (
	"alisms-go-master/SmsClient"
	"log"
	"net/http"
)

const (
	accessKeyID     = "LTAIoJ9QOQK24bbm"
	secretAccessKey = "s2rb2ovh2hzTyBBaiX4GBiKJHWKCCx"
)

func main() {
	sc, err := SmsClient.NewSMSClient(accessKeyID, secretAccessKey)
	if err != nil {
		return
	}
	statusCode, _, _ := sc.SendSMS(SmsClient.Params{"15008227269", "段应许", "SMS_139910267", `{"code":"12345"}`})
	if statusCode == http.StatusOK {
		log.Println("发送成功")
	} else {
		log.Println("发送失败")
	}
}
