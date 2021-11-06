package implement

type FibV1 struct{}

func (fib *FibV1) CalculateFibN(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fib.CalculateFibN(n-1) + fib.CalculateFibN(n-2)
}
