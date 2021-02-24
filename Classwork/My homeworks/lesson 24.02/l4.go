package main

import "fmt"

func main (){
	var a,b int

	chetniy :=0
	nechetniy :=0

	fmt.Scanf("%d %d", &a, &b)
	for x := a; x <= b; x++ {
		if x%2 == 0 {
			chetniy += 1
		} else {
			nechetniy += 1
		}
	}
	fmt.Println(chetniy)
	fmt.Println(nechetniy)
}
