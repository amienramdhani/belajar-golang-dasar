package main
import "fmt"

type Orang struct {
    Nama, Alamat string
    Umur int
}

func main() {
    var alya Orang
    alya.Nama = "Alya Shafwa Wafiya"
    alya.Alamat = "Bekasi"
    alya.Umur = 23
    fmt.Println(alya)
    
    amien := Orang{
        Nama : "Muhammad Amien Ramdhani",
        Alamat : "Cirebon",
        Umur : 22,
    }
    fmt.Println(amien)
    
    ramdhani := Orang{ "Ramdhani" , "Cirebon" , 22}
    fmt.Println(ramdhani)
}