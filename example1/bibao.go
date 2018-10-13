package main

import (
	"fmt"
	"strings"
)

/*
闭包：一个函数与其相关的引用环境组合而成的实体
*/
func Add() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
func makeSuffixFunc(suffix string) func (string) string{
	return func (name string) string{
		if strings.HasSuffix(name,suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	var f1 = Add()
	var f2 = Add()
	fmt.Println(f1(1))
	fmt.Println(f1(10))
	fmt.Println(f2(1))
	fmt.Println(f2(10))
	f3 := makeSuffixFunc(".png")
	fmt.Println(f3("test"))
	fmt.Println(f3("pic"))
	f4 := makeSuffixFunc(".jpg")
	fmt.Println(f4("test"))
	fmt.Println(f4("pic"))
}
