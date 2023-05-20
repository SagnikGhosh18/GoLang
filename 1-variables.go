package main

import (
	"fmt"
	"strconv"
	//"strconv"
)

var i int = 56

var (
	actor string = "Emilia Clarke"
	role string = "Queen"
	rating int = 5
)

func main() {
	fmt.Println(i)
	var i int = 42					/* Shadowing - a concept where package level declaration takes precedence
													we cannot declare same variable twice within the same scope */
	fmt.Println(i) 
	var k = "Hello"
	j:= 54
	
	fmt.Printf("%v, %T\n",j , j)	//Prints the value and the type
	fmt.Printf("%v, %T\n",i , i)	//Prints the value and the type
	fmt.Printf("%v, %T\n",k , k)	//Prints the value and the type

	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a+b+c+d)

	fmt.Printf("%v\t%v\t%v\n",actor, role, rating)

	/**** TYPE CONVERSION ****/
	var m float32 = float32(i)
	fmt.Printf("%v, %T\n",i , i)	//Prints the value and the type
	fmt.Printf("%v, %T\n",m , m)	//Prints the value and the type

	var str string
	str = string(i)
	fmt.Printf("%v, %T\n",str , str)
	str = strconv.Itoa(i)
	fmt.Printf("%v, %T\n",str , str)
}
