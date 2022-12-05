/*
Advent of Code: Day 5

After the crane has made the moves, find out what crates
end up on top of each of the three stacks (part 1). Do
the same drill but with CrateMover 9001 (part 2).

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
	part1, part2 := getStacks(), getStacks()
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

		// part 1
		part1 = addAndRemovePart1(part1, m, o, d)
		
		// part 2
		part2 = addAndRemovePart2(part2, m, o, d)
	}
	getLetters(part1)
	getLetters(part2)
}

func getStacks() [][]byte {
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
	return stacks
}

func addAndRemovePart1(s [][]byte, m int, o int, d int) [][]byte {
	orig, dest := o-1, d-1
	for i := range s[orig] {
		if len(s[orig])-1-i > len(s[orig])-1-m {
			s[dest] = append(s[dest], s[orig][len(s[orig])-1-i])
		}
	}
	s = removeFromOrigin(s, m, o)
	return s
}

func addAndRemovePart2(s [][]byte, m int, o int, d int) [][]byte {
	orig, dest := o-1, d-1
	if m == 1 {
		s[dest] = append(s[dest], s[orig][len(s[orig])-1])
	} else {
		crates := s[orig][len(s[orig])-m:]
		s[dest] = append(s[dest], crates...)
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
