package main

import (
	"fmt"
)

type testInterface interface{
	Print1()
	Print2()
}

type Token struct {
	a int
	b int
}

func (t *Token)Print1(){
	fmt.Println(t.a)
}

func (t *Token)Print2(){
	fmt.Println(t.b)
}

func main(){

}