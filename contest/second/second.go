package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		count := make([]int, 5)
		for _, part := range parts {
			size, _ := strconv.Atoi(part)
			count[size]++
		}
		cars := 0
		cars += count[4]
		minpair := min(count[3], count[1])
		cars += minpair
		count[3] -= minpair
		count[1] -= minpair
		cars += count[3]
		cars += count[2] / 2
		if count[2]%2 != 0 {
			cars++
			count[1] -= min(2, count[1])
		}
		if count[1] > 0 {
			cars += (count[1] + 3) / 4
		}
		fmt.Fprintln(writer, cars)

	}
}
