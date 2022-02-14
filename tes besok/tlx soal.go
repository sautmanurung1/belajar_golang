package main

import "fmt"
func main(){
	var N int
	fmt.Scan(&N)
	if N>0 && N<=100 {
		fmt.Println("YES")
	}else {
		fmt.Println("Tidak")
	}
}