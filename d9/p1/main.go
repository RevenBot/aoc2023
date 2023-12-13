package main

import (
	"fmt"
	"strconv"
	"ac2023.com/m/v2/utils"
)

func sum(array []int)bool{
	for _,v := range array {
		if v != 0 {
			return false
		}
	}
	return true
}

func getLast(array []int) int  {
	newArray := []int{}
	for i := 0; i < len(array)-1; i++ {
		val := array[i+1]-array[i]
		newArray = append(newArray, val)
	}
	i := len(array)
	if sum(newArray){
		fmt.Println(array)
		return array[i-1]
	}else{
		return array[i-1] + getLast(newArray)
	}
}

func main() {
	scanner, file := utils.GetFileScanner("input.txt")


	sum := 0
	for scanner.Scan(){
		aInt := utils.SplitStringToIntArray(scanner.Text())

		val := getLast(aInt)
		utils.PrintInfoMethod("Val last to sum",map[string]string{
			"val":strconv.Itoa(val),
		})
		sum += val
	}
	fmt.Println(sum)

	file.Close()
}
