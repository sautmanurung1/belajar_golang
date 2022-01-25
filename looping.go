package main

import "fmt"

func main(){
	var I,N int
	N = 45

	for I = 1; I<=5; I++{
		N = N + 5
	}
	fmt.Println(N)
}