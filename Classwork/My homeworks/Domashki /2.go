package main

import "fmt"

func main (){
	var a,b,c int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)

	if a+b>=c || a+c>=b || c+b>=a {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}

