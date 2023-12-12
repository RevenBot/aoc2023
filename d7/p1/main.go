package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"ac2023.com/m/v2/utils"
)

type Hand struct {
	cards    []string
	bindHand int
}

type SortHand []Hand

func (h SortHand) Len() int {
	return len(h)
}

func (u SortHand) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u SortHand) Less(i, j int) bool {
	return sortByLetterCount(u[i], u[j])
}

func sortByLetterCount(left Hand, right Hand) bool {
	countLeft := countCard(left)
	countRight := countCard(right)


	if countLeft == countRight {
		return sortByOrder(left, right)
	}
	return countLeft < countRight
}

func sortByOrder(left Hand, right Hand) bool {

	for i := 0; i < len(left.cards); i++ {
		sLeft := getStrength(left.cards[i])
		sRight := getStrength(right.cards[i])
		if sLeft != sRight {
			return sLeft < sRight
		}
	}
	return false
}
func getStrength(s string) int {
	strengthArray := []string{"2","3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	for i := 0; i < len(strengthArray); i++ {
		if s == strengthArray[i] {
			return i + 1
		}
	}
	return 0
}

func countCard(hand Hand) int {
	args := map[string]int {}
	for i := 0; i < len(hand.cards); i++ {
		args[hand.cards[i]]+=1
	}
	l := len(args)
	if l == 1 {
		return 7
	}
	max := 0
	for _,v := range args {
		if v > max {
			max = v
		}
	}
	if l == 2 && max == 4 {
		return 6
	}
	if l == 2 && max == 3 {
		return 5
	}
	if l == 3 && max == 3 {
		return 4
	}
	if l == 3 && max == 2 {
		return 3
	}
	if l == 4 && max == 2 {
		return 2
	}
	return 1
}

type WinningHand struct {
	strenghHand int
	order       []int
}

func GetHands(scannner *bufio.Scanner) []Hand {
	var arrayToReturn []Hand
	for scannner.Scan() { line := scannner.Text()
		handsString := strings.Split(line, " ")
		val, err := strconv.Atoi(handsString[1])
		if err == nil {
			newHand := &Hand{
				cards:    strings.Split(handsString[0], ""),
				bindHand: val,
			}
			arrayToReturn = append(arrayToReturn, *newHand)
		}

	}
	return arrayToReturn
}

func calculateTotalWinning(hands []Hand) int {
	sum := 0
	for i, v := range hands {
		sum += (i + 1) * v.bindHand
		fmt.Printf("=%v= %v * %v = %v tot=%v\n",v,i+1,v.bindHand,(i + 1) * v.bindHand,sum)
	}
	return sum
}

func main() {

	scanner, file := utils.GetFileScanner("input.txt")

	hands := GetHands(scanner)
	fmt.Println(hands)
	fmt.Println("***")
	fmt.Println("***")
	sort.Sort(SortHand(hands))
	fmt.Println(calculateTotalWinning(hands))
	file.Close()
}
