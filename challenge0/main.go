package main

import (
	"fmt"
	"os"

	"github.com/kaleabAlemayehu/Aoc2024/helpers"
	"github.com/kaleabAlemayehu/Aoc2024/lexer"
)

func main() {
	// Reading file
	data, err := os.ReadFile("./input.txt")
	helpers.CheckError(err)
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
	left = *helpers.Sort(left)
	right = *helpers.Sort(right)
	// Subtract the difference
	var diff = []int{}
	for i := 0; i < 1000; i++ {
		if left[i] > right[i] {
			diff = append(diff, left[i]-right[i])
		} else {
			diff = append(diff, right[i]-left[i])
		}
	}
	// Add the difference
	var sum int
	for _, val := range diff {
		sum += val
	}
	fmt.Println("sum is: ", sum)
}
