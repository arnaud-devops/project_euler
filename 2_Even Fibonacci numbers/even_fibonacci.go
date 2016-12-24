package main

import "fmt"

func main() {
	i := 1
	tmp := 0
	ret := 0
	for j := 1 ; j <= 4000000 ; { 
		tmp = j
		j += i
		i = tmp
		if j % 2 == 0 {
			ret += j
		}
	}
	fmt.Println(ret);
}
