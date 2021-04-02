package main

import (
	"fmt"

	"github.com/natural-affinity/aoc/day1repair"
)

func main() {
	report, mapped, err := day1repair.ReadReport("./day1repair/day1repair.input")
	product, err := day1repair.ProductOfTwo(report, mapped)

	fmt.Println(product)
	fmt.Println(err)
}
