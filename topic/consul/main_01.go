package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

var consulClient, err = api.NewClient(api.DefaultConfig())

func main() {
	catalog := consulClient.Catalog()
	// fmt.Println(catalog.Services(nil))
	// fmt.Println(catalog.Service("redis", "primary", nil))
	s1, _, _ := catalog.Services(nil)
	for k, v := range s1 {
		fmt.Println("service == >", k)
		tag := ""
		if len(v) > 0 {
			tag = v[0]
		}
		s2, _, _ := catalog.Service(k, tag, nil)
		for _, service := range s2 {
			fmt.Println(service.ServiceName, "==>", service.Address, "==>", service.ServicePort, "==>", service.ServiceID)
		}
	}
}
