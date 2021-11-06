package implement

type FibV2 struct{}

func (fib *FibV2) CalculateFibN(n int) int {
	if n <= 0 {
		return 0
	}
	return fib.doCalFibN(n, make(map[int]int, n-2))
}

func (fib *FibV2) doCalFibN(n int, cache map[int]int) int {
	if n <= 2 {
		return 1
	}
	v, ok := cache[n]
	if ok {
		return v
	}
	cache[n] = fib.doCalFibN(n-1, cache) + fib.doCalFibN(n-2, cache)
	return cache[n]
}
