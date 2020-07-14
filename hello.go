package main

import "fmt"

func add(x int, y int, z int) int {
	return x + y - z
}

func main(){
	fmt.Print("Hello, my Go world\n")
	fmt.Println(add(10, 10, 5))
}
