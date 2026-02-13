package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Print("---Maps---")
	m1:=make(map[int]string)
	m1[0]="a"
	m1[1]="b"
	m1[2]="c"
	m1[3]="d"
	fmt.Println("m1:",m1)
	//accessing non existing key in map
	//fmt.Println("access :",m1[11]) -->print empty value 
	//IMP; if key value does not exist in map and you access it , it return zero value for int, string etc

	//deleting any key value from map , pass key in delete function
	delete(m1,3)
	fmt.Println("after deleting 3 now m1:",m1)

	//deleting all element 
	clear(m1)

	fmt.Println("after deleting all keys now length m1:",len(m1))


	//other way to make map  
	m3:= map[string]int {"a":1,"b":2}
	fmt.Println("m3: ",m3)

	k,ok := m3["a"] //ok is true if key exists in map and false if key does not exist in map
	fmt.Println("value of key a is ",k,"and ok is ",ok)
	if ok{
		fmt.Println("key exists in map and value is ",k)
	}else{
		fmt.Println("key does not exist in map")
	
	}

	//operator in map
	m4:= map[string]int {"a":1,"b":2}
	m5:= map[string]int {"a":1,"b":1}
	fmt.Println("compare m4 and m5 => ",maps.Equal(m4,m5)) //map cannot be compared using == operator except for nil map

}