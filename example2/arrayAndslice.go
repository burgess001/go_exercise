package main
import (
	"fmt"
)
/*
数组：同一种数据类型的固定长度序列
数组的定义：var a [len]int ,比如 var a[5]int
长度是数组类型的一部分，因此，var a [5]int 和 var a [10]int 是不同的类型
*/
/*
切片：切片是数组的一个引用，因此切片是引用类型的
切片的长度是可变的，因此切片是一个可变的数组
切片的遍历方式和数组一样，可以用len()求长度
cap可以求出slice的最大容量，0 <= len(slice) <= cap(array),其中array是slice引用的数组
切片的定义：var 变量名 []类型，比如 var str []string    var arr []int
*/
func chushi(){
	//数组的初始化
	var age0 [5]int = [5]int{1,2,3}
	var age1 = [5]int{1,2,3,4,5}
	var age2 = [...]int{1,2,3,4,5,6}
	var str = [5]string{3:"hello world",4:"john"} //指定数组的第三个和第四个元素
	fmt.Println(age0)
	fmt.Println(age1)
	fmt.Println(age2)
	fmt.Println(str)
}
func duowei(){
	//多维数组  定义一个三行四列的二维数组
	var  a [3][4]int = [...][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	for _,v:= range a{
		for _,v1 := range v{
			fmt.Printf("%d\t",v1)
		}
		fmt.Printf("\n")
	}
}

func silce(){
	var a = [10]int{1,2,3,4}
	b := a[1:5]
	fmt.Println(b)
	var c = []int{1,2,3}
	var d = []int{4,5,6,7}
	c = append(c,d...) //切片append切片
	fmt.Println(c)
	//切片的拷贝，内置函数copy
	s1 :=[]int{1,2,3,4,5}
	s2 := make([]int,10)
	copy(s2,s1)
	fmt.Printf("s2:%d\n",s2)

}
func testmodifystring () {
	//slice与string的关系
	/*string本身是不可变的，因此要改变string中的字符串，需要如下操作*/
	str := "hello world"
	s := []byte(str)
	s[0] = '0'
	str = string(s)
	fmt.Println(str)
	//但是字符串如果是中文转换成byte类型的数组就出错了，应该用
	s1 := "我爱你，中国"
	s2 := []rune(s1)
	s2[0] = '你'
	s1 = string(s2)
	fmt.Println(s1)
}
func fibo(n int){
	//用slice非递归的实现Fibonacci数列
	var a [] int64
	a = make([]int64,n)
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++{
		a[i] = a[i-1] +a[i-2]
	} 
	for _,v := range a{
		fmt.Println(v,)
	}
}
func main (){
	var a [10]int
	a[0] = 100
	fmt.Println(a)
	for i := 0; i < len(a); i++{
		fmt.Println(a[i])
	}
	for index,val := range a {
		fmt.Printf("a[%d]=%d\n",index,val)
	}
	fibo(10)
	chushi()
	duowei()
	silce()
	testmodifystring()
}