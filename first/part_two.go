package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type previousResult struct {
	lastIndex, num int
}

func main() {
	var (
		fileName string
		previous previousResult
		largerCount int = 0
		first bool = true
		nums []int
	)

	test := flag.Bool("test", false, "Input file name")

	flag.Parse()

	if *test {
		fileName = "input_test.txt"
	} else {
		fileName = "input.txt"
	}

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		num, err := strconv.Atoi(line)

		if err != nil {
			log.Fatalf("Failed to parse %s.", line)
			return
		}

		nums = append(nums, num)
	}

	for {
		if first {
			first = false
			previous = previousResult{3, sum(nums[0:3])}
			continue
		}

		if previous.lastIndex > len(nums) {
			break
		}

		s := sum(nums[previous.lastIndex-3:previous.lastIndex])

		if s > previous.num {
			largerCount++
		}

		previous = previousResult{previous.lastIndex+1, s}
	}

	log.Println(largerCount)
}

func sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}