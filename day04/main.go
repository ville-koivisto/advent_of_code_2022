/*
Advent of Code: Day 4

Check how many of the assignments are a subset of eachother (part 1)
and how many of them overlap at all (part 2).

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
	var subsets, overlaps int
	assigments, err := os.Open("assignment_pairs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer assigments.Close()

	s := bufio.NewScanner(assigments)
	for s.Scan() {
		rawPair := s.Text()

		// part 1
		elf1, elf2 := getRanges(rawPair)
		if isSubset(elf1, elf2) || isSubset(elf2, elf1) {
			subsets += 1
		}

		// part 2
		if overlap(elf1, elf2) {
			overlaps += 1
		}
	}
	fmt.Println(subsets)
	fmt.Println(overlaps)
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

func overlap(a, b []int) bool {
	for _, v1 := range a {
		for _, v2 := range b {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}
