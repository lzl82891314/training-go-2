package internal

type Fib interface {
	// CalculateFibN 计算斐波那契数列中第N个数的值
	// 斐波那契数列为：前两项都是1，从第三项开始，每一项的值都是前两项的和
	// 例如：1 1 2 3 5 8 13 ...
	CalculateFibN(n int) int
}
