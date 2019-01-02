package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// req,err := http.NewRequest(http.MethodPost,"http://127.0.0.1:8001",)
	// if err != nil {
	// 	log.Fatal("request error:",err)
	// 	return
	// }
	// req.Header.Set("Content-Type","application/json;charset=UTF-8")
	// client := http.Client{}
	// resp,err := client.Do(req)
	// if err != nil {
	// 	log.Fatal("response error:",err)
	// 	return
	// }
	// body,err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// type PostData struct {
	// 	Method  string          `json:"method"`
	// 	Version string          `json:"jsonrpc"`
	// 	Id      uint64          `json:"id"`
	// 	Payload json.RawMessage `json:"params"`
	// }
	// var params json.RawMessage
	// params = json.RawMessage([]byte{1})
	// params.MarshalJSON()
	// data := &PostData{
	// 	Method:  "chain_getBlockByHeight",
	// 	Version: "2.0",
	// 	Id:      1,
	// 	Payload: params,
	// }

	// jsonData ,err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatal("marshal error:",err)
	// 	return
	// }

	jsonData := []byte(`{"jsonrpc": "2.0","method": "tx_sendTx","params": [],"id": 1}`)
	reader := bytes.NewReader(jsonData)

	resp, err := http.Post("http://127.0.0.1:8001", "application/json;charset=UTF-8", reader)

	if err != nil {
		log.Fatal("post error:", err)
		return
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("read response error:", err)
		return
	}
	fmt.Println(string(responseBody))
}
