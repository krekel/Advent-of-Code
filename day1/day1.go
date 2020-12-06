package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	values := make(map[int]int)
	vals := make([]int, 0, 30)

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

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		values[val] = val
		vals = append(vals, val)
	}

	// part 1
	for _, v := range values {
		elem, ok := values[2020-v]
		if ok {
			fmt.Println(elem * v)
			break
		}
	}

	// part 2
	for j := 0; j < len(vals); j++ {
		for k := j + 1; k < len(vals); k++ {
			elem, ok := values[2020-(vals[k]+vals[j])]
			if ok {
				fmt.Println(elem * vals[j] * vals[k])
				return
			}
		}
	}

}
