package main

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		x %= 1000000007
	}
	return x
}

func main() {
	println(fib(3))
}

//为什么要对1000000007取模呢？
//1000000007是10位的最小质数，需要将计算结果 % 1000000007才能保证得出的结果在int 范围中；
