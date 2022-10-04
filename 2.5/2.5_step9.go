package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	strX := scanner.Text()
	strY := scanner1.Text()
	r := strings.Index(strX, strY)
	fmt.Println(r)

	/*
		package main

		import (
			"fmt"
			"strings"
		)

		func main() {
			var s1, s2 string
			fmt.Scan(&s1, &s2)
			res := strings.Index(s1, s2)
			fmt.Println(res)
		}
	*/
}
