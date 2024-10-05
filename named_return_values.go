package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Alya"
	middleName = "Shafwa"
	lastName = "Wafiya"

	return firstName,middleName,lastName
}

func main() {
	a,b,c := getCompleteName()	

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}