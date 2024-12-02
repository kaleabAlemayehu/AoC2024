package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/kaleabAlemayehu/Aoc2024/lexer"
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
	// Sort both array
	slices.Sort(left)
	slices.Sort(right)
	// Subtract the difference
	var sum int
	for i := 0; i < 1000; i++ {
		if left[i] > right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}
	}
	fmt.Println("sum is: ", sum)
	// Add the difference
}
