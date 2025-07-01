package main

import "fmt"

func Soma(a int, b int) int {
	return a + b
}

func main() {
	resultado := Soma(5, 3)
	fmt.Printf("Resultado da soma: %d\n", resultado)
}
