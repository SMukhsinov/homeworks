package main

import "fmt"

func main (){
	var a, b, c int

	print("Ввод \n")
	fmt.Scanf("%d%d%d", &a,&b,&c)

	print("Вывод \n")
	x := a + b + c
	fmt.Println(x)

}