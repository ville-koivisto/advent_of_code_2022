/*
Advent of Code: Day 6

Find combination of four different characters, the
start-of-packet marker, from the datastream (part 1).
Also, tune in your device to find start of the
message marker of 14 chars long (part 2).
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
			if isMarker(data, i, 3) {
				fmt.Printf("Start marker end (packet) found at char: %v\n", i+4)
				break
			}
		}
		for i := range data {
			if isMarker(data, i, 13) {
				fmt.Printf("Start marker end (message) found at char: %v\n", i+14)
				os.Exit(0)
			}
		}
	}
}

func isMarker(data string, index int, marker int) bool {
	checkUnique := make(map[byte]bool,0)
	for i := index; i <= index+marker; i++ {
		if checkUnique[data[i]] {
			return false
		}
		checkUnique[data[i]] = true
	}
	return true
}
