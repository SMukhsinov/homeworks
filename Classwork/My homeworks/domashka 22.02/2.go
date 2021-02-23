package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanf("%d \n", &a)
	fmt.Scanf("%d \n", &b)
	fmt.Scanf("%d \n", &c)

	//if a + b <= c || a + c <= b || b + c <= a {
	//	fmt.Println("Треугольника не существует")
	//}
	if a == b && b == c && c == a {
		fmt.Println("Равностороннии треугольник")
	} else if a != b && a != c && b != c {
		fmt.Println("Разностороннии треугольник")
	} else {
		fmt.Println("Равнобедренный треугольник")
	}

}
