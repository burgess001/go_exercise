package main
import (
	"fmt"
)

type Student struct{
	Name string
	Age int
	score float32
}
func main () {
	var stu Student
	stu.Name = "kou"
	stu.Age = 18
	stu.score = 90
	fmt.Println(stu)
	fmt.Printf("Name:%p\n",&stu.Name)
	fmt.Printf("Age:%p\n",&stu.Age)
	fmt.Printf("socre:%p\n",&stu.score)
	var stu1 *Student = &Student{
		Name: "lo",
		Age: 20,
	}
	fmt.Printf("Name:%p\n",&stu1.Name)
	fmt.Printf("Age:%p\n",&stu1.Age)
	fmt.Printf("score:%p\n",&stu.score)
}