package main

import (
	"fmt"

	"github.com/mnogu/go-calculator"
)

func main() {

	val, err := calculator.Calculate("2 + 2")
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
