package service
import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit-example/utils"
)

type HelloRequest struct {
	Text string `json:"uid"`
}

type HelloResponse struct {
	Text string `json:"result"`
}

type EchoRequest struct {
	Text string `json:"uid"`
}

type EchoResponse struct {
	Text string `json:"result"`
	SId  string `json:"sId"`
}


func HelloEnpoint(api IApi) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(*HelloRequest)
		return HelloResponse{
			Text: api.Hello(r.Text),
		}, nil
	}
}

func EchoEndpoint(api IApi) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(*EchoRequest)
		return EchoResponse{
			Text:api.Echo(r.Text),
			SId:utils.SId,
		}, nil
	}
}