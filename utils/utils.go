package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
)

var client *api.Client
var sName = "s1"
var sPort = 9988
var sId = uuid.New().String()

var ConsulAddr = "127.0.0.1:8500"

func init() {
	cfg := api.DefaultConfig()
	cfg.Address = ConsulAddr

	c, err := api.NewClient(cfg)
	if err != nil{
		log.Fatal("api client error\n")
	}else{
		client = c
	}
}

func SetConfig(serverName string, port int)  {
	sName = serverName
	sPort = port
}

func Register() {

	m := make(map[string]string)
	m["k1"] = "value1"
	m["k2"] = "value2"

	c := api.AgentServiceCheck{}

	h := fmt.Sprintf("http://127.0.0.1:%d%s", sPort, "/health")
	c.HTTP = h
	c.Interval = "5s"

	r := api.AgentServiceRegistration{}
	r.Name = sName
	r.ID = sId
	r.Port = sPort
	r.Address = "127.0.0.1"
	r.Tags = []string{"test"}
	r.Meta = m
	r.Check = &c

	client.Agent().ServiceRegister(&r)

}

func UnRegister()  {
	client.Agent().ServiceDeregister(sId)
}