package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"

	"net/http"
)

func DecodeHello(c context.Context,  r*http.Request) (request interface{}, err error){
 	v := r.URL.Query().Get("text")
 	if v != ""{
 		return &HelloRequest{Text:v}, nil
	}
	return nil, errors.New("hello 参数错误")

}

func EncodeHello(c context.Context,w http.ResponseWriter, r interface{}) error{
	json.NewEncoder(w).Encode(r)
	return nil
}

func EncodeEchoReq(c context.Context, req *http.Request, r interface{}) error{
	er := r.(EchoRequest)
	req.URL.Path += "/echo/" + er.Text
	return nil
}

func EncodeEchoRsp(c context.Context,w http.ResponseWriter, r interface{}) error{
	json.NewEncoder(w).Encode(r)
	return nil
}


func DecodeEchoReq(c context.Context,  r*http.Request) (request interface{}, err error){
	v := mux.Vars(r)
	if s, ok := v["uid"]; ok == true{
		return &EchoRequest{Text:s}, nil
	}
	return nil, errors.New("echo 参数错误")

}


func DecodeEchoRsp(c context.Context, rsp *http.Response) (response interface{}, err error){

	if rsp.StatusCode >= 400{
		return nil, errors.New("req error")
	}else{
		r := EchoResponse{}
		err = json.NewDecoder(rsp.Body).Decode(&r)
		return r, err
	}

}

