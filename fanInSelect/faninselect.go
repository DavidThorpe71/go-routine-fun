package main

import (
	"fmt"
)

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

func getData(input ...string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < len(input); i++ {
			newData := `{
				"title": "message",
				"message": "` + input[i] + `"
				}`
			c <- newData
		}
	}()

	return c
}

func main() {

	c := fanIn(getData("hello", "cheese", "eggs"), getData("what", "now"))

	values := []string{}

	for i := 0; i < 5; i++ {
		values = append(values, <-c)
	}

	fmt.Println("json:", values)
}
