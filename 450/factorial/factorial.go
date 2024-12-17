package main

import (
	"fmt"
	"strconv"
)

func factorial(num int64) []int64 {

	fact := make([]int64, num+1)
	fact[0] = 1
	var i int64
	for i = 1; i <= num; i++ {
		fact[i] = i * fact[i-1]
	}

	factStr := strconv.FormatInt(fact[num], 10)

	digits := make([]int64, len(factStr))

	for j, ch := range factStr {
		digits[j] = int64(ch - '0')
	}

	return digits

}

func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}
