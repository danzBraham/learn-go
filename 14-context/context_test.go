package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "B", "Bobby")
	contextC := context.WithValue(contextA, "C", "Cleo")

	contextD := context.WithValue(contextB, "D", "Dragon")
	contextE := context.WithValue(contextB, "E", "Eagle")

	contextF := context.WithValue(contextC, "F", "Fang")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("F")) // Got
	fmt.Println(contextF.Value("C")) // Got, because C is the parent of F
	fmt.Println(contextF.Value("B")) // Not got, because differernt parent
	fmt.Println(contextA.Value("B")) // Parent cannot get the value of child
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine())
}
