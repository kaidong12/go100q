package basic

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type resume struct {
	Name string `json:"name" doc:"我的名字"`
}

func findDoc(stru interface{}) map[string]string {
	t := reflect.TypeOf(stru).Elem()
	doc := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
	}

	return doc

}

func StructTagDemo_1() {
	var stru resume
	stru = resume{"StructTagDemo_1"}
	doc := findDoc(&stru)
	fmt.Printf("name字段为：%s\n", doc["name"])

	bytes, err := json.Marshal(stru)
	fmt.Println(bytes)
	fmt.Println(string(bytes))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
