package main

import "fmt"

func main() {
	var dia int
	var mes int

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia < 20 {
			fmt.Print("Capricornio")
		} else if dia >= 20 && dia < 32 {
			fmt.Print("Acuario")
		} else {
			fmt.Print("Error")
		}
	case 2:
		if dia < 19 {
			fmt.Print("Acuario")
		} else if dia >= 19 && dia < 29 {
			fmt.Print("Pisis")
		} else {
			fmt.Print("Error")
		}
	case 3:
		if dia < 21 {
			fmt.Print("Pisis")
		} else if dia >= 21 && dia < 32 {
			fmt.Print("Aries")
		} else {
			fmt.Print("Error")
		}
	case 4:
		if dia < 20 {
			fmt.Print("Aries")
		} else if dia >= 20 && dia < 31 {
			fmt.Print("Tauro")
		} else {
			fmt.Print("Error")
		}
	case 5:
		if dia < 21 {
			fmt.Print("Tauro")
		} else if dia >= 21 && dia < 32 {
			fmt.Print("Géminis")
		} else {
			fmt.Print("Error")
		}
	case 6:
		if dia < 21 {
			fmt.Print("Géminis")
		} else if dia >= 21 && dia < 31 {
			fmt.Print("Cáncer")
		} else {
			fmt.Print("Error")
		}
	case 7:
		if dia < 23 {
			fmt.Print("Cáncer")
		} else if dia >= 23 && dia < 32 {
			fmt.Print("Leo")
		} else {
			fmt.Print("Error")
		}
	case 8:
		if dia < 23 {
			fmt.Print("Leo")
		} else if dia >= 23 && dia < 32 {
			fmt.Print("Virgo")
		} else {
			fmt.Print("Error")
		}
	case 9:
		if dia < 23 {
			fmt.Print("Virgo")
		} else if dia >= 23 && dia < 31 {
			fmt.Print("Libra")
		} else {
			fmt.Print("Error")
		}
	case 10:
		if dia < 23 {
			fmt.Print("Libra")
		} else if dia >= 23 && dia < 32 {
			fmt.Print("Escorpio")
		} else {
			fmt.Print("Error")
		}
	case 11:
		if dia < 22 {
			fmt.Print("Escorpio")
		} else if dia >= 22 && dia < 31 {
			fmt.Print("Sagitario")
		} else {
			fmt.Print("Error")
		}
	case 12:
		if dia < 22 {
			fmt.Print("Sagitario")
		} else if dia >= 22 && dia < 32 {
			fmt.Print("Capricornio")
		} else {
			fmt.Print("Error")
		}
	default:
		fmt.Print("Error")
	}
}
