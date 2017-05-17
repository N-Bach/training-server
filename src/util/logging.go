package util

import (
	"fmt"
	"encoding/json"
)

func PrintStr(v string) {
	fmt.Println(v)
}

func PrintObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}