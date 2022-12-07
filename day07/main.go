/*
Advent of Code: Day 7

The communication device Elfs gave is full and
files need to be deleted to free up some space.
Find out which dirs total at most 100000 (part 1).
*/

package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	i, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer i.Close()

	s := bufio.NewScanner(i)
	for s.Scan() {
		line := s.Text()
						
	}
}