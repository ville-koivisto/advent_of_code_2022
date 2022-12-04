/*
Advent of Code: Day 3

- Rucksack has two compartments
- Each type in only one compartment
-

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	items := fillItems()
	// priorities := []int{}
	fmt.Println(items)
	fmt.Println(getPriority('A', items))

	r, err := os.Open("rucksacks.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	for s.Scan() {
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
		fmt.Println("Rucksack:", comp1, comp2)
	}

	/* TODO:
	- lue rivi
	- jaa kahteen yhtä suureen osaan
	- talleta osat omiin slice-rakenteisiin
	- hae intersektio
	- har intersektion arvolle priority
	- summaa priority muuttujaan
	*/
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

func getPriority(item byte, items []byte) (int, error) {
	for i, v := range items {
		if v == item {
			return i+1, nil
		}
	}
	return 0, fmt.Errorf("item not in rucksack")
}