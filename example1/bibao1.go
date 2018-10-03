package main

import (
	"fmt"
	"strings"
)
func makeSuffixFunc(suffix string) func (string) string{
	return func (name string) string{
		if strings.HasSuffix(name,suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	f1 := makeSuffixFunc(".png")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))
	f2 := makeSuffixFunc(".jpg")
	fmt.Println(f2("test"))
	fmt.Println(f2("pic"))
}