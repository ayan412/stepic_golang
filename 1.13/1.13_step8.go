/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:

5
1
8
100
0
12
Sample Output:

1
*/

package main

import "fmt"

func main() {
	var n, list int
	fmt.Scan(&n)
	slc := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&list)
		slc = append(slc, list)
		//fmt.Println(slc)
	}
	var max int = 0
	for _, el := range slc {
		if el == 0 {
			max++
		}
		//fmt.Println(max)
	}
	fmt.Println(max)
}
