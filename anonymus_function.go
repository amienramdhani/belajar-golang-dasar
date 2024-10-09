package main

import "fmt"

type Blacklist func(name string)bool

func isRegister(name string,  blacklist Blacklist){
    if blacklist(name){
        fmt.Println("You are blocked", name)
    }else{
        fmt.Println("Hello", name)
    }
}

func main() {
  blacklist := func(name string)bool{
      return name == "anjing"
  }
  
  isRegister("Alya", blacklist)
  isRegister("anjing", func(name string)bool{
      return name == "anjing"
  })
}