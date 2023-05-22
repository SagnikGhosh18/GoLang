package main

import(
	"fmt"
)

type greeter struct{
	greeting string
	name string
}
func main(){
	sayMessage("Hyyyy")
	sayGreeting("Hi","Stacey")

	s := sum(1,2,3,4,5)
	fmt.Println(*s)
	fmt.Println(sum1(1,2,3,4,5))

	a,b := 5.0,0.0
	d,err := div(a,b)
	fmt.Println(d, err)

	var f func(int, int) int = func (a,b int) int{
		return a+b
	}
	fmt.Println(f(2,3))


	g:= greeter{
		"Hello",
		"Go",
	}
	g.greet()
}

func sayMessage(msg string){
	fmt.Println(msg)
}

func sayGreeting(greet, name string){
	fmt.Println(greet,name)
}

func sum(values ...int) *int{
	fmt.Printf("%v, %T\n",values,values)
	result := 0
	for _,v := range values{
		result+=v
	}
	fmt.Println(result)
	return &result
}
func sum1(values ...int) (result int){
	fmt.Printf("%v, %T\n",values,values)
	for _,v := range values{
		result+=v
	}
	return
}

func div(a,b float64) (float64, error){
	if b==0.0 {
		return 0.0,fmt.Errorf("Cannot divide by zero")
	}
	return a/b,nil
}

func (g greeter) greet(){
	fmt.Println(g.greeting,g.name)
}