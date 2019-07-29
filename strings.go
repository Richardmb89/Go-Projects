package main

import (
	"fmt"
	"strings"
)
func main()  {
	//Define a string
	var aString string = "Linux user and Developer"
	//Declare a constant string
	const anotherString = "Ricard Bee"
	for i := 0; i < len(aString);i++{
		fmt.Printf("%x ",aString[i])
	}
	fmt.Print("\n")
	//Using Binary
	for i := 0; i<len(aString);i++{
		fmt.Printf("%b ",aString[i])
	}
	fmt.Print("\n")
	//Using a Strings package
	fmt.Println(strings.Contains(aString,"Linux"))
	fmt.Println(strings.Contains(aString,"linux"))
	fmt.Println(strings.Count(anotherString,"e"))
	fmt.Println(strings.Count(anotherString,"r"))

	//Using a String Array
	var aStringArray = []string{"Linux","User","and Developer"}
	fmt.Println(strings.Join(aStringArray,"_*_*_*_"))

}