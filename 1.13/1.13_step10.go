/*
Цифровой корень
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .

По данному числу определите его цифровой корень.

Входные данные

Вводится одно натуральное число n, не превышающее 10^710
7
 .

Выходные данные
Вывести цифровой корень числа n.

Sample Input:

3456
Sample Output:

9
*/
package main

import "fmt"

func main() {
	var n int
	var slc, slc2 []int
	fmt.Scan(&n)
	if n < 9999999 {
		for n > 0 {
			num := n % 10
			slc = append(slc, num)
			n = n / 10
		}
		sum1 := 0
		for i := 0; i < len(slc); i++ {
			sum1 += slc[i]
		}
		//fmt.Println(sum1)
		if sum1 > 9 {
			for sum1 > 0 {
				num2 := sum1 % 10
				slc2 = append(slc2, num2)
				sum1 = sum1 / 10
			}
			sum2 := 0
			for i := 0; i < len(slc2); i++ {
				sum2 += slc2[i]
			}
			//fmt.Println(sum1)
			fmt.Println(sum2)
		} else if sum1 < 9 {
			fmt.Println(sum1)
		}

	}
}
