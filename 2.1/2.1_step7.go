/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7
Sample Output:

4
*/
package main

import "fmt"

var n int
var slc []int

func minimumFromFour() int {
	for i := 1; i <= 4; i++ {
		fmt.Scan(&n)
		slc = append(slc, n)
	}
	smallest := slc[0]
	for _, val := range slc[1:] {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}
func main() {

}
