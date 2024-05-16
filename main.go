package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// 创建一个Person对象
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	// 使用Marshal函数将Person对象序列化为JSON格式的字节数组
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}
	// 打印序列化后的JSON数据
	fmt.Println("Serialized JSON data:", string(jsonData))

	// 使用Unmarshal函数将JSON格式的字节数组反序列化为Person对象
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return
	}

	// 打印反序列化后的Person对象
	fmt.Println("Decoded Person:", decodedPerson)
}
