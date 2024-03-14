package main

import (
	"fmt"
)

func init() {
	fmt.Println("zaxo")
}
func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n /= 10
	}

	result := []int{}
	for i := range slc {
		result = append(result, slc[len(slc)-1-i])
	}

	return result
}
func countDigits(number int) int {
	var Cnt int
	Cnt = 0

	for number > 0 {
		number = number / 10
		Cnt = Cnt + 1
	}
	return Cnt
}

func main() {
	var flag bool = true

	for flag {

		length := 16
		fmt.Println("Enter the inputs")

		var cc int
		fmt.Scanln(&cc)
		if cc == -1 {
			flag = false
		}
		if countDigits(cc) != 16 {
			fmt.Println("Card number must be 16 digit")
			return
		}
		numbers := splitInt(cc)
		sum := 0
		for i := 0; i < length; i++ {

			switch i % 2 {
			case 0:
				if (numbers[i]*2)/10 >= 1 {
					twoDigitNumber := numbers[i] * 2
					digits := splitInt(twoDigitNumber)
					sum += digits[0] + digits[1]
				} else {
					sum += numbers[i] * 2
				}

			default:
				sum += numbers[i]
			}

		}
		if sum%10 == 0 {
			fmt.Println("CARD IS VALID")
		} else {
			fmt.Println("CARD IS INVALID")
		}

	}
}
