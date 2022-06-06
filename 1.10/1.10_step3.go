/*
Требуется написать программу, при выполнении которой с клавиатуры считываются
два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
*/
package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	sum := 0
	for i := A; i <= B; i++ {
		sum += i
		fmt.Println(A, i, sum)
	}
	fmt.Println(sum)
}
