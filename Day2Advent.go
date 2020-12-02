package main

import (
	adventutilities "AdventOfCode/utils"
	"log"
	"strconv"
	"strings"
)

func splitStringIntoJuicyBits(passwAndPolicy string) (minOcc int, maxOcc int, charToMatch string, password string) {

	// trekk ut policy
	policy := strings.Split(passwAndPolicy, " ")[0]
	minOcc, err := strconv.Atoi(strings.Split(policy, "-")[0])
	adventutilities.Check(err)

	maxOcc, err = strconv.Atoi(strings.Split(policy, "-")[1])
	adventutilities.Check(err)

	//find :
	i := strings.Index(passwAndPolicy, ":")
	charToMatch = passwAndPolicy[i-1 : i]

	passwordIndex := strings.LastIndex(passwAndPolicy, " ")
	password = passwAndPolicy[passwordIndex+1:]
	//log.Println("extracted password", password, "from raw:", passwAndPolicy)
	//password = strings.TrimSpace(password)
	//log.Println("trimmed password", password, "NOSPACES")
	//log.Println("line: ",passwAndPolicy, "passwd: ", password," policy:", policy," minOcc: ", minOcc, " maxOcc: ", maxOcc, "char to match is", charToMatch)

	return minOcc, maxOcc, charToMatch, password
}

func abidesByPolicyPuzzle1(password string, minOcc int, maxOcc int, charToMatch string) (ok bool) {

	count := strings.Count(password, charToMatch)

	//log.Println(password, count)

	return count >= minOcc && count <= maxOcc
}

func abidesByPolicyPuzzle2(password string, pos1 int, pos2 int, charToMatch string) (ok bool) {

	//Toboggan Corporate Policies have no concept of "index zero"!
	adjustedPos1 := pos1 - 1
	adjustedPos2 := pos2 - 1

	trimmedPwd := strings.TrimSpace(password)

	runes := []rune(trimmedPwd)
	pos1Char := string(runes[adjustedPos1:pos1]) //we're taking one char, and orig pos is adjusted plus 1
	pos2Char := string(runes[adjustedPos2:pos2])

	if pos1Char == charToMatch && pos2Char != charToMatch {

		log.Println(trimmedPwd, "char at pos1(", pos1, ") is", pos1Char, "char at pos2 (", pos2, ") is", pos2Char, "char to match", charToMatch)

		return true
	}
	if pos1Char != charToMatch && pos2Char == charToMatch {

		log.Println(trimmedPwd, "char at pos1(", pos1, ") is", pos1Char, "char at pos2 (", pos2, ") is", pos2Char, "char to match", charToMatch)

		return true
	}

	return false
}

func solvePuzzle1(passwords []string) (numValidPasswords int) {

	numValidPasswords = 0

	for _, line := range passwords {

		minOcc, maxOcc, charToMatch, password := splitStringIntoJuicyBits(line)

		if abidesByPolicyPuzzle1(password, minOcc, maxOcc, charToMatch) {
			//log.Println("valid password:", password)
			numValidPasswords++
		}

	}

	return numValidPasswords
}

func solvePuzzle2(passwords []string) (numValidPasswords int) {

	numValidPasswords = 0

	for _, line := range passwords {

		pos1, pos2, charToMatch, password := splitStringIntoJuicyBits(line)

		if abidesByPolicyPuzzle2(password, pos1, pos2, charToMatch) {
			//log.Println("valid password:", password)
			numValidPasswords++
		}
	}

	return numValidPasswords

}

func main() {

	lines, err := adventutilities.ReadStringsFromFile("data/inputs_02_12.txt")
	adventutilities.Check(err)

	numValid := solvePuzzle1(lines)
	log.Println("got a file with ", len(lines), " passwords to check")
	log.Println("numValid passwords puzzle 1:", numValid)

	numValid = solvePuzzle2(lines)
	log.Println("numValid passwords puzzle 2:", numValid)

}
