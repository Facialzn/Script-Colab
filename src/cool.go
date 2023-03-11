package main

import "fmt"
import "strings"

func sum_a_b(a int, b int) int {
	return a + b
}

func rest_a_b(a int, b int) int {
	return a - b
}

func main() {
	fmt.Printf("Hola, sin función: %d\n", 1+2)
	fmt.Printf("Hola, con función es: %d\n", sum_a_b(1, 2))
	fmt.Printf("Hola, con ister eg es: %d\n", rest_a_b(1, 2))
}
