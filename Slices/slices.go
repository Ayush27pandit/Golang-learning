package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []int //uninitialized slice has nil value
	fmt.Println("Slice:", s)
	fmt.Println("Is slice nil?", s == nil)
    fmt.Println("Length of slice", len(s))


	/*
	Creating a slice using Make function
	Important to note 
	syntax: make([]Type, intital length, capacity) --> diff in capacity and length is that length is the number of elements in the slice and capacity is the total number of elements that can be stored in the slice before it needs to be resized. If you don't specify capacity, it will be set to the initial length.
	Below example
	*/
	sliceMake:=make([]int, 3,7)
	fmt.Println("Slice created using make:", sliceMake)
	fmt.Println("Length of slice created using make:", len(sliceMake))
	fmt.Println("Capacity of slice created using make:", cap(sliceMake))
	sliceMake = append(sliceMake, 10, 20, 30, 40,90,9100) //appending more elements than capacity to see how it resizes
	fmt.Println("Slice created using make:", sliceMake)
	fmt.Println("Capacity of current slice after appending more element:", cap(sliceMake))//which becomes 14 as it doubles the capacity when it needs to resize
	fmt.Println("Length of current slice after appending more element:", len(sliceMake))


	//Copy feature
	nums:=make([]int ,0,3)
	nums= append(nums,1,2,4)
	nums2:=make([]int,len(nums),cap(nums))

	copy(nums2,nums)
	fmt.Println("nums  is", nums)
	fmt.Println("nums2 is copy of nums so nums2 is", nums2)

	//slice operator

	nums3:=[]int {0,1,2,3,4,5,6,7,8,9}
	//[start from this index: one element before this index]
	fmt.Println(nums3[6:])

	//Compare slices

	nums4:=[]int {0,1,2,3,4,5,6,7,9,8}
	fmt.Println("Compare nums4 and nums3 => ", slices.Compare(nums3,nums4))//0 means equal and 1 means not equal, 1 means S1>S2 and -1 means S1<S2
	

	//2d slices

	numsMatrix:=[][]int {{1,2,3},{1,2,3}}
	fmt.Println("2d slices",numsMatrix)




}