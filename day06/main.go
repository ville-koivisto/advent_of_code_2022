/*
Advent of Code: Day 6

Find combination of four different characters, the
start-of-packet marker, from the datastream (part 1).
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	d, err := os.Open("datastream.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer d.Close()

	s := bufio.NewScanner(d)
	for s.Scan() {
		data := s.Text()
		for i := range data {
			if isUnique(data, i) {
				fmt.Printf("Start marker end found at %vth character", i+4)
				fmt.Println()
				os.Exit(0)
			}
		}
	}
}

func isUnique(data string, index int) bool {
	checkUnique := make(map[byte]bool,0)
	for i := index; i <= index+3; i++ {
		if checkUnique[data[i]] {
			return false
		}
		checkUnique[data[i]] = true
	}
	return true
}
