package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	validPartOne := 0
	validPartTwo := 0

	for scanner.Scan() {
		val := scanner.Text()

		values := strings.Split(val, " ")
		minMax := strings.Split(values[0], "-")

		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		letter := strings.TrimSuffix(values[1], ":")
		password := values[2]

		letterCount := make(map[string]int, 0)

		for i := 0; i < len(password); i++ {
			letterCount[string(password[i])]++
		}

		count, ok := letterCount[letter]
		if ok && count >= min && count <= max {
			validPartOne++
		}

		// part 2
		if (string(password[min-1]) == letter) != (string(password[max-1]) == letter) {
			validPartTwo++
		}

	}

	fmt.Printf("part one: %d\npart two: %d\n", validPartOne, validPartTwo)

}
