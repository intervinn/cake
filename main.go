package main

import (
	"fmt"
	"os"
)

func main() {
	p := NewParser()
	c := NewCake(p)
	err := c.Init()

	if len(os.Args) <= 1 {
		fmt.Println("not enough")
	}

	err = c.Execute(os.Args[1])

	if err != nil {
		panic(err)
	}
}
