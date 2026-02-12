package main

import "fmt"

type Person struct {
	name string
	age  int

}


func main(){
	
	p1 := Person{name: "Alice", age: 30}

	var a,b int = 10, 20
	const c, d int = 30, 40
	const (
	e, f int = 50, 60
	g, h int = 70, 80
	)

	fmt.Printf("a: %d, b: %d\n", a, b)
	fmt.Printf("c: %d, d: %d\n", c, d)
	fmt.Printf("e: %d, f: %d\n", e, f)
	fmt.Printf("g: %d, h: %d\n", g, h)
	fmt.Printf("Person name %s and age %d",p1.name ,p1.age)

	

	
}
