package logop

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomTest() {
	var g, l, e int = 0, 0, 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		if r := randomInt(1, 10); r >= 7 {
			fmt.Println("Greater is ", r)
			//fmt.Println(r, "Greater than five")
			g++
		} else if r >= 3 && r <= 6 {
			fmt.Println("Equal is ", r)
			e++
		} else {
			fmt.Println("Less is ", r)
			//fmt.Println(r, "Less than five")
			l++
		}
	}

	fmt.Println("Greater is ", g, " Less is ", l, " equal is ", e)
}

func GotoHere() {
	i := 0
Here:
	if i > 10 {
		return
	}
	println(i)
	i++
	goto Here
}

func LoopIter() {
	var arr = [5]string{"Apple", "Orange", "Banana", "Plum", "Grape"}
	for k, fruit := range arr {
		println(k, fruit)
	}
}

func SimplePointer() {
	var ptr *[5]string
	var arr = [5]string{"Apple", "Orange", "Banana", "Plum", "Grape"}
	ptr = &arr
	fmt.Println(ptr)
	for i , value := range ptr{
		fmt.Println(i , value)
	}
	fmt.Println(ptr[0])
}

func RunPointer(){
	var ptr *[5]string
	var arr = [5]string{"Apple", "Orange", "Banana", "Plum", "Grape"}
	ptr = &arr
	fmt.Println(ptr)
	fmt.Println("ptr address is ", &ptr)
	FucPointer(&ptr, 5)
	fmt.Println(ptr)
}

func FucPointer(a **[5]string, len int){
	fmt.Println("a address is ", &(*a))
	for i := 0; i < 5; i++{
		fmt.Println((*a)[i])
	}
	(*a)[0] = "Hello"
}

func Fibon(end int){
	a, b := 0, 1
	for i := 0; i < end; i++{
		fmt.Print(a, " ")
		a, b = b, a + b
	}
	fmt.Print("\n")
}


