package xmlconf

import (
	"encoding/xml"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool   `xml:"enabled"`
	Path    string `xml:"path"`
	Info    struct {
		Aa int `xml:"aa"`
	}
}

// Conf 全局
var Conf configuration

// InitConfig 初始化
func InitConfig(file string) error {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer xmlFile.Close()

	if err := xml.NewDecoder(xmlFile).Decode(&Conf); err != nil {
		fmt.Println("Error Decode file:", err)
		return err
	}
	return nil
}

/*

import (
	"fmt"
	"mytmpl/config/xmlconf"
)

func main() {
	err := xmlconf.InitConfig("./config/conf.xml")
	if err != nil {
		return
	}
	fmt.Println(xmlconf.Conf.Path)
	fmt.Println(xmlconf.Conf.Info.Aa)
}

*/
