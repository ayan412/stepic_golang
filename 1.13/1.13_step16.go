/*

Из натурального числа удалить заданную цифру.

Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные

Вывести число без заданных цифр.

Sample Input:

38012732
3
Sample Output:

801272

*/

package main

import "fmt"

func main() {
	var n, t int
	var slc, slc2 []int
	fmt.Scan(&n, &t)
	for n > 0 {
		num := n % 10
		if num != t {
			slc = append(slc, num)
		}
		n = n / 10
	}
	for ind := len(slc) - 1; ind >= 0; ind-- {
		slc2 = append(slc2, slc[ind])
	}
	for _, val := range slc2 {
		if val >= 0 {
			fmt.Print(val)
		}
	}
}
