package main

import (
	"context"
	"flag"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"go-kit-example/service"
	"net/url"
)



func main(){

	/*
		直连api服务
	*/

	urlstr := flag.String("url", "", "服务url")
	flag.Parse()

	if *urlstr == ""{
		fmt.Println("请指定服务url")
		return
	}


	if u, err := url.Parse(*urlstr); err == nil{
		client := httptransport.NewClient("GET", u, service.EncodeEchoReq, service.DecodeEchoRsp)
		endpoint := client.Endpoint()

		req := service.EchoRequest{Text:"hello echo"}
		rsp, err := endpoint(context.Background(), req)
		if err != nil{
			fmt.Printf("error: %s\n", err.Error())
		}else{
			fmt.Printf("echo rsp:%v\n", rsp)
		}
	}else{
		fmt.Println("服务url错误")
	}

}
