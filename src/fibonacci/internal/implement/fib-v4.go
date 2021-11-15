package implement

type FibV4 struct{}

func (fib *FibV4) CalculateFibN(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
