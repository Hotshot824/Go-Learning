package hello

import (
	"fmt"
	"reflect"
)

func Sayhello() {
	fmt.Print("Hello, GO!\n")
}

func Sayhi() {
	fmt.Print("Hi, GO!\n")
}

func Str_ch() {
	data := "Hello"
	fmt.Println(reflect.TypeOf(data))
	fmt.Printf("%s\n", data)
	b := []byte(data)
	fmt.Println(reflect.TypeOf(b))
	fmt.Printf("%s\n", b)
}

func Input_code(){
	var text string
	fmt.Scanf("%s", &text)
	return text
}
