package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func countWinningNumbers(textW []string, textN []string) int {

	count := 0
	for _, wC := range textW {
		for _, nC := range textN {
			if wC != "" && wC != " " && wC == nC {
				fmt.Println(nC)
				count++
			}
		}
	}
	fmt.Println("***")
	fmt.Println(count)
	return count
}

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

func main() {
	toSum := 0
	scanner, file := getFileScanner("input.txt")
	for scanner.Scan() {
		numbers := scanner.Text()[8:]
		numberStrings := strings.Split(numbers, "|")
		winningNumbers := strings.Split(numberStrings[0], " ")
		myNumbers := strings.Split(numberStrings[1], " ")
		//   for _,wC := range(winningNumbers){
		//   fmt.Println(wC)
		//   fmt.Println(len(winningNumbers))
		//   }
	  winningNumbersinRow := countWinningNumbers(winningNumbers,myNumbers)
  
	  toSum += int(math.Pow(2, float64(winningNumbersinRow)-1))

	}

	fmt.Println("-----")
	fmt.Println(toSum)

	file.Close()

}
