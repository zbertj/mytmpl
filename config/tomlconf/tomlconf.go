package tomlconf

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

// Conf 定义全局变量
var Conf tomlConfig

// InitConfig 定义初始化函数
func InitConfig(file string) error {
	if _, err := toml.DecodeFile(file, &Conf); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

/*
import (
	"fmt"
	"mytmpl/config/tomlconf"
)

func main() {
	err := tomlconf.InitConfig("./config/conf.toml")
	if err != nil {
		return
	}
	fmt.Printf("Title: %s\n", tomlconf.Conf.Title)
	fmt.Printf("Owner: %s (%s, %s), Born: %s\n",
		tomlconf.Conf.Owner.Name, tomlconf.Conf.Owner.Org, tomlconf.Conf.Owner.Bio,
		tomlconf.Conf.Owner.DOB)
	fmt.Printf("Database: %s %v (Max conn. %d), Enabled? %v\n",
		tomlconf.Conf.DB.Server, tomlconf.Conf.DB.Ports, tomlconf.Conf.DB.ConnMax,
		tomlconf.Conf.DB.Enabled)
	for serverName, server := range tomlconf.Conf.Servers {
		fmt.Printf("Server: %s (%s, %s)\n", serverName, server.IP, server.DC)
	}
	fmt.Printf("Client data: %v\n", tomlconf.Conf.Clients.Data)
	fmt.Printf("Client hosts: %v\n", tomlconf.Conf.Clients.Hosts)
}

*/
