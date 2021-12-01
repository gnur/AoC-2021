package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := os.Getenv("IN")
	if in == "" {
		in = "input"
	}
	file, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
