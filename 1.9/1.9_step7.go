/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:

1234
Sample Output:

1
*/

package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n)
	if 1 <= n && n <= 9 {
		fmt.Print(n)
	} else if 10 <= n && n <= 99 {
		c = n / 10
		fmt.Println(c)
	} else if 100 <= n && n <= 999 {
		c = n / 100
		fmt.Println(c)
	} else if 1000 <= n && n <= 9999 {
		c = n / 1000
		fmt.Println(c)
	} else if n == 10000 {
		c = n / 10000
		fmt.Println(c)
	} else {
		fmt.Println("Read a condion!")
	}
}
