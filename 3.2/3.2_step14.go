package main

import (
	"bufio"
	"fmt"

	//"image/color/palette"
	"os"
	//"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//inputData := scanner.Text()

	rFormat := strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), ",", "."), ";")
	fmt.Printf("%T", rFormat)
	// 	Listn := make([]int, 0, 0)
	// 	for _, v := range rFormat {
	// 		num, err := strconv.Atoi(v)
	// 		if err != nil {
	// 			fmt.Println("errrrro")
	// 		}
	// 		Listn = append(Listn, num)
	// 	}
	// 	for range Listn {
	// 		res := Listn[0] / Listn[1]
	// 		fmt.Println(res)
	// 	}

}
