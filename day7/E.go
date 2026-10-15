package main

import (
	"bufio"
	"fmt"
	"os"
)

func canAchieve(s string, n, k int, diff []int) bool {
	for i := 0; i <= n; i++ {
		diff[i] = 0
	}

	currentFlips := 0

	for i := 0; i < n; i++ {
		currentFlips ^= diff[i]

		originalBit := int(s[i] - '0')
		actualBit := originalBit ^ currentFlips

		if actualBit == 0 {
			if i+k > n {
				return false
			}
			currentFlips ^= 1
			diff[i+k] ^= 1
		}
	}

	return true
}

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n int
	fmt.Fscan(reader, &n)

	var s string
	fmt.Fscan(reader, &s)

	diff := make([]int, n+1)

	for k := n; k >= 1; k-- {
		if canAchieve(s, n, k, diff) {
			fmt.Fprintln(writer, k)
			return
		}
	}
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
