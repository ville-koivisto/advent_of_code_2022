/*
Advent of Code: Day 2

Elfs are playing rock paper scissors and you want to
know the tournament outcome without (part 1) and
with following a strategy guide (part 2).

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	g, err := os.Open("game.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	s := bufio.NewScanner(g)
	pointList := []int{}
	for s.Scan() {
		game := s.Text()
		o, p := game[0], game[2]
		pointList = append(pointList, playGame(o, p))
	}

	total := 0
	for _, p := range pointList {
		total += p
	}

	fmt.Println(total)
}

func playGame(o byte, p byte) int {
	op := shapeToValue(o)
	me := shapeToValue(p)
	
	// lose, win, draw
	if op-me == -2 || op-me == 1 {
		return me
	} else if op-me == -1 || op-me == 2 {
		return me + 6
	} else { return me + 3 }
}

func shapeToValue(m byte) int {
	switch value := string(m); value {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	return 0
}
