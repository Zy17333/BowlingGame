package BowlingGame

import (
	"fmt"
	"testing"
)

var testCases = []string{
	"X|X|X|X|X|X|X|X|X|X||XX",
	"9-|9-|9-|9-|9-|9-|9-|9-|9-|9-||",
	"5/|5/|5/|5/|5/|5/|5/|5/|5/|5/||5",
	"X|7/|9-|X|-8|8/|-6|X|X|X||81",
	"X|71|12|23|34|45|5/|--|X|9/||X",
}

func TestScoreCalc(t *testing.T) {
	for index, testCase := range testCases {
		fmt.Printf("Case %v:\n", index+1)
		score := ScoreCalc(testCase)
		fmt.Println("Total Score = ", score)
		fmt.Println("============================")
	}
}
