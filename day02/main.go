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

var getWinner = map[byte]byte{
	'A' : 'B',
	'B' : 'C',
	'C' : 'A',
}

var getLoser = map[byte]byte{
	'A' : 'C',
	'B' : 'A',
	'C' : 'B',
}

func main() {
	g, err := os.Open("game.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	normalScores := []int{}
	strategyScores := []int{}
	
	s := bufio.NewScanner(g)
	for s.Scan() {
		game := s.Text()
		o, p := game[0], game[2]
		normalScores = append(normalScores, playNormal(o, p))
		strategyScores = append(strategyScores, playStrategy(o, p))
	}

	totalNormal, totalStrategy := 0, 0
	for _, p := range normalScores {
		totalNormal += p
	}
	for _, p := range strategyScores {
		totalStrategy += p
	}

	fmt.Println(totalNormal, totalStrategy)
}

func playNormal(o byte, p byte) int {
	op, me := shapeToValue(o), shapeToValue(p)
	if op-me == -2 || op-me == 1 {
		return me
	} else if op-me == -1 || op-me == 2 {
		return me + 6
	} else { return me + 3 }
}

func playStrategy(o byte, p byte) int {	
	if string(p) == "X" {
		losingShape := getLoser[o]
		return shapeToValue(losingShape)
	} else if string(p) == "Y" {
		return shapeToValue(o) + 3
	} else {
		winningShape := getWinner[o]
		return shapeToValue(winningShape) + 6
	}
}

func shapeToValue(shape byte) int {
	switch value := string(shape); value {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	return 0
}
