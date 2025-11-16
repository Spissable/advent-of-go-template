package day12_test

import (
	"spissable/advent-of-go-template/day12"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	result := day12.SolvePuzzle1(input)
	t.Log(result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	result := day12.SolvePuzzle2(input)
	t.Log(result)
}
