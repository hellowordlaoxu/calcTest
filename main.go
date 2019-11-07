package main

import (
	"fmt"
	"go5期/gitTest/calc"
)

func main() {
	fmt.Println("calc.Add called!")
	res := calc.Add(10, 20)
	fmt.Println("Add(10 ,20) :", res)

	fmt.Println("calc.Sub called!")
	res = calc.Sub(30, 20)
	fmt.Println("Sub(10 ,20) :", res)
	//fmt

	fmt.Println("calc.Div called!")
	res = calc.Div(100, 20)
	fmt.Println("Div(100 ,20) :", res)
	fmt.Println("calc.Multi called!")
	res = calc.Multi(30, 20)
	fmt.Println("Multi(10 ,20) :", res)
}
