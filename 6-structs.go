package main

import (
	"fmt"
	"reflect"
)

type doctor struct{
	number int
	name string
	companions []string
}

type Animal struct{
	name string `required max:"100"`			//these are tags, which are space delimited key value pairs
	origin string
}

type Bird struct{
	Animal
	speed int
	canFly bool
}

func main(){
	aDoctor := doctor{
		5,
		"A name",
		[]string{"abc","cde","fgh"},		// must be in order
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.name)
	fmt.Println(aDoctor.companions[0])

	anotherDoc := struct{name string}{name:"A"}
	fmt.Println(anotherDoc)
	fmt.Println(anotherDoc.name)

	//Embedding structs
	b := Bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speed = 55
	b.canFly = false
	fmt.Println(b)


	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}