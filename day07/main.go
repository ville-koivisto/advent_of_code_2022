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
	"strings"
)

func main() {
	rootDir := "/"
	parentDir := ""
	currentDir :=""
	dirTree := make(map[string][]int)
	
	i, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer i.Close()

	s := bufio.NewScanner(i)
	for s.Scan() {
		line := s.Text()
		cd := changeDir(line)
		if cd != "" {
			// TODO: continue from here
		}

	}
}

func changeDir(cmd string ) string {
	if strings.Contains(cmd, "$ cd") {
		parseCmd := strings.Fields(cmd)
		moveTo := parseCmd[2]
		return moveTo
	}
	return ""
}
