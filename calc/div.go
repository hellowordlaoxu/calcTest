package calc

import "fmt"

func Div(a, b int) int {
	fmt.Println("Div called!")
	if b == 0 {
		panic("除数不应为0")
	}
	return a / b
}
