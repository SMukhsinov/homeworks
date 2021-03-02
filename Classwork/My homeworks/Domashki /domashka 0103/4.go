package main

import "fmt"

//найти разницу между максимальным и минимальными элементами массива
//5
//10
//21
//31
//41
//-3
//максимальны - 41
//минимальный - -3
//41 - (-3) = 44

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
	maxi := 0
	min := 0
	raznica := arr[0]
	for i := 0; i < n; i++ {
		if arr[i] > maxi {
			maxi = arr[i]
		} else if arr[i] < min {
			//if arr[i] < min
			min = arr[i]
		}
		//}else{
		raznica = maxi + (-min)
	}
	//	fmt.Println(maxi)
	fmt.Println("максимальное значение", maxi)
	fmt.Println("минимальное значение", min)
	fmt.Println("разница", raznica)

}
