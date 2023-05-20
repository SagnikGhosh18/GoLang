package main

import(
	"fmt"
)

func main(){
	//grades := [3]int{97,85,93}
	grades := [...]int{97,85,93}

	var students [3]string

	fmt.Printf("%v\n",grades)
	fmt.Printf("%v\n",students)
	students[0] = "Lisa"
	students[1] = "B"
	students[2] = "C"
	fmt.Printf("%v\n",students)
	fmt.Printf("%v\n",students[1])
	fmt.Printf("%v\n",len(students))

	var matrix [3][3]int = [3][3]int{[3]int{1,0,0},[3]int{0,1,0},[3]int{0,0,1}}
	fmt.Printf("%v\n",matrix)

	var matrix2 [3][3]int
	matrix2[0] = [3]int{1,1,1}
	matrix2[1] = [3]int{1,1,1}
	matrix2[2] = [3]int{1,1,1}
	fmt.Printf("%v\n",matrix2)

	matrix3 := matrix2							/* go makes a copy of the array instead of pointing to it, - the way to bypass 
													this is pointers */
	matrix3[2] = [3]int{0,0,0}
	fmt.Printf("%v\n",matrix3)

	fmt.Println("\n\n\n")
	//SLICES 

	a := []int{1,2,3}							/* for slices however, the reference is pointed, so the value change is 
													reflected */
	b := a
	b[1] = 5
	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",b)
	fmt.Printf("Length: %v\n",len(a))
	fmt.Printf("Capacity: %v\n",cap(a))

	a = []int{1,2,3,4,5,6,7,8,9,10}
	b = a[:]									// these operations are possible in arrays too
	c := a[4:]
	d := a[:6]
	e := a[2:9]
	fmt.Printf("%v\n",b)
	fmt.Printf("%v\n",c)
	fmt.Printf("%v\n",d)
	fmt.Printf("%v\n",e)


	f := make([]int, 3, 100)					// int length capacity (passing only 1 makes length=cap)
	
	fmt.Printf("%v\n",f)
	fmt.Printf("Length: %v\n",len(f))
	fmt.Printf("Capacity: %v\n",cap(f))

	g := []int{}
	fmt.Printf("%v\n",g)
	fmt.Printf("Length: %v\n",len(g))
	fmt.Printf("Capacity: %v\n",cap(g))
	g = append(g,1)
	fmt.Printf("%v\n",g)
	fmt.Printf("Length: %v\n",len(g))
	fmt.Printf("Capacity: %v\n",cap(g))
	g = append(g,2,3,4,5)
	fmt.Printf("%v\n",g)
	fmt.Printf("Length: %v\n",len(g))
	fmt.Printf("Capacity: %v\n",cap(g))
	g = append(g, []int{6,7,8,9}...)
	fmt.Printf("%v\n",g)
	fmt.Printf("Length: %v\n",len(g))
	fmt.Printf("Capacity: %v\n",cap(g))


	h := g[:len(a)-1]
	i := append(a[:2],a[3:]...)
	fmt.Printf("%v\n",h)
	fmt.Printf("%v\n",i)
}