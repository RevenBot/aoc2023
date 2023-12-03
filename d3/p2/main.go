package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func getNumber(x int, y int, array [][]rune) (int, error) {

	charLeft := array[y][x]

	if charLeft == '.' {
		return 0, errors.New("not")
	}

	number := string(charLeft)

	// printInfoMethod("GetNumber start", map[string]string{
	// 	"chart at x, y": number,
	// 	"y":             strconv.Itoa(y),
	// 	"x":             strconv.Itoa(x),
	// })

	left := x - 1

	for left >= 0 && unicode.IsNumber(array[y][left]) {

		fmt.Println(string(array[y][left]))

		number = string(array[y][left]) + number

		left--
	}

	right := x + 1

	for right < len(array[y]) && unicode.IsNumber(array[y][right]) {

		fmt.Println(string(array[y][right]))
		number = number + string(array[y][right])

		right++
	}
	printInfoMethod("GetNumber after right", map[string]string{
		"charToSUM": number,
	})
	strings.ReplaceAll(number, ".", "")
	return strconv.Atoi(number)
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

	for _, y := range rowPosition {

		if y > 0 && y < len(array) {

			for _, x := range positionsToCheck {

				if x > 0 && x < rowL {
					toCheck := array[y][x]
					printInfoMethod("Check Symbol", map[string]string{"a": string(toCheck),
						"rowIndex": strconv.Itoa(rowIndex),
						"y":        strconv.Itoa(y),
						"x":        strconv.Itoa(x),
					})
					if unicode.IsDigit(toCheck) {

						// num, err := getNumber(x,y,array)

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

func getSymbolResult(positionX int, positionY int, array [][]rune) int {
	// printInfoMethod("GetSymbolResult", map[string]string{
	// 	"*":         string(array[positionY][positionX]),
	// 	"PositionX": strconv.Itoa(positionX),
	// 	"PositionY": strconv.Itoa(positionY),
	// })
	sum := 1
	count := 0

	numberTopCenter, errTC := getNumber(positionX, positionY-1, array)

	if errTC == nil {
		count++
		sum *= numberTopCenter
	} else {
		numberTopLeft, errTL := getNumber(positionX-1, positionY-1, array)
		numberTopRight, errTR := getNumber(positionX+1, positionY-1, array)
		if errTL == nil {
			count++
			sum *= numberTopLeft
			fmt.Println("---- tc ----")
		}
		if errTR == nil {
			count++
			sum *= numberTopRight
		}
		printInfoMethod("GetSymbolResult", map[string]string{
			"TL": strconv.Itoa(numberTopLeft),
			"TR": strconv.Itoa(numberTopRight),
		})
	}
	numberBottomCenter, errBC := getNumber(positionX, positionY+1, array)

	if errBC == nil {
		count++
		sum *= numberBottomCenter
		fmt.Println("---- bc ----")
	} else {
		numberBottomLeft, errBL := getNumber(positionX-1, positionY+1, array)
		numberBottomRight, errBR := getNumber(positionX+1, positionY+1, array)

		printInfoMethod("GetSymbolResult", map[string]string{
			"BL": strconv.Itoa(numberBottomLeft),
			"BR": strconv.Itoa(numberBottomRight),
		})
		if errBL == nil {
			count++
			sum *= numberBottomLeft
		}
		if errBR == nil {
			sum *= numberBottomRight
			count++
		}
	}

	numberLeft, errL := getNumber(positionX-1, positionY, array)
	numberRight, errR := getNumber(positionX+1, positionY, array)
	printInfoMethod("GetSymbolResult", map[string]string{
		"L": strconv.Itoa(numberLeft),
		"R": strconv.Itoa(numberRight),
	})

	if errL == nil {
		count++
		sum *= numberLeft
	}
	if errR == nil {
		count++
		sum *= numberRight
	}
	if sum == 1 {
		return 0
	}
	if count > 1 {
		printInfoMethod("GetSymbolResult", map[string]string{
			"Valid sum of *": strconv.Itoa(sum),
			"Count":          strconv.Itoa(count),
		})
		return sum
	}
	return 0
}

func main() {

	scanner, file := getFileScanner("input.txt")
	lines, lineLenght := getDimMultiArray(*scanner)
	array := createMultiArray(lines, lineLenght)
	file.Close()

	scanner1, file1 := getFileScanner("input.txt")

	popolateArray(*scanner1, *array)

	sum := 0

	for y, row := range *array {
		for x, val := range row {
			if val == '*' {
				fmt.Println()
				printInfoMethod("Found *", map[string]string{"index": strconv.Itoa(x)})
				sum += getSymbolResult(x, y, *array)
				printInfoMethod("Found *", map[string]string{
					"sum": strconv.Itoa(sum),
				})
				fmt.Println()
			}

		}
	}

	fmt.Println(sum)

	file1.Close()
}
