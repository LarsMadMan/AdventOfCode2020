package main

import (
	"log"
	"strings"
	"strconv"
	adventutilities "AdventOfCode/utils"
)

func splitStringIntoJuicyBits(passwAndPolicy string)(minOcc int, maxOcc int, charToMatch string, password string){

	// trekk ut policy
	policy := strings.Split(passwAndPolicy," ")[0]
	minOcc,err := strconv.Atoi(strings.Split(policy, "-")[0])
	adventutilities.Check(err)

	maxOcc,err = strconv.Atoi(strings.Split(policy, "-")[1])
		adventutilities.Check(err)

	//find :
	i := strings.Index(passwAndPolicy, ":")
	charToMatch = passwAndPolicy[i-1:i]
	password = passwAndPolicy[i:]


	log.Println("line: ",passwAndPolicy, "passwd: ", password," policy:", policy," minOcc: ", minOcc, " maxOcc: ", maxOcc, "char to match is", charToMatch)

	return minOcc, maxOcc, charToMatch, password
}

func abidesByPolicy(password string, minOcc int, maxOcc int, charToMatch string)(ok bool){

	count := strings.Count(password, charToMatch)
		
	log.Println(password, count)

	return count >= minOcc && count <= maxOcc
}

func solvePuzzle1(passwords []string) (numValidPasswords int) {

	numValidPasswords = 0

	for _, line := range passwords{

		
		
		minOcc, maxOcc, charToMatch, password := splitStringIntoJuicyBits(line)
		
		

		if abidesByPolicy(password,minOcc,maxOcc,charToMatch){
			log.Println("valid password:", password)
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
	log.Println("numValid passwords:", numValid)
	
}
