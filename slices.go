package main

import "fmt"

func main(){
	var a [3]int // a normal array
	for i:=0;i<len(a);i++{
		a[i]=(i+1)
	}
	fmt.Println(a)

	// a slice in go

	slice:=[]int{1,2}
	// before appending 3
	fmt.Println(slice)


	// appending elements in a slice

	slice=append(slice,2,3,4,5,6,7)
	// after appending 2,3
	fmt.Println(slice)
}
