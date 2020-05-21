package main

import (
	"fmt"
	"mytmpl/config/yamlconf"
)

func main() {
	err := yamlconf.InitConfig("./config/conf.yaml")
	if err != nil {
		return
	}
	fmt.Println(yamlconf.Conf.Host)
	fmt.Println(yamlconf.Conf.Info.Aaa)
}

