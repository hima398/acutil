package acutil

import (
	"bufio"
	"io"
	"strconv"
)

type Scanner struct {
	sc *bufio.Scanner
}

func NewScanner(r io.Reader) *Scanner {
	sc := bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}

func (s Scanner) nextInt() int {
	s.sc.Scan()
	i, _ := strconv.Atoi(s.sc.Text())
	return i
}
