package implement

type FibV4 struct{}

func (fib *FibV4) CalculateFibN(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	i, j := 2, 3
	for x := 3; x < n; x++ {
		j += i
		i = j - i
	}
	return i
}
