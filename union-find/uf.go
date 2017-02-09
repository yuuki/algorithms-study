package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QuickFindUF struct {
	id []int
}

func NewQuickFindUF(n int) *QuickFindUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &QuickFindUF{id: id}
}

func (u *QuickFindUF) Union(p, q int) {
	pid := u.id[p]
	qid := u.id[q]
	for i := 0; i < len(u.id); i++ {
		if u.id[i] == pid {
			u.id[i] = qid
		}
	}
}

func (u *QuickFindUF) Connected(p, q int) bool {
	return u.id[p] == u.id[q]
}

func (u *QuickFindUF) Find(p int) int {
	return -1
}

func (u *QuickFindUF) Count() int {
	return -1
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	uf := NewQuickFindUF(n)
	for s.Scan() {
		o := strings.Split(s.Text(), " ")
		p, _ := strconv.Atoi(o[0])
		q, _ := strconv.Atoi(o[1])
		if !uf.Connected(p, q) {
			uf.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
		}
	}
}
