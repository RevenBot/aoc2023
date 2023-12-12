package main

import (
	"bufio"
	"fmt"

	"ac2023.com/m/v2/utils"
)

func getDataStructure(scanner *bufio.Scanner) *map[string][]string {

	args := map[string][]string{}
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
		if txt != "" {
			key := txt[0:3]
			left := txt[7:10]
			right := txt[12:15]
			args[key] = []string{left, right}
		}

	}
	return &args
}

func countSteps(steps string, mapToFollow map[string][]string, search string) int {

	key := search
	count := 0
	fmt.Println(mapToFollow)
	for _, v := range steps {
		if string(v) == "L" {
			fmt.Println(mapToFollow[key])
			key = mapToFollow[key][0]
		} else {
			fmt.Println(mapToFollow[key])
			key = mapToFollow[key][1]
		}
		count++
		if key == "ZZZ" {
			return count
		}

	}
	return count + countSteps(steps,mapToFollow,key)
}

func main() {
	scanner, file := utils.GetFileScanner("input.txt")
	scanner.Scan()
	firstRow := scanner.Text()
	args := getDataStructure(scanner)

	count := countSteps(firstRow, *args, "AAA")

	fmt.Println(count)

	file.Close()
}
