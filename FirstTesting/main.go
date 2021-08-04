package main

import "fmt"

func Factorial(input int)(result int){
	if(input == 0){
		return 1;
	}
	result = input*Factorial(input-1)
	return
}

func main() {
	fmt.Println(Factorial(4))
}
