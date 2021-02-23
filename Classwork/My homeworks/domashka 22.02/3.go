package main

import "fmt"

func main () {
	var a, b, c int

	fmt.Scanf("%d \n", &a)
	fmt.Scanf("%d \n", &b)
	fmt.Scanf("%d \n", &c)

	if a > b && b > c || a < b && b < c {
		fmt.Println("Средняя", b)
	}
	if b > a && a > c || b < a && a < c {
		fmt.Println("Средняя", a)
	}
	if a > c && c > b || a < c && c < b {
			fmt.Println("Средняя", c)
		}


}
