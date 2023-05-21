package main

import(
	"fmt"
)

type myStruct struct{
	foo int
}

func main(){
	var a int= 42
	var b *int = &a
	fmt.Println(a,b)
	fmt.Println(a,*b)
	a = 27
	fmt.Println(a,*b)
	*b = 14
	fmt.Println(a,*b)

	arr := [3]int{1,2,3}
	i := &arr[0]
	j := &arr[1]
	fmt.Println("%v, %p, %p",arr,i,j)

	var ms *myStruct
	ms = new(myStruct)						//initializes an empty object
	fmt.Println(ms)
	(*ms).foo = 42
	fmt.Println((*ms).foo)
	ms.foo = 34								//also valid
	fmt.Println(ms.foo)

	
}