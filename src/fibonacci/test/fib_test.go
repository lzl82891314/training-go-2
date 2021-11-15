package test

import "fibonacci/internal/implement"
import "testing"

var n = 40

func BenchmarkFibV1_CalculateFibN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib := &implement.FibV1{}
		fib.CalculateFibN(n)
	}
}

func BenchmarkFibV2_CalculateFibN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib := &implement.FibV2{}
		fib.CalculateFibN(n)
	}
}

func BenchmarkFibV3_CalculateFibN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib := &implement.FibV3{}
		fib.CalculateFibN(n)
	}
}

func BenchmarkFibV4_CalculateFibN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib := &implement.FibV4{}
		fib.CalculateFibN(n)
	}
}
