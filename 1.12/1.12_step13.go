/*
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:

5
41 -231 24 49 6
Sample Output:

49
*/

package main

import "fmt"

func main() {
	// здесь должен быть ваш код
	var lng, elm int
	var slc []int
	fmt.Scan(&lng)

	for i := 0; i < lng; i++ {
		fmt.Scan(&elm)
		slc = append(slc, elm)
	}
	//fmt.Println(slc)
	fmt.Println(slc[3])
}
