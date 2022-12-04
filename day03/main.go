/*
Advent of Code: Day 3

- Rucksack has two compartments
- Each type in only one compartment
-

*/

package main

import "fmt"

func main() {
	items := fillItems()
	fmt.Println(items)
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

func getPriority(item byte, itemlist []byte) (int, error) {
	for i, value := range itemlist {
		if value == item {
			return i+1, nil
		}
	}
	return 0, fmt.Errorf("item not in rucksack")
}