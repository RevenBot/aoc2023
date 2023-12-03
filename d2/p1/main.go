package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		panic(e)
	}
}
func getFileScanner(fileName string) *bufio.Scanner {
	fmt.Printf("---- START READFILE -- %v ----\n", fileName)
	file, err := os.Open(fileName)
	check(err)
	fmt.Printf("---- END READFILE -- %v ----\n", fileName)
	return bufio.NewScanner(file)
}
func getIndex(indexString string) int {
	// fmt.Printf("---- START getIndex -- %v ----\n", indexString)

	re := regexp.MustCompile("[0-9]+")
	number, err := strconv.Atoi(re.FindString(indexString))
	check(err)

	// fmt.Printf("---- END getIndex -- res string: %v -- number: %v ----\n", indexString, number)
	return number
}
func getGames(gamesString string) []string {
	fmt.Printf("---- START getGames -- %v ----\n", gamesString)

	games := strings.Split(gamesString, ";")
	fmt.Printf("---- END getGames -- %v ----\n", games)
	return games
}
func checkColor(colors []string) (int, int, int) {
	fmt.Printf("---- START checkColor -- %v ----\n", colors)
	cRed,cGreen, cBlue := 0,0,0
	for i := 0; i < len(colors); i++ {
		if strings.Contains(colors[i],"red"){
			cRed += getIndex(colors[i])
		}
		if strings.Contains(colors[i],"green"){
			cGreen += getIndex(colors[i])
		}
		if strings.Contains(colors[i],"blue"){
			cBlue += getIndex(colors[i])
		}
		
	}
	return cRed, cGreen, cBlue
	
}
func checkGames(index int, games []string) int  {
	fmt.Printf("---- START checkGames -- %v ----\n", index)


	for i := 0; i < len(games); i++ {

		red, green, blue := checkColor(strings.Split(games[i], ","))
		if red > 12 {
			return 0
		}
		if green > 13 {
			return 0
		}
		if blue > 14 {
			return 0
		}

	}



	fmt.Printf("---- END getGames -- %v ----\n", index)
	return index
}

func getData(text string) int {
	fmt.Printf("---- START getData -- %v ----\n", text)
	pos := strings.IndexByte(text, ':')
	index := getIndex(text[:pos])
	games := getGames(text[pos+1:])
	res := checkGames(index, games)
	fmt.Printf("---- END getData -- Index: %v --- res: %v ----\n\n",index,res)
	return res
	
}

func main() {
	fmt.Printf("---- START MAIN ----\n")
	scanner := getFileScanner("input.txt")
	sum:=0
	for scanner.Scan() {
		sum += getData(scanner.Text())
	}

	fmt.Printf("---- END MAIN %v ----\n",sum)
}
