package passport

import (  
    "log"
)

type Passport struct {  
    BirthYear   int
    IssueYear    int
    ExpirationYear int
    HairColor string
    EyeColor string
    Height string
    PassportId int
    CountryId int
}

func (p Passport) Print() {  
   
    log.Println("BirthYear:",p.BirthYear)
    log.Println("IssueYear:",p.IssueYear)
    log.Println("ExpirationYear:",p.ExpirationYear)
    log.Println("HairColor:",p.HairColor)
    log.Println("EyeColor:",p.EyeColor)
    log.Println("Height:",p.Height)
    log.Println("PassportId:",p.PassportId)
    log.Println("CountryId:",p.CountryId)
}

func MapStringToPassport(rawData string)(p Passport){
    p = Passport{}


    return p
}