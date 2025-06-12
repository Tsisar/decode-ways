package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i] number of ways to decode the substring s[0:i]
	dp := make([]int, n+1)
	dp[0] = 1

	// first character cannot be '0'
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		// one-digit decoding
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// two-digit decoding
		twoDigits, _ := strconv.Atoi(s[i-2 : i])
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the string to decode: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := numDecodings(input)
	fmt.Printf("Output: %d\n", result)
}
