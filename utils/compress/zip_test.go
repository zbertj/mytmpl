package compress

import (
	"testing"
)

func Test_Zip(t *testing.T) {
	soruce := []string{"../../fortest/1.txt"}
	err := Zip(soruce, "../../fortest/1.zip")
	if err != nil {
		t.Error(err)
	}
	t.Log("finished")
}

func Test_Unzip(t *testing.T) {
	soruce := "../../fortest/1.zip"
	err := Unzip(soruce, "../../fortest")
	if err != nil {
		t.Error(err)
	}
	t.Log("finished")
}
