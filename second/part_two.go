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
	cmd = iota
	am
)

var (
	fileN string
)

type positionAim struct {
	horizontal, depth, aim int
}

func newPositionAim() *positionAim {
	return &positionAim{}
}

func (position *positionAim) update(command string, amount int) {
	if command == "forward" {
		position.horizontal += amount
		position.depth += position.aim * amount
	} else if command == "down" {
		position.aim += amount
	} else {
		position.aim -= amount
	}
}

func (position *positionAim) final() int {
	return position.horizontal * position.depth
}

func main() {

	test := flag.Bool("test", false, "Input file name")

	flag.Parse()

	if *test {
		fileN = "input_test.txt"
	} else {
		fileN = "input.txt"
	}

	file, err := os.Open(fileN)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	p := newPositionAim()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		depth, err := strconv.Atoi(line[am])

		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		p.update(line[cmd], depth)
	}

	log.Println(p.final())
}
