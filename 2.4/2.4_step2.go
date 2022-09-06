package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circeArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}
/*
Очень важно помнить о том, что аргументы в Go всегда копируются.
Если мы попытаемся изменить любое поле в функции circleArea,
оригинальная переменная не изменится.
Именно поэтому мы будем писать функции используя указатели:
 */

func main() {
	var c = Circle{0, 0, 5}
	fmt.Println(circeArea(c))
}







