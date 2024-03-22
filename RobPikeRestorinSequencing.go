package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func fanIn(input1 <-chan Message, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Anna"))

	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg1.wait <- true

		//msg1 := <-c
		//fmt.Println(msg1.str)
		//msg1.wait <- true
	}

	fmt.Println("Eres Chido , me voy")
}

func boring(s string) <-chan Message {
	c := make(chan Message)
	WaitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", s, i), WaitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-WaitForIt
		}
	}()
	return c
}
