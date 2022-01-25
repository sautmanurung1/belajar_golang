package main

import "fmt"
func main(){
	var nilai int
	println("Masukkan nilai : ")
	fmt.Scanln(&nilai)
	if nilai>85 && nilai<=100{
		fmt.Println("Nilai anda A")
	} else if nilai>75 && nilai<=84{
		fmt.Println("Nilai anda B")
	} else if nilai>60 && nilai <=74{
		fmt.Println("Nilai anda C")
	} else if nilai>40 && nilai <=59{
		fmt.Println("Nilai anda D")
	} else{
		fmt.Println("Nilai anda E")
	}
}