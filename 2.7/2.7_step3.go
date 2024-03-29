/*
Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.

Выходные данные

Вывести строку, которая получится после добавления символов '*'.

Sample Input:

LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
Sample Output:

L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
Sample Input:

LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
Sample Output:

L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O

*/
// логику решения подсмотрел в комментариях
package main

import (
	"fmt"
	"strings"

	//"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	for i, v := range input {
		if i == len(input)-1 {
			fmt.Printf("%s", string(v))
			break
		} else {
			fmt.Printf("%s*", string(v))
		}
	}

}

// решение №2
package main

import (
"fmt"
"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Print(strings.Trim(strings.ReplaceAll(input, "", "*"), "*"))
}