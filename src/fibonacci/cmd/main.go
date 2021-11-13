package main

import (
	"fibonacci/internal"
	"fibonacci/internal/implement"
	"fmt"
)

func main() {
	n := 10

	v1 := calFib(&implement.FibV1{}, n)
	fmt.Printf("算法 FibV1 计算 n=%v 的结果为：%v\n", n, v1)

	v2 := calFib(&implement.FibV2{}, n)
	fmt.Printf("算法 FibV2 计算 n=%v 的结果为：%v\n", n, v2)

	v3 := calFib(&implement.FibV3{}, n)
	fmt.Printf("算法 FibV3 计算 n=%v 的结果为：%v\n", n, v3)

	v4 := calFib(&implement.FibV4{}, n)
	fmt.Printf("算法 FibV4 计算 n=%v 的结果为：%v\n", n, v4)
}

func calFib(fib internal.Fib, n int) int {
	return fib.CalculateFibN(n)
}
