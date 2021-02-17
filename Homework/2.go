package main

import "fmt"

func main () {
	var a, b, c int
	fmt.Scanf("%d", &a)
	b = a + 3
	c = b * 3
	fmt.Printf("Первое число = %d \n", a)
	fmt.Printf("Второе число = %d \n", b)
	fmt.Printf("Третье число = %d \n", c)
}
