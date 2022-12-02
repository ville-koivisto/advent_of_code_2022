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

var getWinningShape = map[byte]byte{
	'A' : 'B',
	'B' : 'C',
	'C' : 'A',
}

var getLosingShape = map[byte]byte{
	'A' : 'C',
	'B' : 'A',
	'C' : 'B',
}

func main() {
	f, err := os.Open("game.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	normalScores := []int{}
	strategyScores := []int{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		game := s.Text()
		opponent, me := game[0], game[2]
		normalScores = append(normalScores, playNormal(opponent, me))
		strategyScores = append(strategyScores, playStrategy(opponent, me))
	}
	
	totalNormal, totalStrategy := 0, 0
	for _, s := range normalScores {
		totalNormal += s
	}
	for _, s := range strategyScores {
		totalStrategy += s
	}
	fmt.Println(totalNormal, totalStrategy)
}

func playNormal(o byte, m byte) int {
	ov, mv := shapeToValue(o), shapeToValue(m)
	if ov-mv == -2 || ov-mv == 1 {
		return mv
	} else if ov-mv == -1 || ov-mv == 2 {
		return mv+6
	} else { return mv+3 }
}

func playStrategy(o byte, m byte) int {	
	if string(m) == "X" {
		losingShape := getLosingShape[o]
		return shapeToValue(losingShape)
	} else if string(m) == "Y" {
		return shapeToValue(o)+3
	} else {
		winningShape := getWinningShape[o]
		return shapeToValue(winningShape)+6
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
