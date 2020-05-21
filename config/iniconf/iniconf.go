package iniconf

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

type config struct {
	Section struct {
		Enabled bool
		Path    string
	}
}

// Conf 全局使用
var Conf config

// InitConfig 初始化函数
func InitConfig(file string) error {

	err := gcfg.ReadFileInto(&Conf, file)
	if err != nil {
		fmt.Printf("Failed to parse config file: %s", err)
		return err
	}
	return nil
}

/*

import (
	"fmt"
	"mytmpl/config/iniconf"
)

func main() {
	err := iniconf.InitConfig("./config/conf.ini")
	if err != nil {
		return
	}
	fmt.Println(iniconf.Conf.Section.Enabled)
	fmt.Println(iniconf.Conf.Section.Path)
}
*/
