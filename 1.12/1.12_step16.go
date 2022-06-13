/*
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.

Входные данные

Сначала задано число NN — количество элементов в последовательности (1\leq N\leq1001≤N≤100). Далее через пробел записаны NN чисел — элементы последовательности. Последовательность состоит из целых чисел.

Выходные данные

Необходимо вывести единственное число - количество положительных элементов в последовательности.

Sample Input:

5
1 2 3 -1 -4
Sample Output:

3
*/

package main

import "fmt"

func main() {
	// здесь должен быть ваш код
	var lng, elm int
	var slc []int
	var pos []int

	fmt.Scan(&lng)

	for i := 0; i < lng; i++ {
		fmt.Scan(&elm)
		slc = append(slc, elm)
	}
	//fmt.Println(slc)
	for _, el := range slc {
		if el > 0 {
			pos = append(pos, el)
		}
	}
	fmt.Print(len(pos))
}