package main

import (
	"fmt"
	"os"
)


func main() {
	// 0から始まる整数リテラルは8進数として扱われる
	// 8進数の10は10進数で8
	sum := 100 + 010
	fmt.Println(sum)

	file, err := os.OpenFile("foo", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}