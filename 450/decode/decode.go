package main

import "fmt"

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	if s[0] == '0' {
		return 0
	}
	for i := 2; i <= len(s); i++ {
		if int(s[i-1]-'0') >= 1 && int(s[i-1]-'0') <= 9 {
			dp[i] += dp[i-1]
		}
		twodigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twodigit >= 10 && twodigit <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

//'A', 'B', 'B', 'C' (i.e., "1 2 2 3")
//'A', 'B', 'W' (i.e., "1 2 23")
//'A', 'V', 'C' (i.e., "1 22 3")
//'L', 'B', 'C' (i.e., "12 2 3")
//'L', 'W' (i.e., "12 23")

func main() {
	fmt.Println(numDecodings("12"))
}
