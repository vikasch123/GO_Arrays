package main

import "fmt"

func main(){


	var fruitList= []string{}
	fmt.Printf("Type of fruitList is %T \n",fruitList)


var vegList = make([]string,4)
vegList[0]="potato"
vegList[1]="beans"
vegList[2]="ladyfinger"
vegList[3]="Brocolli"

var index = 2
// we want to remove ladyfinger from vegList

vegList=append(vegList[:index],vegList[index+1:]...)

fmt.Println(vegList)

// creating a map

var marks = make(map[string]int)

marks["English"]=92
marks["Maths"]=98
marks["Computer Science"]=100
marks["Physics"]=91

fmt.Println(marks)

// creating an another map

newmap := map[string]int{"a":1, "b":2}


for _,v := range newmap{
	fmt.Printf("the value for key is %v\n",v)

}






}
