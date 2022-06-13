/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:

843
Sample Output:

348
*/

package main

import "fmt"

func main() {
	var num int
	var slc []int
	fmt.Scan(&num)

	for num > 0 {
		n := num % 10
		slc = append(slc, n)
		num = num / 10
		//fmt.Print(slc)
	}
	for _, nn := range slc {
		if nn > 0 {
			fmt.Print(nn)
		}
	}
}
