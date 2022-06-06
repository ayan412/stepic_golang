/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO
*/

package main

import "fmt"

func main() {
	var Fn, Sn, Tn, n int
	fmt.Scan(&n)
	Fn = n / 100
	Sn = (n / 10) % 10
	Tn = (n % 10)
	//fmt.Print(Fn,Sn,Tn)
	if (Fn != Sn) && (Sn != Tn) && (Tn != Fn) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
