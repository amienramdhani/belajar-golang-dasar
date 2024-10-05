package main

import "fmt"

func getFullName() (string,string) {	
	return "Alya","Shafwa"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName,lastName)

}