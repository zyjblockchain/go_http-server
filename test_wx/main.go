package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	AppID     = "wxa6ae8fbb4d391d8e"
	AppSecret = "dca3c0e108f5abf1506670fcb6c0b022"
)

// 微信公众号带参数的二维码生成实例
func main() {
	// 获取access_token
	getResp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxa6ae8fbb4d391d8e&secret=dca3c0e108f5abf1506670fcb6c0b022")
	if err != nil {
		log.Println("get access_token error:", err)
	}
	if getResp.Body == nil {
		log.Println("Get 返回空")
	}
	byteGet, _ := ioutil.ReadAll(getResp.Body)
	defer getResp.Body.Close()
	token := new(AccToken)
	json.Unmarshal(byteGet, token)
	access_token := token.Token
	fmt.Println("access_token:", access_token)
	// 休眠一秒
	time.Sleep(time.Second * 1)

	// 创建二维码ticket
	jsonData := []byte(`{"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_id": 123}}}`)
	reader := bytes.NewReader(jsonData)
	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+access_token, "application/json;charset=UTF-8", reader)
	if err != nil {
		log.Println("post error:", err)
	}
	if resp.Body == nil {
		log.Println("post 返回空")
	}
	byteResp, _ := ioutil.ReadAll(resp.Body)
	defer getResp.Body.Close()
	fmt.Printf("byteResp:%s", string(byteResp))

	creatTicket := new(CreatTicket)
	err = json.Unmarshal(byteResp, creatTicket)
	if err != nil {
		log.Println("json unmarshal error:", err)
	}
	fmt.Println("ticket:", creatTicket.ticket)
	fmt.Println("limitTime:", creatTicket.limitTime)
	fmt.Println("url:", creatTicket.url)
}

// access_token 返回的数据结构类型
type AccToken struct {
	Token     string `json:"access_token"`
	limitTime int    `json:"expires_in"`
}

// ticket 返回的数据结构类型
type CreatTicket struct {
	ticket    string `json:"ticket"`
	limitTime int    `json:"expire_seconds"`
	url       string `json:"url"`
}
