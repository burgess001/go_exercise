package main
import (
	"fmt"
	"sort"
)
/* map:key-value的数据结构，又叫字典或关联数组
声明map：
var map1 map[keytype]valuetype 例如：var a map[string]string  var a[staring]int
var a map[string]map[string]string
声明是不会分配内存的，初始化需要make
map是引用类型
*/
func testMap() {
	var a map[string]string
	a = make(map[string]string,10)
	a["a"] = "fgh"
	a["a"] = "cfg"
	a["b"] = "cfg"
	//map相关操作
	a["hello"] = "world"  //插入一个值
	val ,ok := a["hello"] //查询一个值
	if ok {
		fmt.Println(val)
	}else{
		a["hello"] = "world"
	}
	delete(a,"hello") //删除
	for k,v := range a {  //遍历map
		fmt.Println(k,v)
	}
	fmt.Println(a)
}
func testsortMap() {
	//map排序，map里的key都是无序的
	var a map[int]int
	a = make(map[int]int)
	a[1] = 10
	a[5] = 3
	a[3] = 4
	a[89] = 1
	a[10] = 99
	fmt.Println("map排序前：")
	for k,v := range a {
		fmt.Println(k,v)
	}
	fmt.Println("map排序后：")
	var keys []int
	for k,_ := range a {
		keys = append(keys,k)
	}
	sort.Ints(keys)
	for _,v := range keys {
		fmt.Println(v,a[v])
	}
}
func testMapsort1() {
	//map 反转 key变成value
	var a map[string]int
	var b map[int]string
	a = make(map[string]int,5)
	b = make(map[int]string)
	a["abc"] = 100
	a["vfg"] = 200
	fmt.Println("反转前：",a)
	for k,v := range a {
		b[v] = k
	}
	fmt.Println("反转后：",b)
}
func main() {
	a := make(map[string]map[string]string,100)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	fmt.Println(a)
	testMap()
	testsortMap()
	testMapsort1()
}