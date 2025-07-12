package main

import "fmt"

// shiftChar shifts a lowercase letter c by 'shift' positions in the alphabet,
// wrapping around and handling negative shifts correctly.
func shiftChar(c byte, shift int) byte {
	fmt.Println(int(c - 'a'))
	shifted := (int(c-'a') + shift%26 + 26) % 26
	fmt.Println(shifted)
	return byte(shifted) + 'a'
}

func main() {
	tests := []struct {
		c     byte
		shift int
	}{
		{'a', 1},   // b
		{'z', 1},   // a
		{'c', 28},  // e (28 % 26 = 2, c->e)
		{'d', -1},  // c
		{'a', -28}, // y (-28 % 26 = -2 + 26 = 24, a->y)
		{'m', 52},  // m (52 % 26 = 0, no change)
		{'n', -52}, // n (same as above)
	}

	for _, test := range tests {
		result := shiftChar(test.c, test.shift)
		fmt.Printf("shiftChar('%c', %d) = '%c'\n", test.c, test.shift, result)
	}
}
