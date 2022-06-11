/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

Входные данные

Сначала задано число NN — количество элементов в массиве (1 \leq N \leq 1001≤N≤100). Далее через пробел записаны NN чисел — элементы массива. Массив состоит из целых чисел.

Выходные данные

Необходимо вывести все элементы массива с чётными индексами.

Sample Input:

6
4 5 3 4 2 3
Sample Output:

4 3 2
*/

package main

import "fmt"

func main() {
	// здесь должен быть ваш код
	var lng, elm int
	var slc []int

	fmt.Scan(&lng)

	for i := 0; i < lng; i++ {
		fmt.Scan(&elm)
		slc = append(slc, elm)
	}
	//fmt.Println(slc)
	for ind, el := range slc {
		if ind%2 == 0 {
			fmt.Print(el, " ")
		}
	}
}
