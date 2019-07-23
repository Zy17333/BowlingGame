package BowlingGame

import (
	"fmt"
	"testing"
)

var testCases map[int]map[int]map[int][][]int

func testCaseInit() {
	testCases = make(map[int]map[int]map[int][][]int)
	testCases[10] = make(map[int]map[int][][]int)
	testCases[10][10] = make(map[int][][]int)
	testCases[10][10][2] = [][]int{
		[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 0, 7, 2, 1, 1, 10, 0, 10, 0, 10},
	}

}

func TestFormatTurns(t *testing.T) {
	testCaseInit()
	for index, testCase := range testCases[NumberOfTurnsPerGame][NumberOfPinsPerTurn][NumberOfBallsPerTurn] {
		_, ok := FormatTurns(testCase)
		if ok {
			fmt.Printf("Case %v: %v\n", index+1, Solution(testCase))
		} else {
			fmt.Printf("Case %v: Input Invalid\n", index+1)
		}
	}
}
