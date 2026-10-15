package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF int = 1e9

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &c[i][j])
		}
	}

	if n == 1 {
		fmt.Println(0)
		return
	}

	numStates := 1 << n

	dp := make([][]int, numStates)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	dp[1][0] = 0

	for mask := 1; mask < numStates; mask++ {
		for u := 0; u < n; u++ {
			if dp[mask][u] == INF {
				continue
			}

			for v := 0; v < n; v++ {
				if (mask & (1 << v)) == 0 {
					nextMask := mask | (1 << v)
					cost := dp[mask][u] + c[u][v]
					if cost < dp[nextMask][v] {
						dp[nextMask][v] = cost
					}
				}
			}
		}
	}

	ans := INF
	allVisitedMask := numStates - 1

	for u := 0; u < n; u++ {
		if dp[allVisitedMask][u] != INF {
			totalCost := dp[allVisitedMask][u] + c[u][0]
			if totalCost < ans {
				ans = totalCost
			}
		}
	}

	fmt.Println(ans)
}
