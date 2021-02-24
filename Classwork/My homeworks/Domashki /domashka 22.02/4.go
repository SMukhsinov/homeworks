package main

import "fmt"

func main () {
	var a int

	fmt.Scanf("%d \n", &a)

	if a%2 == 0 {
		fmt.Println("Четное число")
	} else {
		fmt.Println("Нечетное число")
		}

}