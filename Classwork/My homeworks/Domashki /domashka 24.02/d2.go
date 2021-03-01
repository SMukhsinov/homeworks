package main

import "fmt"

func main() {
	//var a,n int
	for a, n := 0, 1; n < 10; a, n = a + n, a {
		fmt.Println(a)
	}
}
