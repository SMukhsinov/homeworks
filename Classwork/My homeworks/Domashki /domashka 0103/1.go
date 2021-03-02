package main

import "fmt"

func main() {
	var n int
	fmt.Print("введите кол-во ввода:")
	fmt.Scanf("%d", &n)
	arr := [10]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	counter := 0
	nechet := 0
	for i := 0; i < n; i++ {
		number := arr[i]
		if number%2 == 0 {
			counter += 1
		} else {
			nechet += 1
		}
	}
	//fmt.Println(arr)
	fmt.Println("количество четных", counter)
	fmt.Println("количество нечетных", nechet)
}