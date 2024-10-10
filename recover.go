package main

import "fmt"

func logging(){
    fmt.Println("Selesai Melakukan Logging")
    message := recover()
    fmt.Println("Terjadi Panic", message)
}

func afterLogging(error bool){
    defer logging()
    if error{
        panic("Terjadi Error pada fungsi ini")
    }
}

func main() {
  afterLogging(true)
}