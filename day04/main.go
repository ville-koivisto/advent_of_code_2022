/*
Advent of Code: Day 4

Check how many pairs overlap fully (part 1).

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	countSubsets := 0
	assigments, err := os.Open("assignment_pairs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer assigments.Close()

	s := bufio.NewScanner(assigments)
	for s.Scan() {
		rawPair := s.Text()
		elf1, elf2 := getRanges(rawPair)
		if isSubset(elf1, elf2) || isSubset(elf2, elf1) {
			countSubsets += 1
		}
	}
	fmt.Println(countSubsets)
}


func getRanges(rawPair string) ([]int, []int) {
	var ranges, elf1, elf2 []int
	s, err := regexp.Compile(`-|,`)
	if err != nil {
		log.Fatal(err)
	}
	values := s.Split(rawPair, 4)
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, num)
	}
	for i := ranges[0]; i <= ranges[1]; i++ {
		elf1 = append(elf1, i)
	}
	for i := ranges[2]; i <= ranges[3]; i++ {
		elf2 = append(elf2, i)
	}
	return elf1, elf2
}

func isSubset(first, second []int) bool {
    set := make(map[int]int)
    for _, value := range second {
        set[value] += 1
    }

    for _, value := range first {
        if count, found := set[value]; !found {
            return false
        } else if count < 1 {
            return false
        } else {
            set[value] = count - 1
        }
    }
    return true
}
