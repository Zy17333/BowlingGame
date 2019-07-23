package BowlingGame

import (
	"fmt"
	"strconv"
	"strings"
)

const NumberOfPinsPerTurn = 10

type Hits []int

func ScoreCalc(input string) int {
	mainGameHits, totalHits, bonus := GetHitsInfo(input)
	score := 0
	for index, _ := range totalHits {
		for len(bonus) > 0 && index == bonus[0] {
			score += totalHits[bonus[0]]
			bonus = bonus[1:]
		}
		score += mainGameHits[index]
		fmt.Printf("Hit %v: Score = %v\n", index, score)
	}
	return score
}

func GetHitsInfo(input string) (Hits, Hits, []int) {
	mainGame, extraGame := SplitGame(input)
	mainGameHits, bonus := ParseMainGame(mainGame)
	totalHits := make(Hits, 0, 0)
	totalHits = append(totalHits, mainGameHits...)
	for _, c := range extraGame {
		switch c {
		case '-':
			totalHits = append(totalHits, 0)
		case '/':
			totalHits = append(totalHits, NumberOfPinsPerTurn-totalHits.Last())
		case 'X':
			totalHits = append(totalHits, NumberOfPinsPerTurn)
		default:
			num, _ := strconv.Atoi(string(c))
			totalHits = append(totalHits, num)
		}
		mainGameHits = append(mainGameHits, 0)
	}
	return mainGameHits, totalHits, bonus
}

func SplitGame(input string) (string, string) {
	games := strings.Split(input, "||")
	if len(games) <= 1 {
		return games[0], ""
	}
	return games[0], games[1]
}

func ParseMainGame(mainGame string) (Hits, []int) {
	mainGameHits := make(Hits, 0, 0)
	bonus := make([]int, 0, 0)
	for _, c := range mainGame {
		switch c {
		case '|':
			continue
		case '-':
			mainGameHits = append(mainGameHits, 0)
		case '/':
			mainGameHits = append(mainGameHits, NumberOfPinsPerTurn-mainGameHits.Last())
			bonus = append(bonus, mainGameHits.Len())
		case 'X':
			mainGameHits = append(mainGameHits, NumberOfPinsPerTurn)
			bonus = append(bonus, []int{mainGameHits.Len(), mainGameHits.Len() + 1}...)
		default:
			num, _ := strconv.Atoi(string(c))
			mainGameHits = append(mainGameHits, num)
		}
	}
	return mainGameHits, bonus
}

func (s Hits) Last() int {
	return s[len(s)-1]
}

func (s Hits) Len() int {
	return len(s)
}
