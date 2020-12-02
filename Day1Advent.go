package main

import (
	"fmt"
	"log"
	"sort"
	adventutilities "AdventOfCode/utils"
)



func printNumbers(nums []int) {

	for i, s := range nums {
		fmt.Println(i, s)
	}
}



func solvePuzzle1(numbers []int) (matchOne int, matchTwo int, product int) {

	if !(sort.IntsAreSorted(numbers)) {
		sort.Ints(numbers)
	}

	for i, s := range numbers {

		match2020 := 2020 - s

		log.Println("the integer at position ", i, "is ", s, " and the match we need is ", match2020)

		foundIndex := sort.SearchInts(numbers, match2020)

		if foundIndex > 0 && numbers[foundIndex] == match2020 {
			log.Println("We have found something! indexed at ", foundIndex)

			return s, match2020, s * match2020
		}

	}

	return 0, 1, 2
}

func solvePuzzle2(numbers []int) (matchOne int, matchTwo int, matchThree int, product int) {

	if !(sort.IntsAreSorted(numbers)) {
		sort.Ints(numbers)
	}

	for i, matchAttempt1 := range numbers {

		for j, matchAttempt2 := range numbers {

			log.Println("the integer at position ", i, "is ", matchAttempt1, "and j ", j, " is ", matchAttempt2)

			match2020 := 2020 - (matchAttempt1 + matchAttempt2)

			log.Println("We are looking for ", match2020)
			foundIndex := sort.SearchInts(numbers, match2020)

			if foundIndex > 0 && numbers[foundIndex] == match2020 {
				log.Println("We have found something! indexed at ", foundIndex)

				return matchAttempt1, matchAttempt2, match2020, matchAttempt1 * match2020 * matchAttempt2
			}
		}

	}

	return 0, 0, 1, 2
}

func main() {

	numbers, err := adventutilities.ReadNumbersFromFile("data/inputs_01_12.txt")
	adventutilities.Check(err)

	//printNumbers(numbers)

	//m1, m2, product := solvePuzzle1(numbers)

	//log.Println("Solved the puzzle! m1: ", m1, " m2: ", m2, "product: ", product)

	m1, m2, m3, product := solvePuzzle2(numbers)

	log.Println("Solved the puzzle! m1: ", m1, " m2: ", m2, " m3: ", m3, "product: ", product)

}
