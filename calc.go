package BowlingGame

import "fmt"

type Turn struct {
	KnockNumbers Hits
}

const (
	NumberOfTurnsPerGame = 10
	NumberOfBallsPerTurn = 2
	NumberOfPinsPerTurn  = 10
)

var BonusBalls = []int{2, 1}

// input is the numbers of pins knocked down by each ball
func Solution(input Hits) int {
	turns, ok := FormatTurns(input)
	if !ok {
		return -1
	}
	return Calc(turns, input)
}

func Calc(turns []Turn, input Hits) int {
	score := 0
	currentBall := 0
	for index, turn := range turns {
		score += turn.Sum()
		if turn.Sum() == NumberOfPinsPerTurn && index < NumberOfTurnsPerGame-1 {
			for ball := currentBall + 1; ball <= currentBall+BonusBalls[turn.Len()-1]; ball++ {
				score += input[ball]
			}
		}
	}
	return score
}

// judge whether a input is valid
func FormatTurns(input Hits) ([]Turn, bool) {
	turns := make([]Turn, NumberOfTurnsPerGame)
	currentTurn := 0
	isLastTurn := false
	for _, pins := range input {
		if currentTurn == NumberOfTurnsPerGame-1 {
			isLastTurn = true
		}
		if turns[currentTurn].IsFinished() && !isLastTurn {
			currentTurn++
		}
		if pins+turns[currentTurn].Sum() > NumberOfPinsPerTurn && !isLastTurn {
			fmt.Println("One Turn knocked too much pins")
			return nil, false
		}
		turns[currentTurn].KnockNumbers = append(turns[currentTurn].KnockNumbers, pins)
	}
	if !isLastTurn || !LastTurnOK(turns[NumberOfTurnsPerGame-1]) {
		return nil, false
	}
	return turns, true
}

func LastTurnOK(t Turn) bool {
	// still some pins left
	if t.Sum() <= NumberOfPinsPerTurn {
		fmt.Println("pins left in last turn")
		return t.Len() == NumberOfBallsPerTurn
	}
	sum := 0
	for index, pins := range t.KnockNumbers {
		sum += pins
		if sum == NumberOfPinsPerTurn && index+1 <= NumberOfBallsPerTurn {
			return index+1+BonusBalls[index] == t.Len()
		}
	}
	fmt.Println("pins number incorrect in last turn")
	return false
}

func (t *Turn) Sum() int {
	sum := 0
	for _, points := range t.KnockNumbers {
		sum += points
	}
	return sum
}

func (t *Turn) Len() int {
	return len(t.KnockNumbers)
}

func (t *Turn) IsFinished() bool {
	return t.Len() == NumberOfBallsPerTurn || t.Sum() == NumberOfPinsPerTurn
}
