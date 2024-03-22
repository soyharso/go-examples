package main

import "fmt"

func f(left, right chan string) {
	a := <-right
	left <- a + string(byte(a[len(a)-1])+1)

}

func main() {

	const n = 58

	leftmost := make(chan string)

	right := leftmost

	left := leftmost

	for i := 0; i < n; i++ {

		right = make(chan string)

		go f(left, right)

		left = right

	}

	go func(c chan string) { c <- "A" }(right)

	fmt.Println(<-leftmost)
}
