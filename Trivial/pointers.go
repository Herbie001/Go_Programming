package main

import (
	"fmt"

)

func main() {

	i, j, k := 10, 15, 20
	p, q, s := &i, &j, &k
	fmt.Println(*p,*q,*s)

}