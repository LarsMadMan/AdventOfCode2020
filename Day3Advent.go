package main

import (
	adventutilities "AdventOfCode/utils"
	"log"
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

func checkResult(testName string, actual int, expected int)(success bool){

	if(actual == expected){
		log.Println("Test:", testName, "successful, actual",actual,"== expected",expected)
		return true
	}else{
		log.Println("Test:", testName, "failed, expected", expected, "got",actual)
		return false
	}
}



func main() {

	//need to work out how to write tests in Go, but leaving for another day
	test := false

	log.Println("test:",test)

	fileName := "data/inputs_03_12.txt"

	if(test){
		fileName = "data/testinputs_03_12.txt"
	}

	lines, err := adventutilities.ReadStringsFromFile(fileName)
	adventutilities.Check(err)

	
	oneRightOneDownResult := solvePuzzleN(lines,1,1)
	log.Println("solvePuzzleN (right 1, down 1: numTrees:",oneRightOneDownResult) 
	adventutilities.CheckResult("oneRightOneDown", oneRightOneDownResult, 2)

	
	threeRightOneDownExpected := 200
	if(test){
		threeRightOneDownExpected = 7
	}
	
	threeRightOneDownResult := solvePuzzle1(lines)
	log.Println("solvePuzzle1 (threeRightOneDown, numTrees:",threeRightOneDownResult)
	adventutilities.CheckResult("threeRightOneDown", threeRightOneDownResult, threeRightOneDownExpected)


	fiveRightOneDownResult := solvePuzzleN(lines,5,1)
	log.Println("solvePuzzleN (right 5, down 1: numTrees:",fiveRightOneDownResult) 
	adventutilities.CheckResult("fiveRightOneDown", fiveRightOneDownResult, 3)
	
	
	sevenRightOneDownResult := solvePuzzleN(lines,7,1)
	log.Println("solvePuzzleN (right 7, down 1: numTrees:",sevenRightOneDownResult) 
	adventutilities.CheckResult("sevenRightOneDown", sevenRightOneDownResult, 4)

	oneRightTwoDownResult := solvePuzzleN(lines,1,2)
	log.Println("solvePuzzleN (right 1, down 2: numTrees:",oneRightTwoDownResult) 
	adventutilities.CheckResult("oneRightTwoDown", oneRightTwoDownResult, 2)

	product := oneRightOneDownResult * threeRightOneDownResult * fiveRightOneDownResult * sevenRightOneDownResult * oneRightTwoDownResult
	log.Println("Product:",product)
	adventutilities.CheckResult("Product", product, 336)
}
