package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	command = iota
	amount
)

var (
	fileName string
)

type position struct {
	horizontal, depth int
}

func newPosition() *position {
	return &position{}
}

func (position *position) update(command string, amount int) {
	if command == "forward" {
		position.horizontal += amount
	} else if command == "down" {
		position.depth += amount
	} else {
		position.depth -= amount
	}
}

func (position *position) final() int {
	return position.horizontal * position.depth
}

func main() {

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

	p := newPosition()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		depth, err := strconv.Atoi(line[amount])

		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		p.update(line[command], depth)
	}

	log.Println(p.final())
}