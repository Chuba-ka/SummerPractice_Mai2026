package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for w := n; w >= i; w-- {
			dp[w] = (dp[w] + dp[w-i]) % MOD
		}
	}

	fmt.Println(dp[n])
}
