package main
import (
	"fmt"
	"math/rand"
)
type Student struct{
	Name string
	Age int
	Score float32
	next *Student
}
func trans (p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}
func Inserttail(p *Student) {
	//批量从尾部插入
	var tail = p
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name : fmt.Sprintf("stu%d",i),
			Age : rand.Intn(100),
			Score : rand.Float32() *100,
		}
		tail.next = stu
		tail = stu
	}
}

func Inserthead(p **Student) {
	for i := 0; i < 10; i++{
		stu := &Student{
			Name : fmt.Sprintf("stu%d",i),
			Age : rand.Intn(100),
			Score : rand.Float32() *100, 
		}
		stu.next = *p
		*p = stu
	}
}

func deleteNode(p *Student) {
	var prev *Student = p
	for p != nil {
		if p.Name == "stu2"{
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}
func InsertNode (p *Student,newNode *Student) {
	for p != nil {
		if p.Name == "stu3" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}

func Insertmy (p *Student) {
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name : fmt.Sprintf("stu%d",i),
			Age : rand.Intn(100),
			Score : rand.Float32() *100,
		}
		p.next = stu
		p = stu
	}
}
func main() {
	var head *Student = new(Student)
	head.Name = "jw"
	head.Age = 18
	head.Score = 100
	var stu1 Student
	stu1.Name = "stu1"
	stu1.Age = 20
	stu1.Score = 80
	head.next = &stu1
	//遍历链表
	trans(head)
	//给链表再插入一个值
	var stu2 Student
	stu2.Name = "stu2"
	stu2.Age = 23
	stu2.Score = 23
	stu1.next = &stu2 //这种方式我们称为尾部插入
	trans(head)

	fmt.Println("尾插")
	var head1 *Student = new(Student)
	Inserttail(head1)
	trans(head1)

	//头插
	var head2 *Student = new(Student)
	head2.Name = "kou"
	head2.Age = 19
	head2.Score = 60
	fmt.Println("头插")
	Inserthead(&head2)
	trans(head2)

	//删除一个节点
	deleteNode(head2)
	trans(head2)

	//插入一个节点
	var newNode = &Student{
		Name: "ji",
		Age: 19,
		Score: 90.1,
	}
	fmt.Println("中间插入一个节点")
	InsertNode(head1,newNode)
	trans(head1)

	//自己胡写的
	fmt.Println("自己胡写")
	var myNode = &Student{
		Name: "wen",
		Age: 19,
		Score: 90.1,
	}
	Insertmy(myNode)
	trans(myNode)
}