package main

import (
	"fmt"
	"time"
)

func main() {
	// IF - ELSE
	number := 57
	if number%2 == 0 {
		fmt.Print(number, " is an EVEN number")
	} else {
		fmt.Println(number, "is an ODD number")
	}
	fmt.Println()
	// SWITCH - BY VALUE
	switch number {
	case 1:
		fmt.Println(number, " is equal to 1")
	case 100:
		fmt.Println(number, " is equal to 100")
	}
	// SWITCH - BY CONDITION
	switch {
	case number > 80:
		fmt.Println(number, " is greater than 80")
	case number < 20:
		fmt.Println(number, " is greater than 20")
	default:
		fmt.Println(number, " didn't match any special case")
	}

	// FOR - CLASSIC
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println("Sum ->", sum) // 10 (1+2+3+4)
	// FOR - AS WHILE
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println("N ->", n) // 8 (1*2*2*2)
	// FOR - INFINITE
	//sum = 0
	//for {
	//	sum++ // repeated forever
	//}
	//fmt.Println(sum) // never reached
	// FOR - AS A FOR EACH
	strings := []string{"hello", "world"}
	for i, v := range strings {
		fmt.Println(i, v)
	}
	for i, ch := range "HELLO" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}
	// FOR - EXIT LOOP
	sum = 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println("Sum ->", sum) // 6 (2+4)
	// WITH MAPS
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// WITH CHANNELS
	ch := make(chan int)
	go func() {
		ch <- 5
		ch <- 10
		ch <- 15
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}

	// SELECT
	c1 := make(chan string)
	c2 := make(chan string)
	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(6 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "two"
	}()
	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received from Channel 1 ->", msg1)
		case msg2 := <-c2:
			fmt.Println("received from Channel 2 ->", msg2)
		}
	}

}
