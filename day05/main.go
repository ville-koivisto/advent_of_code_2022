/*
Advent of Code: Day 5

After the crane has made the moves, find out what crates
end up on top of each of the three stacks (part 1).

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	stacks := [][]byte{
		{'S','C','V','N'},
		{'Z','M','J','H','N','S'},
		{'M','C','T','G','J','N','D'},
		{'T','D','F','J','W','R','M'},
		{'P','F','H'},
		{'C','T','Z','H','J'},
		{'D','P','R','Q','F','S','L','Z'},
		{'C','S','L','H','D','F','P','W'},
		{'D','S','M','P','F','N','G','Z'},
	}

	m, err := os.Open("moves.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer m.Close()

	s := bufio.NewScanner(m)
	for s.Scan() {
		splitMoves := strings.Fields(s.Text())
		move, origin, destination := splitMoves[1], splitMoves[3], splitMoves[5]
		m, _ := strconv.Atoi(move)
		o, _ := strconv.Atoi(origin)
		d, _ := strconv.Atoi(destination)
		stacks = addToDestination(stacks, m, o, d)
	}
	
	// part 1
	getLetters(stacks)
}

func addToDestination(s [][]byte, m int, o int, d int) [][]byte {
	orig, dest := o-1, d-1
	for i := range s[orig] {
		if len(s[orig])-1-i > len(s[orig])-1-m {
			s[dest] = append(s[dest], s[orig][len(s[orig])-1-i])
		}
	}
	s = removeFromOrigin(s, m, o)
	return s
}

func removeFromOrigin(s [][]byte, m int, o int) [][]byte {
	orig := o-1
	s[orig] = s[orig][:len(s[orig])-m]
	return s
}

func getLetters(s [][]byte) {
	for i := range s {
		c := string(s[i][len(s[i])-1])
		fmt.Print(c)
	}
	fmt.Println()
}
