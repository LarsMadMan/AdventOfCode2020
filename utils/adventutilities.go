package adventutilities

import (
	"io/ioutil"
	"strconv"
	"strings"
	"log"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckResult(testName string, actual int, expected int)(success bool){

	if(actual == expected){
		log.Println("Test:", testName, "successful, actual",actual,"== expected",expected)
		return true
	}else{
		log.Println("Test:", testName, "failed, expected", expected, "got",actual)
		return false
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