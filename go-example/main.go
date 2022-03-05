package main

import (
	"go-example/hello"
	"go-example/variable"
	"go-example/table99"
	"fmt"
)

func main(){
	fmt.Print("IO - example:\n")
	hello.Sayhello()
	hello.Sayhi()

	fmt.Print("Var - example:\n")
	variable.ShowVarType()
	fmt.Print("Const - example:\n")
	variable.ShowConst()
	fmt.Print("Bool - example:\n")
	variable.ShowBool()
	fmt.Print("String - example:\n")
	variable.ShowString()
	fmt.Print("Chstring - example:\n")
	variable.ChangeString()

	fmt.Print("Table99\n")
	table99.Table99()
}
