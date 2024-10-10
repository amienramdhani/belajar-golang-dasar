package main

import "fmt"

func logging(){
    fmt.Println("Selesai Melakukan Logging")
}

func afterLogging(error bool){
    defer logging()
    if error{
        panic("Terjadi Error")
        
    }
}

func main() {
  afterLogging(true)
}