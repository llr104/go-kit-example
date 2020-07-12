package main

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go-kit-test/service"
	"go-kit-test/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {


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
		http.ListenAndServe(":9988", r)
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


