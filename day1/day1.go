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
	for _, i := range values {
		for _, j := range values {
			elem, ok := values[2020-(i+j)]
			if ok {
				fmt.Println(elem * i * j)
				return
			}
		}
	}

}
