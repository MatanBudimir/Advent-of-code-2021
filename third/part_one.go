package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type commonValues []string

func main() {
	var (
		fileName string
	)
	test := flag.Bool("test", false, "Input file name")

	flag.Parse()

	if *test {
		fileName = "input_test.txt"
	} else {
		fileName = "input.txt"
	}

	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	mostCommon := getCommonValues(b, true)
	gamma, err := number(mostCommon)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	leastCommon := getCommonValues(b, false)
	epsilon, err := number(leastCommon)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	log.Println(gamma * epsilon)
}

func getCommonValues(data []byte, mostCommon bool) commonValues {
	vals := strings.Split(string(data), "\n")

	var common commonValues

	for i := 0; i < len(vals[0])-1; i++ {
		var zero, one int

		for _, val := range vals {
			if val[i] == '0' {
				zero++
			} else {
				one++
			}
		}

		if mostCommon {
			if zero > one {
				common = append(common, "0")
			} else {
				common = append(common, "1")
			}
		} else {
			if zero > one {
				common = append(common, "1")
			} else {
				common = append(common, "0")
			}
		}
	}

	return common
}

func number(common commonValues) (int64, error) {
	binary := strings.Join(common, "")

	num, err := strconv.ParseInt(binary, 2, 64)

	if err != nil {
		return 0, err
	}

	return num, nil
}
