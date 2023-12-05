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

func getSeeds(scan *bufio.Scanner) []int {

	scan.Scan()
	firstRowText := scan.Text()[6:]
	seedsText := strings.Split(firstRowText, " ")
	var seedNumbers []int
	for _, v := range seedsText {
		if v != "" {
			n, err := strconv.Atoi(v)
			if err != nil {
				check(err)
			}
			seedNumbers = append(seedNumbers, n)
		}
	}
	return seedNumbers

}

type Converter interface {
	GetNewData(input int) int
}

type soilConverter struct {
	number     int
	converters []mapConverter
}
type fertilizerConverter struct {
	number     int
	converters []mapConverter
}
type waterConverter struct {
	number     int
	converters []mapConverter
}
type lightConverter struct {
	number     int
	converters []mapConverter
}
type temperatureConverter struct {
	number     int
	converters []mapConverter
}
type humidityConverter struct {
	number     int
	converters []mapConverter
}
type locationConverter struct {
	number     int
	converters []mapConverter
}

func (soil soilConverter) GetNewData(input int) int {
	printInfoMethod("SoilDec", map[string]string{
		"seed": strconv.Itoa(input),
	})
	var val = input
	for _, v := range soil.converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("SoilDec", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("SoilDec", map[string]string{
				"after-val": strconv.Itoa(val),
			})
			return val
		}
	}
	return val
}
func (fertilizer fertilizerConverter) GetNewData(input int) int {
	printInfoMethod("fertilizer", map[string]string{
		"seed": strconv.Itoa(input),
	})
	var val = input
	for _, v := range fertilizer.converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("fertilizer", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("fertilizer", map[string]string{
				"after-val": strconv.Itoa(val),
			})
			return val
		}
	}
	return val
}
func (water waterConverter) GetNewData(input int) int {
	printInfoMethod("water", map[string]string{
		"seed": strconv.Itoa(input),
	})
	var val = input
	for _, v := range water.converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("water", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("water", map[string]string{
				"after-val": strconv.Itoa(val),
			})
			return val
		}
	}
	return val
}
func (light lightConverter) GetNewData(input int) int {
	printInfoMethod("light", map[string]string{
		"seed": strconv.Itoa(input),
	})
	var val = input
	for _, v := range light.converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("light", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("light", map[string]string{
				"after-val": strconv.Itoa(val),
			})
			return val
		}
	}
	return val
}
func (temperature temperatureConverter) GetNewData(input int) int {
	printInfoMethod("temperature", map[string]string{
		"seed": strconv.Itoa(input),
	})
	var val = input
	for _, v := range temperature.converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("temperature", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("temperature", map[string]string{
				"after-val": strconv.Itoa(val),
			})
			return val
		}
	}
	return val
}
func (humidity humidityConverter) GetNewData(input int) int {
	printInfoMethod("humidity", map[string]string{
		"seed": strconv.Itoa(input),
	})
	var val = input
	for _, v := range humidity.converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("humidity", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("humidity", map[string]string{
				"after-val": strconv.Itoa(val),
			})
			return val
		}
	}
	return val
}
func (location locationConverter) GetNewData(input int) int {
	printInfoMethod("location", map[string]string{
		"seed": strconv.Itoa(input),
	})
	var val = input
	for _, v := range location.converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("location", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("location", map[string]string{
				"after-val": strconv.Itoa(val),
			})
			return val
		}
	}
	return val
}

type mapConverter struct {
	index       int
	name        string
	destination int
	source      int
	mesureRange int
}

func getMapConverters(scan *bufio.Scanner) (*[]mapConverter, *[]mapConverter, *[]mapConverter, *[]mapConverter, *[]mapConverter, *[]mapConverter, *[]mapConverter) {
	count := 0
	var cSoils []mapConverter
	var cFertilizer []mapConverter
	var cWater []mapConverter
	var cLight []mapConverter
	var cTemperature []mapConverter
	var cHumidity []mapConverter
	var cLocation []mapConverter
	var name string
	for scan.Scan() {
		line := scan.Text()
		if strings.Contains(line, ":") {
			count = 0
			name = line
			fmt.Println("****")
			fmt.Println(name)
		}
		numbers := strings.Split(line, " ")
		fmt.Println(numbers)
		if len(numbers) == 3 {
			destination, errd := strconv.Atoi(numbers[0])
			source, errs := strconv.Atoi(numbers[1])
			mesureRange, errm := strconv.Atoi(numbers[2])
			if errd != nil && errs != nil && errm != nil {
				check(errd)
			}
			converter := &mapConverter{
				index:       count,
				name:        name,
				destination: destination,
				source:      source,
				mesureRange: mesureRange,
			}
			fmt.Println(converter)
			if strings.Contains(name, "seed-to-soil map:") {
				cSoils = append(cSoils, *converter)
			}
			if strings.Contains(name, "soil-to-fertilizer map:") {
				cFertilizer = append(cFertilizer, *converter)
			}
			if strings.Contains(name, "fertilizer-to-water map:") {
				cWater = append(cWater, *converter)
			}
			if strings.Contains(name, "water-to-light map:") {
				cLight = append(cLight, *converter)
			}
			if strings.Contains(name, "light-to-temperature map:") {
				cTemperature = append(cTemperature, *converter)
			}
			if strings.Contains(name, "temperature-to-humidity map:") {
				cHumidity = append(cHumidity, *converter)
			}
			if strings.Contains(name, "humidity-to-location map:") {
				cLocation = append(cLocation, *converter)
			}
			count++

		}
	}
	return &cSoils, &cFertilizer, &cWater, &cLight, &cTemperature, &cHumidity, &cLocation
}

func getMinValue(seed int, converters []mapConverter) int {
	printInfoMethod("GetMinValue", map[string]string{
		"seed": strconv.Itoa(seed),
	})
	var val = seed
	for _, v := range converters {
		fmt.Println(v)
		sMax := v.source + v.mesureRange
		dMax := v.destination + v.mesureRange
		if v.source <= val && val < sMax {
			printInfoMethod("GetMinValue", map[string]string{
				"name":     v.name,
				"prev-val": strconv.Itoa(val),
			})
			val = (val - v.source) + v.destination
			printInfoMethod("GetMinValue", map[string]string{
				"index":       strconv.Itoa(v.index),
				"after-val":   strconv.Itoa(val),
				"source":      strconv.Itoa(v.source),
				"sMax":        strconv.Itoa(sMax),
				"destination": strconv.Itoa(v.destination),
				"dMax":        strconv.Itoa(dMax),
				"range":       strconv.Itoa(v.mesureRange),
			})
		}
	}
	return val
}

func main() {
	scanner, file := getFileScanner("input.txt")

	seeds := getSeeds(scanner)

	s,f,w,li,t,h,lo := getMapConverters(scanner)
	min := 0
	soil := soilConverter {
		converters: *s,
	}
	fetilizer := fertilizerConverter{
		converters: *f,
	}
	water := waterConverter{
		converters: *w,
	}
	light := lightConverter{
		converters: *li,
	}
	temperature := temperatureConverter{
		converters: *t,
	}
	humidity := humidityConverter{
		converters: *h,
	}
	location := locationConverter{
		converters: *lo,
	}
	// min = location.GetNewData(humidity.GetNewData(temperature.GetNewData(light.GetNewData(water.GetNewData(fetilizer.GetNewData(soil.GetNewData(seeds[1])))))))
	for _, v := range seeds {
		// temp := getMinValue(v, *converters)

		temp := location.GetNewData(humidity.GetNewData(temperature.GetNewData(light.GetNewData(water.GetNewData(fetilizer.GetNewData(soil.GetNewData(v)))))))

		println("____")

		printInfoMethod("min x seed", map[string]string{
			"seed": strconv.Itoa(v),
			"min":  strconv.Itoa(temp),
		})

		if min == 0 || temp < min {
			min = temp
		}

	}

	fmt.Println("***")
	fmt.Println(min)

	file.Close()
}
