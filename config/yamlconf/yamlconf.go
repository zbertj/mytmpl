package yamlconf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// config 配置文件结构体
type config struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	Info struct {
		Aaa int
		Bbb int
	}
}

// Conf 全局的配置变量
var Conf config

// InitConfig 初始化配置函数
func InitConfig(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

/*
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
*/
