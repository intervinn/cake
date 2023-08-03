package main

import (
	"fmt"
	"log"
	"os"

	"github.com/intervinn/cake"
)

func main() {
	p := cake.NewParser()
	c := cake.NewCake(p)
	err := c.Init()

	if len(os.Args) <= 1 {
		msg := "Not enough arguments, see current commands available:\n"
		for _, v := range c.Commands {
			msg += "- " + v.Name + "\n"
		}
		fmt.Print(msg)
		return
	}

	err = c.Execute(os.Args[1])

	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
