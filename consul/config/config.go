package config

import (
	"fmt"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

const (
	host = "127.0.0.1"
	port = 8500
)

var (
	client *consulapi.Client
)

func GetConsulConfig() *consulapi.Client {
	// 创建连接consul服务配置
	if client == nil {
		var err error
		config := consulapi.DefaultConfig()
		config.Address = fmt.Sprintf("%s:%d", host, port)
		client, err = consulapi.NewClient(config)
		if err != nil {
			log.Fatal("consul client error : ", err)
		}
	}
	return client
}
