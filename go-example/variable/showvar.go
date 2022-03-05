package variable

import "fmt"

func ShowVarType()  {
	a, b, c := 1, "One", 0.9999
	fmt.Println("a, b, c := 1, One, 0.9999")
	fmt.Println("a = ", a, "b =", b, "c =", c)
}

func ShowConst() {
	a := 1
	const pi float32= 3.1415926
	fmt.Println("a =", a, "pi = ", pi)
}

func ShowBool(){
	a, b := true, false
	fmt.Println("a =", a, "b =", b)
}

func ShowString() {
	var hello string = "hello"
	hi := "hi"
	fmt.Println(hello, hi)
}

func ChangeString() {
	hello := "Hello father!"
	fmt.Println("first hello =", hello)
	chhello := []byte(hello)
	temp := "Fucker"
	fucker := []byte(temp)
	for i, j:=6, 0 ; i<len(chhello)-1; i++{
		chhello[i] = fucker[j]
		j++
	}
	for i:=0 ; i<len(chhello); i++{
		fmt.Printf("%c", chhello[i])
	}
	fmt.Printf("\n")
}