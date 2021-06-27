package main

import "fmt"

func main() {
	age := 7
	if age >= 5 && age < 10 {
		fmt.Println("5以上10未満")
	} else {
		fmt.Println("それ以外")
	}
}
