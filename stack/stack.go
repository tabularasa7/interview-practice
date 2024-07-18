package stack

type Stack []int

func (s *Stack) Push(val int) {

	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		return -1
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
