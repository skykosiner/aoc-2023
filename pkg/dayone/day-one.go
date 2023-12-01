package dayone

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func input() string {
	content, err := ioutil.ReadFile("./pkg/dayone/day-one-input.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

func charToInt(char string) int {
	output, err := strconv.Atoi(char)

	if err != nil {
		log.Fatal("Error converting rune to int", err)
	}

	return output
}

func partOne(lines []string, numbers *[]int) {
	for _, line := range lines {
		pointOne := 0
		pointTwo := 0

		// Make sure there is no blank lines
		if line == "" {
			return
		}

		var listOfNumbers []int

		for _, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				listOfNumbers = append(listOfNumbers, charToInt(string(char)))
			}
		}

		pointOne = listOfNumbers[0]
		pointTwo = listOfNumbers[len(listOfNumbers)-1:][0]

		numberCombined := fmt.Sprintf("%d%d", pointOne, pointTwo)
		*numbers = append(*numbers, charToInt(numberCombined))
	}
}

func SolveDayOne() {
	lines := strings.Split(input(), "\n")
	var numbers []int
	partOne(lines, &numbers)

	total := 0

	for _, number := range numbers {
		total += number
	}

	fmt.Println(total)
}
