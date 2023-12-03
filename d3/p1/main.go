package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		panic(e)
	}
}

func getFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	fmt.Printf("---- START READFILE -- %v ----\n", fileName)
	file, err := os.Open(fileName)
	check(err)
	fmt.Printf("---- END READFILE -- %v ----\n", fileName)
	return bufio.NewScanner(file), file
}

func printInfoMethod(methodInfo string, args map[string]string) {

	fmt.Printf("---- %v ----\n", methodInfo)
	if len(args) > 0 {
		fmt.Printf("---- ARGS: ")
		for k, v := range args {
			fmt.Printf("- %v : %v -", k, v)
		}
		fmt.Printf("\n")
	}
}

func getDimMultiArray(scanner bufio.Scanner) (int, int) {
	countLines, countLineLenght := 0, 0

	printInfoMethod("Start getDimMultiArray", map[string]string{})

	for scanner.Scan() {
		countLines++
		string := scanner.Text()
		countLineLenght = len(string)
	}
	printInfoMethod("End getDimMultiArray", map[string]string{
		"cL":  strconv.Itoa(countLines),
		"cLL": strconv.Itoa(countLineLenght),
	})

	return countLines, countLineLenght

}

func createMultiArray(lines int, lineLenght int) *[][]rune {

	twoDSlice := make([][]rune, lines)

	for i := range twoDSlice {
		twoDSlice[i] = make([]rune, lineLenght)
	}
	return &twoDSlice

}

func popolateArray(scanner bufio.Scanner, multiArray [][]rune) {
	printInfoMethod("Start popolateArray", map[string]string{})
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			multiArray[index][i] = v
		}
		index++
	}
	printInfoMethod("End popolateArray", map[string]string{})
}

func checkSymbol(positions []int, array [][]rune, rowIndex int, rowL int) (int, error) {

	var debug string
	for _, s := range positions {
		debug += string(array[rowIndex][s])
	}

	printInfoMethod(" check Symbol", map[string]string{"number": debug})

	// add -1 + 1
	positionsToCheck := append(positions, positions[0]-1)
	positionsToCheck = append(positionsToCheck, positions[len(positions)-1]+1)

	rowPosition := []int{rowIndex - 1, rowIndex, rowIndex + 1}

	for _, v := range rowPosition {

		if v > 0 && v < len(array) {

			for _, p := range positionsToCheck {

				if p > 0 && p < rowL {
					toCheck := array[v][p]
					printInfoMethod("Check Symbol", map[string]string{"a": string(toCheck),
						"rowIndex": strconv.Itoa(rowIndex),
						"y":        strconv.Itoa(v),
						"x":        strconv.Itoa(p),
					})
					if !unicode.IsDigit(toCheck) && toCheck != '.' {

						var text string
						for _, s := range positions {
							text += string(array[rowIndex][s])
						}
						printInfoMethod("Check Symbol", map[string]string{"a": string(toCheck),
							"ris": strconv.Itoa(rowIndex),
						})
						return strconv.Atoi(text)

					}

				}

			}

		}

	}

	return 0, errors.New("not")

}

func main() {

	scanner, file := getFileScanner("input.txt")
	lines, lineLenght := getDimMultiArray(*scanner)
	array := createMultiArray(lines, lineLenght)
	file.Close()

	scanner1, file1 := getFileScanner("input.txt")

	popolateArray(*scanner1, *array)

	sum := 0

	for rowIndex, row := range *array {
		var positions []int
		for j, val := range row {
			if unicode.IsDigit(val) {
				positions = append(positions, j)
			} else {

				if len(positions) > 0 {
					num, err := checkSymbol(positions, *array, rowIndex, len(row))
					if err == nil {
						printInfoMethod("NUM TO SUM", map[string]string{"num": string(num)})
						sum += num
					}
					positions = []int{}
				}
			}
			if j+1 == len(row) {
				if len(positions) > 0 {
					num, err := checkSymbol(positions, *array, rowIndex, len(row))
					if err == nil {
						printInfoMethod("NUM TO SUM", map[string]string{"num": string(num)})
						sum += num
					}
				}
			}
		}
	}

	fmt.Println(sum)

	file1.Close()
}
