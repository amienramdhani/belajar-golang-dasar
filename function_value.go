package main

import "fmt"

func getGoodBye(name string) string {
	return "I Love You " + name
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("Alya"))
	fmt.Println(misal("Shafwa"))
}