/*
Advent of Code: Day 5

After the crane has made the moves, find out what crates
end up on top of each of the three stacks (part 1).

*/

package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	
	m, err := os.Open("moves.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer m.Close()

	s := bufio.NewScanner(m)
	for s.Scan() {
		// TODO: all
	}
}