package main

import "fmt"

func main() {
	e := []int{1, 2, 3} //slice
	e = append(e, 4)
	fmt.Println(e, len(e))

	f := make(map[string]int) // map
	f["one"] = 1
	f["two"] = 2
	fmt.Println(f, len(f), f["one"], f["three"])
}