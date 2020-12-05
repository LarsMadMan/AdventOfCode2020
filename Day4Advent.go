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


	for n, line := range lines {
		
		if line == ""  {
			
			//new record
			strings.TrimSpace(record)
			records = append(records, record)
			record = ""
		}else{
			record += " " 
			record += line
		}

		if(n == len(lines)-1){
			
			record += line
			strings.TrimSpace(record)
			records = append(records, record)
		}
		

	}
	return records

}

func isvalid(p passport.Passport)(isvalid bool, reason string){
	

	if p.BirthYear == 0{
		reason = "p.BirthYear is missing"
		return false, reason;
	}

	if p.IssueYear == 0{
		reason = "p.IssueYear is missing"
		return false, reason;
	}

	if p.ExpirationYear == 0{
		reason = "p.ExpirationYear is missing"
		return false, reason;
	}

	if p.PassportId == ""{
		reason = "p.PassportId is missing"
		return false, reason;
	}

	if p.Height == ""{
		reason = "p.Height is missing"
		return false, reason;
	}

	if p.HairColor == ""{
		reason = "p.HairColor is missing"
		return false, reason;
	}
	if p.EyeColor == ""{
		reason = "p.EyeColor is missing"
		return false, reason;
	}
	

	return true, reason
}



func main() {

	//need to work out how to write tests in Go, but leaving for another day
	test := false

	log.Println("test:",test)

	fileName := "data/inputs_04_12.txt"

	if(test){
		fileName = "data/testinputs_04_12.txt"
	}

	lines, err := adventutilities.ReadStringsFromFile(fileName)
	adventutilities.Check(err)

	log.Println("len of lines",len(lines))

	records := splitInputs(lines)
	
	log.Println(len(records))
	
	numValidPassports := 0
	isValid := false
	reason := ""
	for _, record := range records {
		p := passport.MapStringToPassport(record)

		reason = ""
		isValid, reason = isvalid(p)

		log.Println("Passport with Birthyear",p.BirthYear, "is valid:",isValid,"reason:",reason)

		if isValid{
			
			numValidPassports++
		}
		

		
	}

	log.Println("Number of valid passports:",numValidPassports)
	
}
