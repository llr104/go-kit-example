package utils

import (
	"github.com/hashicorp/consul/api"
	"log"
)

var client *api.Client
var serverId = "s1"

const ApiUrl = "http://127.0.0.1:9988"
const ApiPort = "9988"
func init() {
	cfg := api.DefaultConfig()
	c, err := api.NewClient(cfg)
	if err != nil{
		log.Fatal("api client error\n")
	}else{
		client = c
	}
}

func Register() {

	m := make(map[string]string)
	m["k1"] = "value1"
	m["k2"] = "value2"

	c := api.AgentServiceCheck{}
	c.HTTP = "http://127.0.0.1:9988/health"
	c.Interval = "5s"

	r := api.AgentServiceRegistration{}
	r.Name = "s1"
	r.ID = serverId
	r.Port = 9988
	r.Address = "127.0.0.1"
	r.Tags = []string{"test"}
	r.Meta = m
	r.Check = &c

	client.Agent().ServiceRegister(&r)

}

func UnRegister()  {
	client.Agent().ServiceDeregister(serverId)
}