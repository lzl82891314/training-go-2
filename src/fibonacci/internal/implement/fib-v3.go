package implement

type FibV3 struct{}

func (fib *FibV3) CalculateFibN(n int) int {
	if n <= 2 {
		return 1
	}
	s := make([]int, n)
	s[0], s[1] = 1, 1
	for i := 2; i < n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n-1]
}
