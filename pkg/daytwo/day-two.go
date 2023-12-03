package daytwo

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Game struct {
	id int
}

func input() string {
	content, err := ioutil.ReadFile("./pkg/daytwo/day-two-input.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

func parseItems(sets []string) bool {
	limits := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	for _, set := range sets {
		singles := strings.Split(set, ",")

		for _, single := range singles {
			single = strings.TrimSpace(single)
			value := strings.Split(single, " ")

			qty, _ := strconv.Atoi(value[0])

			if qty > limits[value[1]] {
				return false
			}
		}
	}

	return true
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

		sets := strings.Split(strings.Split(line, ":")[1], ";")

		if parseItems(sets) {
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
		total += game.id
	}

	fmt.Println("day two part one", total)
}
