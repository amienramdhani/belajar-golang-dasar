package main

import "fmt"

func logging(){
    fmt.Println("Selesai Melakukan Logging")
}

func afterLogging(){
    defer logging()
    fmt.Println("Fungsi ini akan dijalankan sebelum defer")
}

func main() {
  afterLogging()
}