/*
Advent of Code: Day 3

Get the sum of priority values (part 1).
Get the sum of badge priority values (per group of three Elfs) (part 2).

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/juliangruber/go-intersect"
)

func main() {
	items := fillItems()
	prioritySum, badgeSum := 0, 0
	elfGroup := []string{}

	r, err := os.Open("rucksacks.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	for s.Scan() {
		// part 1
		comp1, comp2 := []rune{}, []rune{}
		rucksack := s.Text()
		cap := len(rucksack)
		for i, item := range rucksack {
			if i <= cap/2-1 {
				comp1 = append(comp1, item)
			} else {
				comp2 = append(comp2, item)
			}
		}
		c := fmt.Sprintf("%v", intersect.Hash(comp1, comp2)[0])
		value, err := getPriorityValue(c, items)
		if err != nil {
			log.Fatal(err)
		}
		prioritySum += value

		// part 2
		elfGroup = append(elfGroup, rucksack)
		if len(elfGroup) == 3 {
			badge := fmt.Sprintf("%v", intersect.Hash(elfGroup[0], intersect.Hash(elfGroup[1], elfGroup[2]))[0])
			value, err := getPriorityValue(badge, items)
			if err != nil {
				log.Fatal(err)
			}
			badgeSum += value
			elfGroup = []string{}
		}
	}

	fmt.Println(prioritySum)
	fmt.Println(badgeSum)
}

func fillItems() []byte {
	items := []byte{}
	var ch byte
	for ch = 'a'; ch <= 'z'; ch++ {
		items = append(items, ch)
	}
	for ch = 'A'; ch <= 'Z'; ch++ {
		items = append(items, ch)
	}
	return items
}

func getPriorityValue(item string, items []byte) (int, error) {
	for i, v := range items {
		if item == fmt.Sprintf("%v",v) {
			return i+1, nil
		}
	}
	return 0, fmt.Errorf("invalid value for an item")
}
