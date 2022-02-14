package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var p,q,m int

	for i := 0; i < n; i++{
		fmt.Scan(&p,&q,&m)
		for m > 0{
			if p > q{
				p = p / 2
				m --
			}else {
				q = q/2
				m --
			}
	}
	if p > q {
            fmt.Println(p," ",q)
        } else {
            fmt.Println(q," ",p)
		}
    }
}