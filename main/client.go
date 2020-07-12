package main

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"go-kit-example/service"
	"net/url"
)

func main(){


	u, _ := url.Parse("http://127.0.0.1:9988")

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
