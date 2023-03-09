package main

import "fmt"

func rest_a_b(int a, int b) {
	fmt.Println("Is equal:", a + b)
	
}

func main() {
	fmt.Printf("Hola, es: %s", 1 + 2)
	fmt.Printf("Hola, es: %s", rest_a_b(1, 2))
}