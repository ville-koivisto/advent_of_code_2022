/*
Advent of Code: Day 1

Check Elf's inventory and find out the one
carrying most calories. Inventory is written
in "inventory.txt".
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// find max calories (this resolves Part 1 of the Day 1)
	max := 0
	for _, c := range inventory {
		if c > max {
			max = c
		}
	}

	fmt.Println(max)
}
