package main

import "fmt"

func main(){
	var a [3]int
	var n int 
	// a normal array
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

	// creating a slice using make()
	fmt.Println("Enter the number of elements : ")

	fmt.Scanln(&n)
	s:=make([]int,n) //[0 0 0 0 0]
	// appending elements
	//var  elem int

	fmt.Println("Enter elements into the slice")
	for i:=0;i<n;i++{
		fmt.Scanln(&s[i])
		
	}

	fmt.Println("The slice is printed below")

	fmt.Println(s)

}
