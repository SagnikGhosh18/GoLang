package main

import(
	"fmt"
)

func main(){
	myMap := make(map[string]int) 
	myMap = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["a"])

	myMap["d"] = 1
	fmt.Println(myMap)
	delete(myMap,"d")
	fmt.Println(myMap)

	_, ok := myMap["d"]
	fmt.Println(ok)

	fmt.Println(len(myMap))			// maps are passed by reference, changes are reflected

	
}