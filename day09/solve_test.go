package day09_test

import (
	"spissable/advent-of-go-template/day09"
	"spissable/advent-of-go-template/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	result := day09.SolvePuzzle1(input)
	t.Log(result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	result := day09.SolvePuzzle2(input)
	t.Log(result)
}
