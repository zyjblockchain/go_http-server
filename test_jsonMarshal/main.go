package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	lemoAdd := "Lemo83KCTD3H97YCAAANN69HGK7J6G5AAG5CDC52"
	jsonlemoAdd, err := json.Marshal(lemoAdd)
	if err != nil {
		fmt.Println("json01 error:", err)
	}
	data := PostData{
		Version: "2.0",
		Id:      1,
		Method:  "account_getBalance",
		Payload: []json.RawMessage{jsonlemoAdd},
	}

	res, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json02 error:", err)
	}
	fmt.Println(string(res))
}

type PostData struct {
	Version string            `json:"jsonrpc"`
	Id      uint64            `json:"id"`
	Method  string            `json:"method"`
	Payload []json.RawMessage `json:"params,omitempty"`
}
