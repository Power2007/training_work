package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	//添加一个healthz接口
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", defaultz)
	log.Println("start http server")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func defaultz(response http.ResponseWriter, request *http.Request) {
	//获取IP地址
	ip:=request.RemoteAddr
	ip=ip[0:strings.LastIndex(ip, ":")]
	log.Println("request IP is ", ip)
	//把request 头写入 response头
	for key, value := range request.Header {
		log.Println(key, value)
		response.Header().Set(key, strings.Join(value, ""))

	}
	version_string := os.Getenv("VERSION")
	response.Header().Set("VERSION", version_string)
	response.Write([]byte("sucess"))
}

func healthz(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}
