package main

import "fmt"

func main (){

	x := 0
	for n := 0; n < 20; n ++ {
		x += n

		fmt.Println(x)
	}
}