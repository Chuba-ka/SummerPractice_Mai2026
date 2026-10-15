package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	if n == 1 {
		fmt.Println(0)
		return
	}

	dp := make([]int64, n+1)
	dp[1] = 0

	for i := 2; i <= n; i++ {
		val := int64(i)
		dp[i] = dp[i-1] + val

		if i%2 == 0 {
			cost2 := dp[i/2] + val
			if cost2 < dp[i] {
				dp[i] = cost2
			}
		}

		if i%3 == 0 {
			cost3 := dp[i/3] + val
			if cost3 < dp[i] {
				dp[i] = cost3
			}
		}
	}

	fmt.Println(dp[n])
}
