package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	"go-kit-example/service"
	"go-kit-example/utils"
	"io"
	"net/url"
	"os"
	"time"
)

func main() {
	/*
		通过consul发现api服务
	*/

	sName := flag.String("sName", "", "服务名")
	flag.Parse()
	if *sName == ""{
		fmt.Println("请指定要访问的服务名")
		return
	}

	cfg := consulapi.DefaultConfig()
	cfg.Address = utils.ConsulAddr
	apiClient, err := consulapi.NewClient(cfg)
	if err != nil{
		fmt.Printf("apiClient: %s\n", err.Error())
		os.Exit(0)
	}

	consulClient := consul.NewClient(apiClient)
	logger := log.NewLogfmtLogger(os.Stdout)

	instancer := consul.NewInstancer(consulClient, logger, *sName, []string{"test"}, true)

	f :=  func(instance string) (endpoint.Endpoint, io.Closer, error){
		t, _ := url.Parse("http://"+instance)
		c := httptransport.NewClient("GET", t, service.EncodeEchoReq, service.DecodeEchoRsp)
		return c.Endpoint(), nil,nil
	}

	endpointer := sd.NewEndpointer(instancer, f, logger)
	endpoints, _ := endpointer.Endpoints()
	fmt.Printf("endpoints len:%d\n", len(endpoints))
	if len(endpoints) >0{

		{
			//指定访问第一个服
			endpoint := endpoints[0]
			req := service.EchoRequest{Text:"hello echo"}
			rsp, err := endpoint(context.Background(), req)
			if err != nil{
				fmt.Printf("first endpoint error: %s\n", err.Error())
			}else{
				fmt.Printf("first endpoint echo rsp:%v\n", rsp)
			}
		}


		{
			//简单负载均衡
			bin := lb.NewRoundRobin(endpointer)

			for   {
				endpoint, err := bin.Endpoint()
				if err != nil {
					fmt.Printf("NewRoundRobin error:%s\n", err.Error())
				}else{
					req := service.EchoRequest{Text:"hello echo"}
					rsp, err := endpoint(context.Background(), req)
					if err != nil{
						fmt.Printf("NewRoundRobin endpoint error: %s\n", err.Error())
					}else{
						fmt.Printf("NewRoundRobin endpoint echo rsp:%v\n", rsp)
					}
				}

				time.Sleep(1*time.Second)
			}

		}

	}

}