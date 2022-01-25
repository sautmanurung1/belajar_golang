package main

import "fmt"

func main(){
	var numChild int
	fmt.Scanln(&numChild)
	switch numChild {
		case 1:
			fmt.Println("You have 1 child")
			break
		case 2:
			fmt.Println("You have 2 children")
			break
		case 3:
			fmt.Println("You have 3 children")
			break
		case 4:
			fmt.Println("You have 4 children")
			break
		case 5:
			fmt.Println("You have 5 children")
			break
		case 6:
			fmt.Println("You have 6 children")
			break
		case 7:
			fmt.Println("You have 7 children")
			break
		case 10:
			fmt.Println("You have 10 children")
			break
		default:
			fmt.Println("You have 15 children")
			break
	}
}