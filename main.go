package main

import (
	"fmt"
	"mytmpl/utils/compress"
)

func main() {
	fmt.Println("--------------running-----------------")
	soruce := []string{"./fortest/1.txt"}
	err := compress.Zip(soruce, "fortest/1.zip")
	if err != nil {
		fmt.Println(err)
	}
	compress.Unzip("fortest/1.zip", "fortest")
}
