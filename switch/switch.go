package main

import (
	"fmt"
	"time"
)

func main() {
	//Switch 

	switch day := 12; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")

		
	
	}

	if time.Now().Weekday()== time.Saturday || time.Now().Weekday()== time.Sunday {
		fmt.Println("It's the weekend!")
	}else {
		fmt.Println("It's a weekday!" , time.Now().Format("Monday , 2 January 2006 15:04:05"))
	
	}

	whoAmI := func (i interface{}){
		switch i:=i.(type){
		case int:
			fmt.Println("Integer",i)
		case string:
			fmt.Println("String",i)
		case bool:
			fmt.Println("Boolean",i)
		default:
			fmt.Println("Unknown",i)
		
		}
	}
	whoAmI(42)
	whoAmI("Hello")
	whoAmI(true)
}