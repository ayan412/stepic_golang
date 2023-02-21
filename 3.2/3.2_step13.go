package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func ParseString(s string) int64 {

	list := make([]rune, 0, 0)
	ListStr := make([]string, 0, 0)
	var ContList string

	for _, v := range s {
		if unicode.IsNumber(v) {
			list = append(list, v)
		}
	}
	//return list

	for _, v := range list {
		ListStr = append(ListStr, string(v))
	}

	for _, v := range ListStr {
		ContList += v
	}

	num, err := strconv.ParseInt(ContList, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func adding(a, b string) int64 {
	return ParseString(string(a)) + ParseString(string(b))
}

func main() {
	fmt.Println(adding("sd#$fa232*&^", "10d@@@#$#@$!><AV0"))
}
