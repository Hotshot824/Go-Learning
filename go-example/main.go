package main

import (
	// "go-example/hello"
	// "go-example/table99"
	// "go-example/variable"
	// "go-example/logop"
	"go-example/inter"
)

func main() {
	//basic IO
	/*
		hello.Sayhello()
		hello.Sayhi()
	*/

	//basic variable declaration
	/*
		variable.ShowVarType()
		variable.ShowConst()
		variable.ShowBool()
		variable.ShowString()
		variable.ChangeString()
	*/

	// basic for loop
	// table99.Table99()
	// table99.Table99plus()

	//basic logic operation
	//logop.RandomTest()
	//logop.GotoHere()
	//logop.LoopIter()
	//logop.SimplePointer()
	//logop.RunPointer()
	//logop.Fibon(20)

	//hello.Str_ch()
	//hello.Input_code()

	// logop.RunDecode()
	dog := inter.Dog{Name:"Kenny"}
	inter.ShowEat(&dog)
	inter.ShowRun(&dog)
}
