package main

import "fmt"

func main() {

	var sugar = "This is testing sugar"

	var char rune
	char = 97
	fmt.Println("Hello Bangladesh!")
	fmt.Println(sugar)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!false)

	// fmt.Println(char)
	fmt.Println(char)

	fmt.Printf("%c", char)
}
