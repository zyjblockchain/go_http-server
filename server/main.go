package main

import (
	"io"
	"net/http"
	"os/exec"
)

func main03() {
	// http请求处理
	http.HandleFunc("/gitpull", handler1)
	http.HandleFunc("/gitdev", handler2)
	// 	绑定监听地址和端口
	http.ListenAndServe(":8080", nil)
}

// 请求处理函数
func handler1(w http.ResponseWriter, r *http.Request) {
	// 下面是执行服务器上指定的shell脚本
	command1 := `./gitpull.sh`
	cmd := exec.Command("/bin/bash", "-c", command1)
	output, err := cmd.Output()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, string(output))
}

// 请求处理函数2
func handler2(w http.ResponseWriter, r *http.Request) {
	command2 := `./gitdev.sh`
	cmd := exec.Command("/bin/bash", "-c", command2)
	output, err := cmd.Output()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, string(output))
}
