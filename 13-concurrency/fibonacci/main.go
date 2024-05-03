package main

func concurrentFib(n int) []int {
	ch := make(chan int)
	fib := make([]int, 0)

	go fibonacci(n, ch)
	for x := range ch {
		fib = append(fib, x)
	}

	return fib
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
