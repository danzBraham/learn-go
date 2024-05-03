package main

func countReports(numSentCh chan int) int {
	total := 0
	for {
		num, ok := <-numSentCh
		if !ok {
			break
		}
		total += num
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
