/*
Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.

Sample Input:

2 4
Sample Output:

4 2
*/

package main

import "fmt"

func main() {
	a := 3
	b := 6
	test(&a, &b)
}

func test(x1, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}
