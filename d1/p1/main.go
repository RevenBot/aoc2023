package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func getFirstDigit(text string) rune  {
	for _, char := range text {

		if unicode.IsDigit(char) {
			return char
		}

	}
	return '0'
}

func getLastDigit(text string) rune {

	runes := []rune(text)

	for x := len(runes) - 1; x >= 0; x-- {

		if unicode.IsDigit(runes[x]) {
			return runes[x]
		}

	}

	return '0'
}

func getSum(file bufio.Scanner) (int32, error) {

	var sum int = 0

	file.Split(bufio.ScanLines)
	for file.Scan() {
		first := strconv.Itoa(int(getFirstDigit(file.Text())-'0'))
		last := strconv.Itoa(int(getLastDigit(file.Text())-'0'))

		fullNumber := first+last

		num, err := strconv.Atoi(fullNumber)
		check(err)

		sum += num
	}

	return int32(sum), nil

}

func main() {
	file, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(file)

	sum, err := getSum(*fileScanner)

	fmt.Println(sum)

	file.Close()

}
