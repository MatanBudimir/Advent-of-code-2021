package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

	oxygen, err := numToInt(filter(b, true))

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	coTwoScrubber, err := numToInt(filter(b, false))

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	log.Println(oxygen*coTwoScrubber)
}

func filter(numbers []byte, max bool) string {
	nums := strings.Split(string(numbers), "\n")

	for i := 0; i < len(nums[0])-1; i++ {
		var zero, one int

		if len(nums) == 1 {
			break
		}

		for _, num := range nums {
			if num[i] == '0' {
				zero++
			} else {
				one ++
			}
		}

		if one > zero || one == zero {
			if max {
				nums = getNumbers(nums, '1', i)
			} else {
				nums = getNumbers(nums, '0', i)
			}
		} else {
			if max {
				nums = getNumbers(nums, '0', i)
			} else {
				nums = getNumbers(nums, '1', i)
			}
		}
	}

	return nums[0]
}

func getNumbers(numbers []string, val byte, index int) []string {
	var res []string
	for _, num := range numbers {
		if num[index] == val {
			res = append(res, num)
		}
	}

	return res
}

func numToInt(num string) (int64, error) {
	number := strings.TrimSuffix(num, "\r")
	return strconv.ParseInt(number, 2, 64)
}