/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:

50
Sample Output:

1 2 4 8 16 32
*/

package main

import "fmt"

func main() {
	var n, num int
	num = 1
	var slc0 []int
	slc0 = make([]int, 1)
	slc0[0] = 1
	var slc []int
	fmt.Scan(&n)
	for i := 0; num < n; i++ {
		num = num * 2
		if num > n {
			break
		}
		slc = append(slc, num)
	}
	slc0 = append(slc0, slc...)
	//fmt.Println(slc2)
	for _, el := range slc0 {
		fmt.Printf("%d ", el)
	}

}
