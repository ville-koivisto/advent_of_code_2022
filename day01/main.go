/*
Advent of Code: Day 1

Check Elf's inventory and find the one
carrying most calories. Return max calories
(part 1). Also find total sum for the best
of three (part 2).

The inventory is written in "inventory.txt".
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main () {
	// read the inventory file
	i, err := os.Open("inventory.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer i.Close()

	// read lines
	scanner := bufio.NewScanner(i)
	sum := 0
	inventory := []int{}
	
	// sum inventory items per Elf
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			intLine, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			sum += intLine
		} else {
			inventory = append(inventory, sum)
			sum = 0
		}
	}

	// find max calories (resolves Part 1 / Day 1)
	max := 0
	for _, c := range inventory {
		if c > max {
			max = c
		}
	}

	// sum of best three (resolves Part 2 / Day 1)
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	maxThree := 0
	for i := 0; i < 3; i++ {
		maxThree += inventory[i]
	}

	fmt.Println(max)
	fmt.Println(maxThree)
}
