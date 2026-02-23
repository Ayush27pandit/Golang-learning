package main

import "fmt"

func main() {
	fmt.Println("--------Ranges---------")
	nums:=[]int {1,3,4,5,6,6,7,8,9}
	fmt.Println(nums)
	// for i:=range len(nums){
	// 	fmt.Println(nums[i])
	// }
	sum:=0;

	for i,nums:=range nums{
		
		sum+=nums
		fmt.Println("Sum in iterations",i,sum)
	}

	m1:=map[int]string{11:"a",12:"b",3:"c",4:"d"}
	fmt.Println(m1)
	for k,s:=range m1{
		fmt.Println(" key-",k ," => ",s)
	}

}