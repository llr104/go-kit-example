package main

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"go-kit-example/service"
	"go-kit-example/utils"
	"net/url"
)



func main(){

	/*
		直连api服务
	*/
	u, _ := url.Parse(utils.ApiUrl)

	client := httptransport.NewClient("GET", u, service.EncodeEchoReq, service.DecodeEchoRsp)
	endpoint := client.Endpoint()

	req := service.EchoRequest{Text:"hello echo"}
	rsp, err := endpoint(context.Background(), req)
	if err != nil{
		fmt.Printf("error: %s\n", err.Error())
	}else{
		fmt.Printf("echo rsp:%v\n", rsp)
	}

}
