package calc

func Div(a, b int) int {
	if b == 0 {
		panic("除数不应为0")
	}
	return a / b
}
