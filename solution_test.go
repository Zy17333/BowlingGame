package BowlingGame

import (
	"fmt"
	"testing"
)

func TestScoreCalc(t *testing.T) {
	fmt.Println(ScoreCalc("X|X|X|X|X|X|X|X|X|X||XX"))
	fmt.Println(ScoreCalc("9-|9-|9-|9-|9-|9-|9-|9-|9-|9-||"))
	fmt.Println(ScoreCalc("5/|5/|5/|5/|5/|5/|5/|5/|5/|5/||5"))
	fmt.Println(ScoreCalc("X|7/|9-|X|-8|8/|-6|X|X|X||81"))
}