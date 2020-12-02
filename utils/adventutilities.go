package adventutilities

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadStringsFromFile(fileName string) (lines []string, err error) {
	data, err := ioutil.ReadFile(fileName)
	Check(err)

	lines = strings.Split(string(data), "\n")

	return lines, nil
}

func ReadNumbersFromFile(fileName string) (nums []int, err error) {

	data, err := ioutil.ReadFile(fileName)
	Check(err)

	lines := strings.Split(string(data), "\n")

	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	var val int = 0

	for _, s := range lines {

		val, err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		nums = append(nums, val)

	}

	return nums, nil
}