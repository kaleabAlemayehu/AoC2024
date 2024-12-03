package main

import (
	"fmt"
	"github.com/kaleabAlemayehu/Aoc2024/lexer"
	"os"
)

func main() {
	// Reading file
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	// Parse file input and store them into their two array
	var left = []int{}
	var right = []int{}
	l := lexer.NewLexer(data)
	for i := 0; i < 2000; i++ {
		v := l.NextToken()
		if i%2 == 0 {
			left = append(left, int(v))
		} else {
			right = append(right, int(v))
		}
	}
	// Create fq counter the right list
	fq := make(map[int]int)
	for _, v := range right {
		i, ok := fq[v]
		if !ok {
			fq[v] = 1
		} else {
			fq[v] = i + 1
		}
	}
	var sum int
	// Loop over element of left and sum the value * frequency_of_value
	for _, v := range left {
		i, ok := fq[v]
		if ok {
			sum += v * i
		}
	}
	fmt.Println("the sum is:", sum)

}
