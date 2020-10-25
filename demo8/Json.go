package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`    // 更具不同类型区分其值
	Age  int64 `json:"age"`
}

func main() {
	j1 := person{Name: "Payne", Age: 32}
	//fmt.Println(j1)
	// 序列化
	a, err := json.Marshal(j1)
	if err != nil{
		fmt.Printf("err %v\n", err)
	}
	fmt.Println(string(a))
	// 反序列化
	str := `{"name"":"payne", "age":10}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)
	fmt.Printf("%#v\n",p2)
}
func Jsons()  {
	fmt.Println("a.txt")
}
