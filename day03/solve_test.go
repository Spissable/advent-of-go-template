package day03_test

import (
	"spissable/advent-of-go-template/day03"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	result := day03.SolvePuzzle1(input)
	t.Log(result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	result := day03.SolvePuzzle2(input)
	t.Log(result)
}
