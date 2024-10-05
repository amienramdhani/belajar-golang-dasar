package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Alya")
	fmt.Println((result))

	fmt.Println(getHello("Amien"))
	fmt.Println(getHello("Shafwa"))

}