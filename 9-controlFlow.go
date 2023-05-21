package main

import(
	"fmt"
)

func main(){
	// defer
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	a:= 5
	defer fmt.Println(a)
	a=10

	//panic
	d,e := 1,0
	panic("Bad division")
	ans := d/e
	fmt.Println(ans)
}