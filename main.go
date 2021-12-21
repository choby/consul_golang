package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/choby/consul_go/registry"
	"github.com/choby/consul_go/util"
	"github.com/gin-gonic/gin"
)

func main() {
	host := "127.0.0.1"
	port := 8500
	token := ""
	registryDiscoveryClient, err := registry.NewConsulServiceRegistry(host, port, token)

	ip, err := util.FindFirstNonLoopbackIP()
	if err != nil {
		panic(err)
	}

	fmt.Println(ip)
	rand.Seed(time.Now().UnixNano())

	si, _ := registry.NewDefaultServiceInstance("go-user-server", "", 8011,
		false, map[string]string{"user": "zyn2"}, "")

	registryDiscoveryClient.Register(si)

	r := gin.Default()
	r.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err = r.Run(":8011")
	if err != nil {
		registryDiscoveryClient.Deregister()
	}
}
