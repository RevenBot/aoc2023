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



func stringIsNumber(text string) (int, error){
	if strings.Contains(text,"one") {
		return 1,nil
	}
	if strings.Contains(text,"two") {
		return 2,nil
	}
	if strings.Contains(text,"three") {
		return 3,nil
	}
	if strings.Contains(text,"four") {
		return 4,nil
	}
	if strings.Contains(text,"five") {
		return 5,nil
	}
	if strings.Contains(text,"six") {
		return 6,nil
	}
	if strings.Contains(text,"seven") {
		return 7,nil
	}
	if strings.Contains(text,"eight") {
		return 8,nil
	}
	if strings.Contains(text,"nine") {
		return 9,nil
	}
	return 0, errors.New("not")
}

func numberFromString(text string) (int,error) {
	fmt.Println(text)
	number, err := stringIsNumber(text)
	
	if err == nil{
		return number ,nil
	}
	return 0 , err
}



func getFirstDigit(text string) int  {
	
	var line string

	for _, char := range text {

		if unicode.IsDigit(char) {
			return int(char-'0')
		}else {
			line += string(char)
			number, err := numberFromString(line)
			if err == nil{
				fmt.Println(number)
				return number
			}
		}
	}
	return 0
}

func getLastDigit(text string) int {

	runes := []rune(text)
	var line string

	for x := len(runes) - 1; x >= 0; x-- {
		if unicode.IsDigit(runes[x]) {
			return int(runes[x]-'0')
		}else{
			line = string(runes[x]) + line
			number, err := numberFromString(line)
			if err == nil{
				return number
			}
		}

	}

	return 0
}

func getSum(file bufio.Scanner) (int32, error) {

	var sum int = 0

	file.Split(bufio.ScanLines)
	for file.Scan() {

		text := file.Text()

		fmt.Printf("----- %v ------\n", text)

		

		first := strconv.Itoa(getFirstDigit(text))
		last := strconv.Itoa(getLastDigit(text))

		fmt.Printf("-----sum: %v + %v ------\n", first, last)

		fullNumber := first+last

		fmt.Printf("-----sum: %v ------\n",fullNumber)
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
