package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	rFormat := strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), ",", "."), ";")

	Listn := make([]float64, 0, 0)
	for i, _ := range rFormat {
		num, err := strconv.ParseFloat(rFormat[i], 64)
		if err != nil {
			fmt.Println("Unable to convert string to int")
		} else {
			Listn = append(Listn, num)
		}
	}
	res := Listn[0] / Listn[1]
	fmt.Printf("%.4f", res)
}
