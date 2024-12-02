package helpers

import "math"

func Sort(arr []int) *[]int {
	maxDigit := MostDigit(arr)
	for k := 0; k < maxDigit; k++ {
		array := make([][]int, 10)
		for i := 0; i < len(arr); i++ {
			array[GetDigit(arr[i], k)] = append(array[GetDigit(arr[i], k)], arr[i])
		}
		arr = []int{}
		for j := 0; j < 10; j++ {
			arr = append(arr, array[j]...)

		}
	}
	return &arr
}

func MostDigit(arr []int) (maxDigit int) {
	for i := 0; i < len(arr); i++ {
		if CountDigit(arr[i]) > maxDigit {
			maxDigit = CountDigit(arr[i])
		}
	}
	return
}

func CountDigit(num int) (digit int) {
	if num == 0 {
		digit = 1
		return
	}
	digit = int(math.Ceil(math.Log10(float64(num))))
	return
}

func GetDigit(num int, index int) int {
	return int(math.Abs(float64(num))/math.Pow(10, float64(index))) % 10
}
