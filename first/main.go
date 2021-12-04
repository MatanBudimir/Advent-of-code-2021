package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		fileName    string
		previous    int
		largerCount int  = 0
		first       bool = true
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

		if first {
			first = false
			previous = num
			continue
		}

		if previous < num {
			largerCount++
		}

		previous = num
	}

	log.Println(largerCount)
}
