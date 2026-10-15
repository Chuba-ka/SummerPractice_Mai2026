package main

import (
	"bufio"
	"fmt"
	"os"
)

type State struct {
	len  int
	link int
	next [26]int
}

type SuffixAutomaton struct {
	sz   int
	last int
	st   []State
}

func NewSAM(maxLen int) *SuffixAutomaton {
	st := make([]State, 2*maxLen)
	st[0].len = 0
	st[0].link = -1
	for i := 0; i < 26; i++ {
		st[0].next[i] = -1
	}
	return &SuffixAutomaton{sz: 1, last: 0, st: st}
}

func (sam *SuffixAutomaton) Extend(c byte) {
	ch := int(c - 'a')
	cur := sam.sz
	sam.sz++

	sam.st[cur].len = sam.st[sam.last].len + 1
	for i := 0; i < 26; i++ {
		sam.st[cur].next[i] = -1
	}

	p := sam.last
	for p != -1 && sam.st[p].next[ch] == -1 {
		sam.st[p].next[ch] = cur
		p = sam.st[p].link
	}

	if p == -1 {
		sam.st[cur].link = 0
	} else {
		q := sam.st[p].next[ch]
		if sam.st[p].len+1 == sam.st[q].len {
			sam.st[cur].link = q
		} else {
			clone := sam.sz
			sam.sz++
			sam.st[clone].len = sam.st[p].len + 1
			sam.st[clone].next = sam.st[q].next
			sam.st[clone].link = sam.st[q].link

			for p != -1 && sam.st[p].next[ch] == q {
				sam.st[p].next[ch] = clone
				p = sam.st[p].link
			}
			sam.st[q].link = clone
			sam.st[cur].link = clone
		}
	}
	sam.last = cur
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	if _, err := fmt.Fscan(reader, &s); err != nil {
		return
	}

	n := len(s)
	sam := NewSAM(n)

	for i := 0; i < n; i++ {
		sam.Extend(s[i])
	}

	var totalSubstrings int64 = 0
	for i := 1; i < sam.sz; i++ {
		totalSubstrings += int64(sam.st[i].len - sam.st[sam.st[i].link].len)
	}

	fmt.Println(totalSubstrings)
}
