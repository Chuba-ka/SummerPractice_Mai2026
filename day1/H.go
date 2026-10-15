package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 676767677

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m int
	if _, err := fmt.Fscan(reader, &n, &m); err != nil {
		return
	}

	b := make([]int, n)
	countT := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
		countT[b[i]]++
	}

	countLess := make([]int, m+1)
	for t := 0; t < m; t++ {
		countLess[t+1] = countLess[t] + countT[t]
	}

	var ans int64 = 1

	for i := 0; i < n; i++ {
		t := b[i]
		if t == 0 {
			continue
		}

		mn := m + 1
		if i > 0 {
			if b[i-1] < mn {
				mn = b[i-1]
			}
		}
		if i < n-1 {
			if b[i+1] < mn {
				mn = b[i+1]
			}
		}

		if mn >= t {
			ans = 0
			break
		}

		var choices int64
		if mn == t-1 {
			choices = int64(countLess[t])
		} else {
			choices = int64(countT[t-1])
		}

		ans = (ans * choices) % MOD
	}

	fmt.Fprintln(writer, ans)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	if _, err := fmt.Fscan(reader, &t); err == nil {
		for i := 0; i < t; i++ {
			solve(reader, writer)
		}
	}
}
