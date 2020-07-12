package main

import (
	"flag"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go-kit-example/service"
	"go-kit-example/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	name := flag.String("name", "", "服务名")
	port := flag.Int("port", 0, "端口")
	flag.Parse()

	if *name == "" || *port == 0{
		fmt.Println("服务名或端口未指定")
		return
	}

	utils.SetConfig(*name, *port)

	utils.Register()
	s1 := httptransport.NewServer(service.HelloEnpoint(&service.Api{}), service.DecodeHello, service.EncodeHello)
	s2 := httptransport.NewServer(service.EchoEndpoint(&service.Api{}), service.DecodeEchoReq, service.EncodeEchoRsp)

	r := mux.NewRouter()


	r.Handle("/", s1)
	r.Handle("/hello", s1)
	r.Handle(`/echo/{uid}`, s2)
	//r.Methods("POST").Path(`/echo/{uid}`).Handler(s2)
	r.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status":"ok"}`))
	})


	exit := make(chan bool, 1)
	go func() {
		s := fmt.Sprintf(":%d", *port)
		http.ListenAndServe(s, r)
		exit <- true
	}()

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		fmt.Println("signalChan exit ")
		exit <- true
	}()

	if <-exit{
		fmt.Println("exit ")
		utils.UnRegister()
	}

}


