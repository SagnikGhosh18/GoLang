package main

import (
	"fmt"
)

func main(){
	if true{
		fmt.Println("Hello world")
	}

	myMap := make(map[string]int) 
	myMap = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}


	// initalizer syntax of the if statement
	if pop, ok := myMap["a"]; ok {
		fmt.Println(pop)
	}

	a:= 5
	b:= 6
	if a==b{
		fmt.Println("True")
	}
	
	if a==5 || a>6{
		fmt.Println(true)
	} else if a==b {
		fmt.Println(false)
	} else {
		fmt.Println("meh")
	}

	switch 4 {
		case 1: fmt.Println("One")
		case 2,3,4: fmt.Println("Two")
		default: fmt.Println("Neither one nor two")
	}
	switch l:=1+1; l {
		case 1: fmt.Println("One")
		case 2,3,4: fmt.Println("Two")
		default: fmt.Println("Neither one nor two")

	}

	i:=98
	switch {
	case i<=100: fmt.Println("Less than 100")
				  fallthrough	
	case i<=20: fmt.Println("Less than 20")
	default: fmt.Println("meh")
	}

	var j interface{} = 1

	switch j.(type){
	case int: fmt.Println("j is an int")
			  break
			  fmt.Println("This won't execute")
	case float64: fmt.Println("j is a float64")
	case string: fmt.Println("j is a string")
	default: fmt.Println("none")
	}

}

