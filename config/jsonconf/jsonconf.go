package jsonconf

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
	Info    info
}

type info struct {
	A int
}

// Conf 定义全局变量
var Conf configuration

// InitConfig 定义初始化函数
func InitConfig(file string) error {
	fp, _ := os.Open(file)
	defer fp.Close()
	decoder := json.NewDecoder(fp)

	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

/*
import (
	"fmt"
	"mytmpl/config/jsonconf"
)

func main() {
	err := jsonconf.InitConfig("./config/conf.json")
	if err != nil {
		return
	}
	fmt.Println(jsonconf.Conf.Path)
	fmt.Println(jsonconf.Conf.Enabled)
	fmt.Println(jsonconf.Conf.Info.A)
}
*/
