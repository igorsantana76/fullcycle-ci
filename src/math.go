package foo

import "fmt"

func main() {
	fmt.Println(soma(69, 200))
	fmt.Println(subitracao(69, 200))
	fmt.Println(multiplicacao(69, 200))
	fmt.Println(divisao(69, 200))
	fmt.Println(soma2(69, 200))
	fmt.Println(subitracao2(69, 200))
	fmt.Println(multiplicacao2(69, 200))
	fmt.Println(divisao2(69, 200))
}

func soma(a int, b int) int {
	return a + b
}

func subitracao(a int, b int) int {
	return a - b
}

func multiplicacao(a int, b int) int {
	return a * b
}

func divisao(a int, b int) int {
	return a / b
}

func soma2(a int, b int) int {
	return a + b
}

func subitracao2(a int, b int) int {
	return a - b
}

func multiplicacao2(a int, b int) int {
	return a * b
}

func divisao2(a int, b int) int {
	return a / b
}
