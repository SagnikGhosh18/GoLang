package main

import(
	"fmt"
)

func main(){
	for i:=0;i<5;i++ {
		fmt.Println(i)
	}
	for a,b:=0,0;a<5;a,b=a+1,b+1 {
		fmt.Println(a," ",b)
	}
	j:=0
	for ;j<5;j++{
		fmt.Println(j)
	}
	j=0
	for ;j<5; { 			// this is the go equivalent of a while loop
		fmt.Println(j)
		j++;
	}

	i := 0
	for i<5 {
		fmt.Println(i)
		i++
	}

	//Go equivalent of do while loop
	i = 0
	for {
		fmt.Println(i)
		i++
		if i>5{
			break
		}
		if i==2{
			continue
		}
	}
	//loop labels
	i=1
	j=1
Loop:
	for ;i<5;i++{
		for ;j<5;j++{
			fmt.Println(i*j)
			if i*j > 3{
				break Loop
			}
		}
	}

	//loops with slices, arrays, maps and strings
	s:= []int{1,2,3}
	for k ,v := range s{
		fmt.Println(k,v)
	}
	for k  := range s{			// to access values, we can use _
		fmt.Println(k)
	}
}
