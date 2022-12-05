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
)

type allStacks struct {
	st1 []byte
	st2 []byte
	st3 []byte
	st4 []byte
	st5 []byte
	st6 []byte
	st7 []byte
	st8 []byte
	st9 []byte
}

var stacks = allStacks{
[]byte{'S','C','V','N'},
[]byte{'Z','M','J','H','N','S'},
[]byte{'M','C','T','G','J','N','D'},
[]byte{'T','D','F','J','W','R','M'},
[]byte{'P','F','H'},
[]byte{'C','T','Z','H','J'},
[]byte{'D','P','R','Q','F','S','L','Z'},
[]byte{'C','S','L','H','D','F','P','W'},
[]byte{'D','S','M','P','F','N','G','Z'},
}

func main() {
	m, err := os.Open("moves.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer m.Close()

	s := bufio.NewScanner(m)
	for s.Scan() {
		rawMove := s.Text()
		fmt.Println(rawMove)
	}
	fmt.Println(stacks)
}

/* TODO:
- parsi syötteestä
	- kuinka monta konttia siirretään
	- lähde
	- kohde
- toteuta algoritmi
	- firt in last out

*/
