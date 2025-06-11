package main
import "fmt"



func findMax(a []int) int{
	max:=0;

	for i:=0;i<len(a);i++{
		if a[i]>max{
			max=a[i]
		}
	}
return max
}

func main(){
	// basic go array 

	var a [3] int
	a[0]=10
	a[1]=20
	a[2]=30

	fmt.Println(a)


	for i:=0; i<3; i++{
		fmt.Println(a[i])
	}


	// other initializati)on shortcuts
	
	b:=[3]int{1,2,3}
	c:=[...]int{4,5,6,7,8,9,10}


	fmt.Println(b)
	fmt.Println(c)



	// creating a string array having my favourite cricketers name


	names:=[...]string{"Virat","AB DE Villiers","Dale Steyn","Brendon Mcculum"}


	fmt.Println(names)


	names[1]="ABD"
// after changing the second element 
	fmt.Println(names,len(names))
	maximum:= findMax(c[:]) // here we are first converting the array c into a slice as the function findMax(a []int) -> expects a slice 
	fmt.Println(maximum)
}
