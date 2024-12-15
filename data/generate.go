package data

import (
	"fmt"
	"golang.org/x/exp/rand"
)

func Generate(dataType string) any {
	switch dataType{
	case TYPE_NAME:
		return generateName()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	case TYPE_DATE:
		return generateDate()
	default:
		return ""
	}
}

func generateDate() string {
	year := 1900 + rand.Intn(124)
	month := 1 + rand.Intn(12)
	var day int
	if month == 2 {
		if year%4 == 0 {
			day = 1 + rand.Intn(29)
		} else {
			day = 1 + rand.Intn(28)
		}	
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		day = 1 + rand.Intn(30)
	} else {
		day = 1 + rand.Intn(31)
	}
	return fmt.Sprintf("%02d-%02d-%d", year, month, day)
}

func generatePhone() string {
	num := "08"
	for i := 0; i < 10; i++ {
		num += fmt.Sprint(rand.Intn(10))
	}
	return num
}

func generateAddress() string {
	randomIndex := rand.Intn(len(address)+ 1)
	return address[randomIndex]
}

func generateName() string {
	randomIndex := rand.Intn(len(name)+1)
	return name[randomIndex]
}