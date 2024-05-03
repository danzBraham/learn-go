package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// test("Hello there Zidan!")
	// tryChan()
	// channelAsParameter()
	// inOutChannel()
	// bufferedChannel()
	// rangeChannel()
	// selectChannel()
	// defaultSelectChannel()
	// raceCondition()
	// mutex()
	// rwMutex()
	// deadlock()
	// waitGroup()
	// once()
	// pool()
	mapSync()
}

/*
Concurrency
Go was designed to be concurrent, which is a trait fairly unique to Go.
It excels at performing many tasks simultaneously safely using a simple syntax.
Concurrency is as simple as using the go keyword when calling a function.
*/
func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string) {
	sendEmail(message)
	time.Sleep((time.Millisecond * 500))
	fmt.Println("====================================")
}

/*
Channels
A typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
Like maps and slices, channels must be created before use. They also use the same make keyword.
*/
func tryChan() {
	ch := make(chan string)
	defer close(ch)

	go func() {
		time.Sleep(time.Second * 1)
		ch <- "Zidan"
		fmt.Println("Done sending data to the channel")
	}()

	data := <-ch
	fmt.Println(data)

	time.Sleep(time.Second * 3)
}

func giveMeResponse(ch chan string, s string) {
	time.Sleep(2 * time.Second)
	ch <- s
}

func channelAsParameter() {
	ch := make(chan string)
	defer close(ch)

	go giveMeResponse(ch, "Hello")

	data := <-ch
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func onlyIn(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Zidan Abraham"
}

func onlyOut(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func inOutChannel() {
	ch := make(chan string)
	defer close(ch)

	go onlyIn(ch)
	go onlyOut(ch)

	time.Sleep(3 * time.Second)
}

func bufferedChannel() {
	ch := make(chan string, 2)
	defer close(ch)

	ch <- "Zidan"
	ch <- "Abraham"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func rangeChannel() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Number %d", i)
		}
		close(ch)
	}()

	// If channel not closed yet, it will always iterate until the channel is closed
	for data := range ch {
		fmt.Println(data)
		fmt.Println("========")
	}
}

func selectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go giveMeResponse(ch1, "Hello")
	go giveMeResponse(ch2, "World")

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data from channel 1:", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data from channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func defaultSelectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go giveMeResponse(ch1, "Hello")
	go giveMeResponse(ch2, "World")

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data from channel 1:", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data from channel 2:", data)
			counter++
		default:
			fmt.Println("Waiting data...")
		}
		if counter == 2 {
			break
		}
	}
}

func raceCondition() {
	x := 0

	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)
}

// To solve race constion use mutex
func mutex() {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)
}

// RWMutex is for Read Write Mutex
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func rwMutex() {
	myAccount := BankAccount{
		RWMutex: sync.RWMutex{},
		Balance: 0,
	}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				myAccount.AddBalance(1)
				fmt.Println(myAccount.GetBalance())
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Balance:", myAccount.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func deadlock() {
	zidan := UserBalance{
		Name:    "Zidan",
		Balance: 1000000,
	}

	abra := UserBalance{
		Name:    "Abra",
		Balance: 500000,
	}

	// DEADLOCK
	go transfer(&zidan, &abra, 200000)
	go transfer(&abra, &zidan, 200000)

	time.Sleep(3 * time.Second)

	fmt.Printf("User: %s; Balance: %d\n", zidan.Name, zidan.Balance)
	fmt.Printf("User: %s; Balance: %d\n", abra.Name, abra.Balance)
}

func runAsync(group *sync.WaitGroup) {
	defer group.Done()

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func waitGroup() {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go runAsync(group)
	}

	group.Wait()
	fmt.Println("Completed")
}

var counter = 0

func onlyOnce() {
	counter++
}

func once() {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(onlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
}

func pool() {
	p := sync.Pool{}
	group := sync.WaitGroup{}

	p.Put("Zidan")
	p.Put("Abra")
	p.Put("Ham")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			data := p.Get()
			fmt.Println(data)
			p.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Completed")
}

func mapSync() {
	data := sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data.Store(i, i)
		}()
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Printf("%d : %d\n", key, value)
		return true
	})
}
