package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	res := math.Hypot(a, b)
	fmt.Println(res)

}
