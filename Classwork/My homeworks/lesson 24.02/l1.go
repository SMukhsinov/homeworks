package main

import "fmt"

func main (){
	for n := 0; n <= 50; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}
