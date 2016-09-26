//this is incorrect
package main

import "fmt"

const prime = 10001

//To check is a number is prime or not
func checkIfPrimeNum(num int) bool{
	if num % 2 == 0{
		return false
	}

	for count := 3; count * count <= num; count +=2{
		if num % count == 0{
			return false;
		}
	}
	return true
}

func main(){
	primeCount, number := 0, 0

	for{
		number++
		if checkIfPrimeNum(number){
			primeCount++
		}
		if primeCount == prime{
			break
		}
	}
	fmt.Print("The 10001th prime number is ", number)
}