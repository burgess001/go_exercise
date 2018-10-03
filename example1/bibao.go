package main

import (
	"fmt"
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

func main() {
	var f1 = Add()
	var f2 = Add()
	fmt.Println(f1(1))
	fmt.Println(f1(10))
	fmt.Println(f2(1))
	fmt.Println(f2(10))
}
