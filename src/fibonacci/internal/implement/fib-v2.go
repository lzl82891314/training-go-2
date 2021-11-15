package implement

type FibV2 struct {
	cache map[int]int
}

func (fib *FibV2) CalculateFibN(n int) int {
	if n <= 2 {
		return 1
	}
	if fib.cache == nil {
		fib.cache = make(map[int]int, n)
		fib.cache[0], fib.cache[1] = 1, 1
	}
	if v, ok := fib.cache[n]; ok {
		return v
	}
	fib.cache[n] = fib.CalculateFibN(n-1) + fib.CalculateFibN(n-2)
	return fib.cache[n]
}
