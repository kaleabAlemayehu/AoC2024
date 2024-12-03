package main

import (
	// "fmt"
	// "github.com/kaleabAlemayehu/Aoc2024/lexer"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Reading file
	// data, err := os.ReadFile("./testd2.txt")
	data, err := os.ReadFile("./inputd2.txt")
	if err != nil {
		panic(err)
	}
	// Create super 2d array
	store := make([][]int, 1000)
	for i := range store {
		store[i] = make([]int, 8)
	}
	var i, j int
	for index := 0; index < len(data); {
		if data[index] == 10 {
			j++
			index++
			// New sub-array
			i = 0

		} else if data[index] == 32 {
			i++
			index++
			// Add new element in sub-array
		} else {
			// Read number
			val := readNumber(data, &index)
			store[j][i], err = strconv.Atoi(string(val))
			if err != nil {
				panic(err)
			}
		}

	}
	var totalSafe int
	for _, value := range store {
		// Check if it is safe
		totalSafe = totalSafe + CheckSafe(value)
	}
	fmt.Printf("totalSafe:%d\n", totalSafe)
}

func CheckSafe(arr []int) int {
	// Determine if it incremental or not
	var isInc bool
	if arr[0] > arr[1] && (arr[0]-arr[1] <= 3 && arr[0]-arr[1] >= 1) {
		isInc = false
	} else if arr[0] < arr[1] && (arr[1]-arr[0] <= 3 && arr[1]-arr[0] >= 1) {
		isInc = true
	} else {
		return 0
	}
	for i := 2; i < len(arr); i++ {
		if ((arr[i] > arr[i-1] && isInc) && (arr[i]-arr[i-1] >= 1 && arr[i]-arr[i-1] <= 3)) || arr[i] == 0 {

		} else if ((arr[i] < arr[i-1] && !isInc) && (arr[i-1]-arr[i] >= 1 && arr[i-1]-arr[i] <= 3)) || arr[0] == 0 {
		} else {
			return 0
		}
	}
	return 1
}

func readNumber(data []byte, i *int) []byte {
	position := *i
	for isDigit(data[*i]) {
		*i++
	}
	return data[position:*i]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
