package main

import "fmt"

func main(){
	//Arrays
	var arr [5]int 
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println("Array:", arr)
	//Array literals
	/*
	###Important thing : If you initialize an array with fewer elements than its size, the remaining elements will be set to their zero value (0 for integers, "" for strings, false for booleans, etc.).
	*/
	arr4:= [10]string{"Go", "is", "a", "great", "language", "!", "I", "love", "it", "."}
	arr2 := [5]int{1, 2}
	fmt.Println("Array2 literal:", arr2)
	fmt.Println("Array4 literal:", arr4)
	//Slices

    slice2:= [] int {1,2,3,4,23,34,4}
	//slice vs Array is that slice is dynamic and array is fixed size
	slice := []int{10, 20, 30}
	fmt.Println("Slice2 literal:", slice2)
	fmt.Println("Slice:", slice)

	//2d array
	twoDArray := [2][3] int {{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D Array:", twoDArray)
}