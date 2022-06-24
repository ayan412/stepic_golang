/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные

Дано число n (0<n<100).

Выходные данные

Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

Sample Input:

10
Sample Output:

10 korov

*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	num1 := n % 10
	num2 := n % 100
	if n > 0 && n < 100 {
		if num1 == 1 && num2 != 11 {
			fmt.Printf("%d korova", n)
		} else if (num1 == 2 && num2 != 12) || (num1 == 3 && num2 != 13) || (num1 == 4 && num2 != 14) {
			fmt.Printf("%d korovy", n)
		} else {
			fmt.Printf("%d korov", n)
		}
	}
}
