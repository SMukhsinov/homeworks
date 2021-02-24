package main

import "fmt"

func main (){

	var a int
	fmt.Scanf("%d", &a)

	n := 1
	for x := 1; x <= a; x++{
		n = x * n

	}
	fmt.Println(n)
}