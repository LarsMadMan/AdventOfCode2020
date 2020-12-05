package main

import (
	adventutilities "AdventOfCode/utils"
	passport "AdventOfCode/utils/structs/Passport"
	"log"
	"strings"
)


func solvePuzzleN(lines []string, right int, down int)(numTrees int){

	lenY := len(lines)

	log.Println("numrows = ", lenY)

	s := lines[0]
	lenX := len(s)

	log.Println("numcols = ", lenX)

	numTrees = 0
	xPos := 0

	log.Println("right =", right)
	log.Println("down =", down)

	for n, line := range lines {

		//strictly not a good solution, but hey
		if(down == 2 && n%2!=0){
			continue
		}


		for i := 0; i < right; i++ {
				
			//do we need to reset?
			if(xPos == lenX ){

				xPos = 0
				log.Println("on line", n+1, " - resetting xPos")
			}

			runes := []rune(line)
			charToMatch := string(runes[xPos:xPos+1])

			if(charToMatch == "#" && i==0){
				numTrees++
				//log.Println(" new tree on line",n+1, " : num trees:",numTrees)
			}

			//log.Println("on line ", n+1,"xPos is", xPos+1, "got a",charToMatch)

			xPos++
		
		}
		
	}

	return numTrees
}

func solvePuzzle1(lines []string)(numTrees int){
	
	return solvePuzzleN(lines, 3,1)
	
}

func splitInputs(lines []string)(records[] string){

	record := ""

	for _, line := range lines {
		
		if line == "" {
			//new record
			log.Println("new record")
		}

	}
	return strings.Split(record, "\n")

}




func main() {

	//need to work out how to write tests in Go, but leaving for another day
	test := true

	log.Println("test:",test)

	fileName := "data/inputs_04_12.txt"

	if(test){
		fileName = "data/testinputs_04_12.txt"
	}

	lines, err := adventutilities.ReadStringsFromFile(fileName)
	adventutilities.Check(err)

	
	
	for _, line := range lines {
		p := passport.MapStringToPassport(line)

		p.Print()
	}

	
	
}
