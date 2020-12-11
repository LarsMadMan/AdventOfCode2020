package main

import (
	adventutilities "AdventOfCode/utils"
	passport "AdventOfCode/utils/structs/Passport"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

func splitInputs(lines []string) (records []string) {

	record := ""

	for n, line := range lines {

		if line == "" {

			//new record
			strings.TrimSpace(record)
			records = append(records, record)
			record = ""
		} else {
			record += " "
			record += line
		}

		if n == len(lines)-1 {

			record += line
			strings.TrimSpace(record)
			records = append(records, record)
		}

	}
	return records

}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func isvalid(p passport.Passport) (isvalid bool, reason string) {

	err := validate.Struct(p)

	if err != nil {

		reason = err.Error()
		return false, reason
	}

	return true, reason
}

func solvePuzzle(fileName string) {
	lines, err := adventutilities.ReadStringsFromFile(fileName)
	adventutilities.Check(err)

	log.Println("len of lines", len(lines))

	records := splitInputs(lines)

	log.Println(len(records))

	numValidPassports := 0
	isValid := false
	reason := ""
	for _, record := range records {
		p := passport.MapStringToPassport(record)

		reason = ""
		isValid, reason = isvalid(p)

		log.Println("Passport with Birthyear", p.BirthYear, "is valid:", isValid, "reason:", reason)

		if isValid {

			numValidPassports++
		}

	}

	log.Println("Number of valid passports:", numValidPassports)
}

func main() {

	validate = validator.New()
	//need to work out how to write tests in Go, but leaving for another day
	test := true

	log.Println("test:", test)

	fileName := "data/inputs_04_12.txt"

	if test {
		fileName = "data/testinputs_04_12.txt"

		log.Println("running for the invalid file:")
		solvePuzzle("data/invalid_04_12.txt")

		log.Println("running for the valid file:")
		solvePuzzle("data/valid_04_12.txt")

	}

	solvePuzzle(fileName)

}
