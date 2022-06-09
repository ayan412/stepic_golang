package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := 1000; i >= 1; i = i / 10 { // сначала выполнится 1-е усл-е внешнего цикла, затем полностью внутренний, а уже затем 2-е усл-е.
		for j := 1000; j >= 1; j = j / 10 {
			fmt.Println(a/i%10, b/j%10)
			if a/i%10 == b/j%10 {
				if a/i%10 != 0 {
					fmt.Print(a/i%10, " ")
				}
			}
		}
	}
}
