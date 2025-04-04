package main
import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var twoD = [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("2d: ", twoD)
}

// compile: go build program_name.go
// execute after compiling: program_name.go
// compile+execute: go run program_name.go

