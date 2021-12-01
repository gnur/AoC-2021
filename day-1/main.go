package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var depths []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, v)
	}
	lastDepth := 0
	decreases := 0

	for _, d := range depths {
		if lastDepth == 0 {
			lastDepth = d
			continue
		}
		if lastDepth < d {
			decreases++
		}
		lastDepth = d
	}
	fmt.Println("Decreases:", decreases)

	decreases = 0
	lastSlidingDepth := 0
	for i := range depths {
		if i+2 >= len(depths) {
			break
		}
		d := depths[i] + depths[i+1] + depths[i+2]
		if lastSlidingDepth == 0 {
			lastSlidingDepth = d
			continue
		}
		if lastSlidingDepth < d {
			decreases++
		}
		lastSlidingDepth = d
	}
	fmt.Println("Sliding decreases:", decreases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
