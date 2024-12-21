package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(strs []string) string {
	var encoded strings.Builder
	for _, str := range strs {
		encoded.WriteString(strconv.Itoa(len(str)))
		encoded.WriteString("#")
		encoded.WriteString(str)
	}
	return encoded.String()
}

func decode(str string) []string {
	var strs []string
	i := 0
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(str[i:j])
		strs = append(strs, str[j+1:j+1+length])
		i = j + 1 + length
	}
	return strs
}

func main() {
	// Example usage
	strs := []string{"hello", "world", "golang", ""}
	encoded := encode(strs)
	fmt.Println("Encoded:", encoded)

	decoded := decode(encoded)
	fmt.Println("Decoded:", decoded)
}
