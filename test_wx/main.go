package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	AppID     = "wx1f6db98d761e4679"
	AppSecret = "fca7d817f6706c240cfbef7d554db891"
)

// 微信公众号带参数的二维码生成实例
func main() {
	// 获取access_token
	getUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppID, AppSecret)
	getResp, err := http.Get(getUrl)
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
	// access_token := "17_aeooho-j4LIGKK3YgIKrWtaHKuy4P9hgjO2nSyOhXdlzFd9IY60l6Ww6Qtat_Isv9AqaIv71UvqSCsr_AfHbLqxbTLNVU7DNzPLbYOy3AoL2pMwzD3htUdTCtT_KUu51q3Nj03ZR0BSfytwTECYiAHAAYU"
	// 创建二维码ticket
	jsonData := []byte(`{"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_str": "lemo"}}}`)
	reader := bytes.NewReader(jsonData)
	postUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", access_token)
	resp, err := http.Post(postUrl, "application/json;charset=UTF-8", reader)
	if err != nil {
		log.Println("post error:", err)
	}
	if resp.Body == nil {
		log.Println("post 返回空")
	}
	byteResp, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("byteResp:%s\n", string(byteResp))

	creatTicket := &CreatTicket{}
	err = json.Unmarshal(byteResp, creatTicket)
	if err != nil {
		log.Println("json unmarshal error:", err)
	}

	fmt.Println("ticket:\r\n", creatTicket.Ticket)
	// fmt.Println("limitTime:", creatTicket.limitTime)
	fmt.Println("url:\r\n", creatTicket.Url)

	// 对Ticket进行urlEncode
	urlEncode := url.QueryEscape(creatTicket.Ticket)
	fmt.Println("urlEncode:\r\n", urlEncode)
	// 通过ticket换取二维码
	urlTicket := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", urlEncode)
	respTicket, err := http.Get(urlTicket)
	if err != nil {
		log.Println("ticket 换取二维码 error:", err)
	}
	byteRespTicket, _ := ioutil.ReadAll(respTicket.Body)
	defer respTicket.Body.Close()
	fmt.Println("two wei ma :", string(byteRespTicket))

}

// access_token 返回的数据结构类型
type AccToken struct {
	Token     string `json:"access_token"`
	LimitTime int    `json:"expires_in"`
}

// ticket 返回的数据结构类型
type CreatTicket struct {
	Ticket string `json:"ticket"`
	// limitTime int    `json:"expire_seconds"`
	Url string `json:"url"`
}
