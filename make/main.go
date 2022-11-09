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

	f, err := os.Create(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Printf("Created file: %q \n", args[1])
}
