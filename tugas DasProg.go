package main

import "fmt"

func main(){
	var ganjil,genap,pilihan,batas int
	fmt.Println("Pilih")
	fmt.Println("1. Ganjil")
	fmt.Println("2. Genap")
	fmt.Printf("Masukkan Pilihan anda")
	fmt.Scanln(&pilihan)
	fmt.Printf("Masukkan batas anda : ")
	fmt.Scanln(&batas)
	if pilihan == 1{
		for i := batas; i >= 1; i--{
			if i%2 != 0{
				ganjil = i
				fmt.Println(ganjil)
			}
		}
	} else if pilihan == 2 {
		for i := batas; i >= 1; i--{
			if i%2 == 0{
				genap = i
				fmt.Println(genap)
			}
		}
	} else {
		fmt.Println("Pilihan anda salah")
	}
}