package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i <= 9; i++ {
		fmt.Println(arr[i])
	}
}
