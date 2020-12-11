package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	c, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	partOne := fixedToboggan(c)
	partTwo := multiSlopeToboggan(c)

	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func countTrees(input []byte, xSlope int, ySlope int) int {
	xPos := 0
	yPos := ySlope
	trees := 0

	grid := bytes.Split(input, []byte("\n"))

	for i := yPos; i < len(grid); i += yPos {
		line := grid[i]
		r := strings.Split(string(line), "")

		xPos += xSlope
		if r[xPos%len(r)] == "#" {
			trees++
		}
	}

	return trees
}

func fixedToboggan(input []byte) int {
	return countTrees(input, 3, 1)
}

func multiSlopeToboggan(input []byte) int {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	trees := 1

	for _, slope := range slopes {
		xSlope := slope[0]
		ySlope := slope[1]

		trees *= countTrees(input, xSlope, ySlope)
	}

	return trees
}
