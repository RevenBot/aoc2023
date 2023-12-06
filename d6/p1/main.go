package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// type   struct {
// 	number     int
// 	converters []mapConverter
// }

func main() {
	scanner, file := getFileScanner("input.txt")

	var times []int
	var distance []int

	for scanner.Scan(){

		text := scanner.Text()
		fmt.Println(text)
		if strings.Contains(text, "Time"){
			temp := (strings.Split(text," "))
			for _, v := range temp {
				if v != "" && v != " "{
					val, err := strconv.Atoi(v)
					if err == nil{
						times = append(times, val)
					}
				}
			}
		}
		if strings.Contains(text, "Distance"){
			temp := (strings.Split(text," "))
			for _, v := range temp {
				if v != "" && v != " "{
					val, err := strconv.Atoi(v)
					if err == nil{
						distance = append(distance, val)
					}
				}
			}
		}
	}

	sum := 1

	for i, time := range times {
		count := 0
		for j:=1; j < time; j++{
			fmt.Println(j*(time-j))
			fmt.Println(distance[i])
			fmt.Println(j*(time-j) >= distance[i])

			if j*(time-j) >= distance[i] {
				count++

			}

		}

		sum *=count
		
		
	}

	fmt.Println("*****")
	fmt.Println(sum)




	file.Close()
	}

