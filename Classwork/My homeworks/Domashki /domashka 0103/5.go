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
	arr1 := [5]int{}
	chet := [5]int{}
	nechet := [5]int{}

	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			arr[i] = chet[i]
		}
		if arr[i]%2 == -1 {
			arr[i] = nechet[i]
		}
	}
	fmt.Println("Введенные", arr1)
	fmt.Println("Четные", chet)
	fmt.Println("Нечетные", nechet)
}



