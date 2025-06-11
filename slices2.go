package main

import "fmt"

func main(){
	// here we are trying to define a slice and input elements in the slice and we don't know the number of elements

	var s[]int //declared the slice
	var val int 
	var choice string

	for{
		fmt.Println("Enter the number:")
		fmt.Scanln(&val)
		s=append(s,val)


		fmt.Print("Add more? (y/n) : ")
		fmt.Scanln(&choice)
		if choice!="y"{
			break
		}

	}
	fmt.Println("Final slice : ",s)
}
