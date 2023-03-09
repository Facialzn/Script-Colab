package main

import "fmt"

func rest_a_b(int a, int b) {
	fmt.Println("Is equal:", a + b)
	
}

func main() {
	a := 1
	b := 2
	c := a + b

	fmt.Printf("Hola, es: %s", c)
	fmt.Printf("Hola, es: %s", rest_a_b(a, b))
}