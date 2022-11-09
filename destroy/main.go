package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatal("One argument expected")
	}

	err := os.Remove(args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Destroyed file: %q \n", args[1])
}
