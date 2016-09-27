package main

import "fmt"

func main(){
	var control int = 0
	var num int
	fmt.Println("Enter number: ")
	fmt.Scan(&num)

	for control != 1{
		if num %2 == 0{
			num = num / 2
			//if the number is equal to 1 we end it as .it goes into an infinite loop
			if num == 1{
				control = 1
			}
			fmt.Println(num)
		} else if num %2 == 1{
			num = (num * 3) + 1
			fmt.Println(num)
		}
	}
}