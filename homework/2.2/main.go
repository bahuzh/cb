package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func header(w http.ResponseWriter, r *http.Request) {
	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "cb version 1.1")
	name := os.Getenv("VERSION")
	fmt.Fprintln(w, name, "\n---------")
	//接收客户端 request，并将 request 中带的 header 写入 response header
	for key, value := range r.Header {
		fmt.Fprintln(w, key, ":", value)
	}
	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	httpCode := http.StatusOK
	remoteIp := r.RemoteAddr
	log.Println("访问IP", remoteIp, "HTTP状态码", httpCode)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}

func main() {
	fmt.Printf("Starting http server...\n")
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", header)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}
