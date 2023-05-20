package main

import(
	"fmt"
)

func main(){
	// BOOLEAN
	fmt.Println("BOOLEAN TYPE\n\n")
	n:= 1==1
	m:= 1==2
	fmt.Printf("%v, %T\n",n,n)
	fmt.Printf("%v, %T\n",m,m)
	fmt.Println("\n\n")

	// INTEGERS
	fmt.Println("INTEGER TYPES\n\n")
	i:= 1
	var j uint16 = 1
	fmt.Printf("%v, %T\n",i,i)
	fmt.Printf("%v, %T\n",j,j)

	//arithmetic operations
	a:=8
	b:=3
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
	
	//bit operations
	
	fmt.Println(a & b)		//and
	fmt.Println(a | b)		//or
	fmt.Println(a ^ b)		//xor
	fmt.Println(a &^ b)		//and not

	// left and right shifts
	fmt.Println( a>>3)		//right shift /2*3
	fmt.Println( a<<3)		// left shift *2^3

	fmt.Println("\n\n")

	// FLOATING POINT VARIABLES
	fmt.Println("FLOATING POINT TYPES\n\n")

	var c float64 = 3.14
	c = 13.7e72
	c = 2.1E14

	fmt.Printf("%v , %T",c,c)

	fmt.Println("\n\n")

	//COMPLEX NUMBERS
	fmt.Println("FLOATING POINT TYPES\n\n")

	var k complex64 = 1 + 2i
	fmt.Printf("%v , %T\n",k,k)
	fmt.Printf("%v , %T\n",real(k),imag(k))

	var d complex128 = complex(5,12)
	fmt.Printf("%v , %T\n",real(d),imag(d))
	fmt.Println("\n\n")

	//STRINGS
	fmt.Println("STRINGS\n\n")
	s:= "this is a string"
	s2:= "this is also a string"
	fmt.Printf("%v , %T\n",s,s)
	fmt.Printf("%v , %T\n",s[2],s[2])
	fmt.Printf("%v , %T\n",string(s[2]),string(s[2]))
	fmt.Printf("%v , %T\n",s+s2,s+s2)					// string concatenation

	g := []byte(s)
	fmt.Printf("%v , %T\n",g,g)							// converting a string to a slice of bytes

	fmt.Println("\n\n")

	//RUNES
	fmt.Println("RUNES\n\n")
	r := 'a'
	fmt.Printf("%v , %T",r,r)
}