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

func inputTwo() string {
	// content, err := ioutil.ReadFile("./pkg/dayone/day-one-input-two-test.txt")
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
		log.Fatal("Error converting string to int", err)
	}

	return output
}

func isInt(s string) bool {
  _, err := strconv.ParseInt(s, 10, 64)
  return err == nil
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

func partTwo(lines []string) {
	answer := 0
	for _, line := range lines {
		var digits []int
		for i, c := range line {
			if c >= '0' && c <= '9' {
				digits = append(digits, charToInt(string(c)))
			}

			for d, val := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], val) {
					digits = append(digits, d+1)
				}
			}
		}

		if len(digits) != 0 {
			p2Int := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
			answer += charToInt(p2Int)
		}
	}

	fmt.Println("answer", answer)
}

func SolveDayOnePartOne() {
	lines := strings.Split(input(), "\n")
	var numbers []int
	partOne(lines, &numbers)

	total := 0

	for _, number := range numbers {
		total += number
	}

	fmt.Println(total)
}

func SolveDayOnePartTwo() {
	lines := strings.Split(inputTwo(), "\n")
	partTwo(lines)
}
