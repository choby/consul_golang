package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/choby/consul_go/consul/registry"
	"github.com/choby/consul_go/util"
	"github.com/gin-gonic/gin"
)

const (
	host  = "127.0.0.1"
	port  = 8500
	token = ""
)

func main() {

	//连接到服务中心
	registryDiscoveryClient, err := registry.NewConsulServiceRegistry(host, port, token)
	defer registryDiscoveryClient.Deregister()
	if err != nil {
		panic(err)
	}

	ip, err := util.FindFirstNonLoopbackIP()
	if err != nil {
		panic(err)
	}

	fmt.Println(ip)
	rand.Seed(time.Now().UnixNano())

	//创建新的实例
	si, _ := registry.NewDefaultServiceInstance("go-user-server",
		"",
		8011,
		false,
		map[string]string{"user": "zyn2"},
		"")

	registryDiscoveryClient.Register(si)

	r := gin.Default()
	r.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8011")
}
