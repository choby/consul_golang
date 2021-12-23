package test

import (
	"fmt"
	"testing"

	"github.com/choby/consul_go/consul/config"
	consulapi "github.com/hashicorp/consul/api"
)

func TestConsulConfig(t *testing.T) {
	client := config.GetConsulConfig()
	// KV, put值
	values := "test"
	key := "go-consul-test/172.16.242.129:8100"
	client.KV().Put(&consulapi.KVPair{Key: key, Flags: 0, Value: []byte(values)}, nil)

	// KV get值
	data, _, _ := client.KV().Get(key, nil)
	fmt.Println(string(data.Value))

	// KV list
	datas, _, _ := client.KV().List("go", nil)
	for _, value := range datas {
		fmt.Println(value)
	}
	keys, _, _ := client.KV().Keys("go", "", nil)
	fmt.Println(keys)
}
