package main

import "fmt"

func main (){
	var a, b int

	print("Введите катет a и b \n")
	fmt.Scanf("%d %d", &a, &b)

	var x int
	x = ((a * a) + (b * b))
	fmt.Printf("Результат %d \n", x)
}
