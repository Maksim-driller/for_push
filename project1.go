package main

import (
	"fmt"
	"strings"
)


func IntToLetters(num int) string {
	var newString string
	if num >= 1000 && num <= 3000 {
		num1 := num / 1000
		num2 := num % 1000
		if num1 != 4 && num1 != 9 {
			newString += strings.Repeat("G", num1)
			if num2 >= 500 {
				num2 -= 500
				newString += "F"
				num6 := num2 / 100
				newString += strings.Repeat("D", num6)
				num5 := num2 % 100
				num9 := num5 / 10
				if num9 != 4 && num9 != 9 {
					if num5 >= 50 {
						newString += "D"
						num10 := num5 - 50
						num11 := num10 / 10
						newString += strings.Repeat("C", num11)
						num12 = num10 % 10 
							num13 := num12 - 5
							newString += "B"
							newString += strings.Repeat("A", num13)
						}
					} else {
						num12 := num5
						switch num12 {
						case 1:
							newString += "A"
						case 2:
							newString += "AA"
						case 3:
							newString += "AAA"
						case 4:
							newString += "AAA"
						case 5:
							newString += "B"
						}
					}
				}
			}
		}
	}
	return newString
}

func main() {
	var a int
	fmt.Scanln(&a)
	lastString := IntToLetters(a)
	fmt.Println(lastString)
}

// далее не понял логику с формой (когда число равно 4 или 9)
