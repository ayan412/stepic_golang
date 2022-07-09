/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1

Sample Input:

Sample Output:
*/
package main

import "fmt"

func main() {
	fmt.Println(sumInt(4))
}
func sumInt(nums ...int) (int, int) {
	sum := 0
	for _, x := range nums {
		sum = sum + x
	}
	return len(nums), sum
}
