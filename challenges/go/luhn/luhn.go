package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Printf("is card valid? %t\n", validateCard("4916252484357563"))
	fmt.Printf("is card valid? %t\n", validateCard("4716654980587001"))
	fmt.Printf("is card valid? %t\n", validateCard("4539735472444168"))
	fmt.Printf("is card valid? %t\n", validateCard("1234567891234567"))
}

func validateCard(pan string) bool {
	
	var multiDigits = doubleOddDigits(intToArray(pan))

	return sumArray(multiDigits, multiDigits[0]) % 10 ==0

}

func sumArray(digits []int64, checkDigit int64) int64 {

	var temp int64 

	for i := 1 ; len(digits) > i; i++ {

		temp = temp + digits[i]
		
	}

	return temp + checkDigit

}

func doubleOddDigits(digits []int64) []int64 {

	for i := 1 ; len(digits) > i; i++ {

		if i % 2 ==1 {

			var temp = digits[i]*2
			if temp>9 {
				digits[i] = (temp - 9)
			} else {
				digits[i] = temp
			}
			
		}
		
	}

	return digits

}

func intToArray(integer string) []int64 {

	var array = make([]int64,len([]rune(integer)))

	var temp,_ = strconv.ParseInt(integer, 10, 64)
    
	for i := 0 ; temp > 0; i++ {
		array[i] = temp % 10
		
		temp /= 10
	}
	return array
}
