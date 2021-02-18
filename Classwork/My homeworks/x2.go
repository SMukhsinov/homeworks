package main

import "fmt"

//2) y = kx + b (k-?,b-?)
//Ввод
//3 4
//5 6
//x1=3,y1=4
//x2=5,y2=6
func main (){
	var x1, x2, y1, y2 int

	print("Ввод х1; х2 \n")
	fmt.Scanf("%d %d", &x1, &y1)

	print("Ввод х2; у2 \n")
	fmt.Scanf("%d %d", &x2, &y2)

	k := (y1 - y2) / (x1 - x2)
	b := y2 - k * x2
	fmt.Printf("Значение к = %d \n", k)
	fmt.Printf("Значение b = %d \n", b)
}
