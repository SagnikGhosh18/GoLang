package main

import(
	"fmt"
)

const (
	_ = iota								//this is a write only value, but used when we have no use for it
	a
	b
	c
)

func main(){
	const myConst int = 56					//typed
	const mySecondConst = 42				//untyped

	fmt.Printf("%v , %T\n",myConst,myConst)

	// enum consts

	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",b)
	fmt.Printf("%v\n",c)					//enums are scope restricted
}