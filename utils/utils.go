package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func GetFileScanner(fileName string) (*bufio.Scanner, *os.File) {
	fmt.Printf("---- START READFILE -- %v ----\n", fileName)
	file, err := os.Open(fileName)
	Check(err)
	fmt.Printf("---- END READFILE -- %v ----\n", fileName)
	return bufio.NewScanner(file), file
}

func PrintInfoMethod(methodInfo string, args map[string]string) {
	fmt.Printf("---- %v ----\n", methodInfo)
	if len(args) > 0 {
		fmt.Printf("---- ARGS: ")
		for k, v := range args {
			fmt.Printf("- %v : %v -", k, v)
		}
		fmt.Printf("\n")
	}
}

func SplitStringToIntArray(text string) []int{

	arrayString := strings.Split(text," ")

	var arrayInt []int

	for _, v := range arrayString {

		if  v != " "{
		val , err := strconv.Atoi(v)
		if err == nil  {
				arrayInt = append(arrayInt,val)
			}
		}

	}

	return arrayInt

}

