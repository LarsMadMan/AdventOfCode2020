package passport

import (
	"log"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      int    `validate:"required,gte=1920,lte=2002`
	IssueYear      int    `validate:"required,gte=2010,lte=2020`
	ExpirationYear int    `validate:"required,gte=2020,lte=2030`
	HairColor      string `validate:"hexcolor|rgb|rgba"`
	EyeColor       string `validate:"required"`
	Height         string `validate:"required"`
	PassportId     string `validate:"required"`
	CountryId      int
}

func (p Passport) Print() {

	log.Println("BirthYear:", p.BirthYear)
	log.Println("IssueYear:", p.IssueYear)
	log.Println("ExpirationYear:", p.ExpirationYear)
	log.Println("HairColor:", p.HairColor)
	log.Println("EyeColor:", p.EyeColor)
	log.Println("Height:", p.Height)
	log.Println("PassportId:", p.PassportId)
	log.Println("CountryId:", p.CountryId)
}

func MapStringToPassport(rawData string) (p Passport) {

	//log.Println("==========MapStringToPassport=======")

	p = Passport{}

	splitIntoKeyValues := strings.Split(rawData, " ")

	for _, keyValue := range splitIntoKeyValues {

		keyValueSplitted := strings.Split(keyValue, ":")

		switch {
		case keyValueSplitted[0] == "hcl":
			p.HairColor = keyValueSplitted[1]
		case keyValueSplitted[0] == "byr":
			val, _ := strconv.Atoi(keyValueSplitted[1])
			p.BirthYear = val
		case keyValueSplitted[0] == "iyr":
			val, _ := strconv.Atoi(keyValueSplitted[1])
			p.IssueYear = val
		case keyValueSplitted[0] == "eyr":
			val, _ := strconv.Atoi(keyValueSplitted[1])
			p.ExpirationYear = val
		case keyValueSplitted[0] == "pid":
			p.PassportId = keyValueSplitted[1]
		case keyValueSplitted[0] == "cid":
			val, _ := strconv.Atoi(keyValueSplitted[1])
			p.CountryId = val
		case keyValueSplitted[0] == "ecl":
			p.EyeColor = keyValueSplitted[1]
		case keyValueSplitted[0] == "hgt":
			p.Height = keyValueSplitted[1]
		}
	}

	return p
}
