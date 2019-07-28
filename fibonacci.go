package main

import "fmt"
func myFib(x int) int {
	if x == 0 || x ==1 {
		return x
	}
	return myFib(x-1) + myFib(x-2)
}
func main()  {
	for i :=20; i < 30; i++{
		fmt.Println(i,"-->",myFib(i))
	}
}
