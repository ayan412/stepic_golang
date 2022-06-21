/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные

Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные

Выведите количество минимальных элементов последовательности.

Sample Input:

3
21
11
4
Sample Output:

1
*/

package main

import "fmt"

func main() {
	var n, lst int
	fmt.Scan(&n)
	slc := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&lst) // почему он начинает считывать не с 3, а с 21?
		slc = append(slc, lst)
		//fmt.Println(slc)
	}
	var min = slc[0]
	var s int
	//fmt.Println(min)
	for _, el := range slc {
		if el < min {
			min = el
			s = 0
			s++
		} else if el == min {
			s++
		}
	}
	fmt.Println(s)
}
