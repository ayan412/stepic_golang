/*
Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.

Sample Input:

2 2
Sample Output:

4
*/

package main

import "fmt"

func testStep5(x1, x2 *int) {
	fmt.Println(*x1 * *x2)
}
func main() {
	a := 6
	b := 4
	testStep5(&a, &b)
}
