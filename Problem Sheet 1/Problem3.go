//Write a program that accepts a users inputted
//string and print out the reverse
package main
import (
	"fmt"
	"bufio"
	"os"
)
func Reverse(s string) string {
	// Here we use runes in the reverse function to reverse a string.
	//each character is taken as an individual element so we can loop through them
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	word1 := text
	reversed1 := Reverse(word1)
	fmt.Println(reversed1)
}