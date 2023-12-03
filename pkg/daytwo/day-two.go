package daytwo

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
	"strconv"
	"strings"
)

type Game struct {
	id int
	red int
	blue int
	green int
}

type Color string

const (
	Red Color = "red"
	Blue = "blue"
	Green = "green"
)

func input() string {
	content, err := ioutil.ReadFile("./pkg/daytwo/day-two-input-test.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

func getColor(color Color, line string) int {
	// Set the limit depening on the color
	limit := 0
	switch color {
	case Red:
		limit = 12
	case Green:
		limit = 13
	case Blue:
		limit = 14
	}

	words := strings.FieldsFunc(line, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	var totalNumbers int

	for i := 0; i < len(words); i++ {
		if words[i] == string(color) && i > 0 {
			if num, err := strconv.Atoi(words[i-1]); err == nil {
				totalNumbers += num
			}
		}
	}

	if totalNumbers >= limit {
		return 0
	}

	return totalNumbers
}

func partOne(lines []string, games *[]Game) {
	for _, line := range lines {
		// No emptey string
		if line == "" {
			return
		}

		var game Game

		// Find the id of the game
		id, _ := strconv.Atoi(strings.Replace(strings.Split(line, ":")[0], "Game ", "", -1))
		game.id = id

		// Get each cube
		game.red = getColor(Red, line)
		game.green = getColor(Green, line)
		game.blue = getColor(Blue, line)

		if game.red != 0 && game.green != 0 && game.blue != 0 {
			*games = append(*games, game)
		}
	}
}

func SolveDayTwoPartOne() {
	lines := strings.Split(input(), "\n")

	var games []Game

	partOne(lines, &games)

	total := 0

	for _, game := range games {
		fmt.Println(game.id)
		total += game.id
	}

	fmt.Println("day two part one", total)
}
