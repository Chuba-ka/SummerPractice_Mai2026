package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	var a string
	fmt.Fscan(reader, &a)

	cell := func(i int) byte {
		if i == 0 || i == n+1 {
			return 'B'
		}
		return a[i-1]
	}

	const INF = 1000000000
	dp := make([]int, n+2)

	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	for i := 0; i <= n; i++ {
		if dp[i] == INF {
			continue
		}

		switch cell(i) {
		case 'B', 'L':
			for j := i + 1; j <= i+m && j <= n+1; j++ {
				if cell(j) != 'C' {
					if dp[i] < dp[j] {
						dp[j] = dp[i]
					}
				}
			}

		case 'W':
			j := i + 1
			if cell(j) != 'C' {
				if dp[i]+1 < dp[j] {
					dp[j] = dp[i] + 1
				}
			}
		}
	}

	if dp[n+1] <= k {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		solve(reader, writer)
	}
}
