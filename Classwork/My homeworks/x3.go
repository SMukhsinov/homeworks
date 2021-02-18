package main

import "fmt"

func main (){
	var x, z int

	fmt.Printf("perviy", x)
	fmt.Scanf("первый катет", x)
	//fmt.Scanf("Второй катет", z)

	var y int
	y = ((x * x) + (z * z))
	fmt.Printf("Результат", y)
}
