package main

import (
	//"bufio"
	//"fmt"

	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	Day3Pt1()
}

func Day1() {
	adventInput := getInput("day1.txt")

	inputs := []int{}

	for i := 0; i < len(adventInput); i++ {
		value := adventInput[i]
		if value == "" {
			continue
		}
		inputs = append(inputs, ConvertToInt(value))
	}

	var result int

	for i := 1; i < len(inputs); i++ {
		if inputs[i] > inputs[i-1] {
			//fmt.Println("LOLW")
			result++
		}
	}

	fmt.Println(result)
}

func Day2Pt1() {
	adventInput := getInput("day2.txt")

	var resultx int
	var resulty int

	for i := 0; i < len(adventInput)-1; i++ {
		value := strings.Fields(adventInput[i])

		switch value[0] {
		case "forward":
			resultx = resultx + ConvertToInt(value[1])
		case "down":
			resulty = resulty + ConvertToInt(value[1])
		case "up":
			resulty = resulty - ConvertToInt((value[1]))
		}
	}

	fmt.Println(resultx * resulty)

}

func Day2Pt2() {
	adventInput := getInput("day2.txt")

	var resultx int
	var resulty int
	var aim int

	for i := 0; i < len(adventInput)-1; i++ {
		value := strings.Fields(adventInput[i])

		switch value[0] {
		case "forward":
			resultx = resultx + ConvertToInt(value[1])
			resulty = resulty + (ConvertToInt(value[1]) * aim)
		case "down":
			aim = aim + ConvertToInt(value[1])
		case "up":
			aim = aim - ConvertToInt(value[1])
		}
	}

	fmt.Println(resultx * resulty)

}

func Day3Pt1() {
	adventInput := getInput("day3.txt")

	s := make([]int, 12)

	//counts the amount of 1's in the array, at each place in string
	for i := 0; i < len(adventInput)-1; i++ {
		value := strings.Split(adventInput[i], "")
		fmt.Println(value)
		for j := 0; j < len(value); j++ {
			fmt.Println(value[j])
			s[j] = s[j] + ConvertToInt(value[j])
		}
	}

	fmt.Println(s)

	//Converts the array into something where there are ones and zeros whether or not it is the majority.
	for i := 0; i < len(s); i++ {
		if s[i] > len(adventInput)/2 {
			s[i] = 1
		} else {
			s[i] = 0
		}
	}

	gamma, err := strconv.ParseInt(arrayToString(s, ""), 2, 64)

	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			s[i] = 1
		} else {
			s[i] = 0
		}
	}

	epsilon, err := strconv.ParseInt(arrayToString(s, ""), 2, 64)

	fmt.Printf("%v %v %v", gamma, epsilon, gamma*epsilon)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func ConvertToInt(lol string) int {
	intValue, err := strconv.Atoi(lol)
	if err != nil {
		panic("Could not convert string to integer")
	}
	return intValue
}

func getInput(file string) []string {
	f, err := os.ReadFile(file)

	adventInput := strings.Split(string(f), "\n")

	if err != nil {
		log.Fatal(err)
	}

	return adventInput
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
