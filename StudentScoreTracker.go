package main

import "fmt"


type Student struct{
	Name string 
	Score int 
}



func addStudent(students Student[], name string,score int) [] Student{
	newStudent:=Student{Name:name, Score:score}

	return append(students,newStudent)
}







func main(){
	var students []Student

	students:=addStudent(students,"Vikas",100)
	students:=addStudent(students,"abc",98)





	for _,s := range students{
		fmt.Printf("name: %s, Score: %d\n",s.Name,s.Score)
	}
}
