package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func main() {
	resultado := soma(5, 3)
	fmt.Printf("Resultado da soma: %d\n", resultado)
}
