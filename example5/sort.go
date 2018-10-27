package main
import(
	"fmt"
)

func bublesort(a []int) {
	//冒泡排序
	for i := 0; i < len(a); i++{
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j],a[j-1] = a[j-1],a[j]
				fmt.Println(a)
			}
		}
	}
}

func selectSort(a []int) {
	//选择排序
	for i := 0; i < len(a); i++{
		for j := i+1; j < len(a); j++{
			if a[i] > a[j]{
				a[i],a[j] = a[j],a[i]
			}
		}
	}
	fmt.Println(a)
}

func InsertSort(a []int) {
	//插入排序
	for i := 1; i < len(a); i++{
		for j := i; j > 0; j--{
			if a[j] < a[j-1] {
				break
			}
			a[j],a[j-1] = a[j-1],a[j]
		}
	}
	fmt.Println(a)
}

func quickSort(a []int,left,right int) {
	//快速排序
	if left >= right {
		return
	}
	val := a[left]
	//确定val所在的位置
	k := left
	for i := left+1; i <= right; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
			a[k] = val
		}
		fmt.Println("quicksork",a)
	}
	quickSort(a,left,k-1)
	quickSort(a,k+1,right)
	
}

func main() {
	a := [...]int{8,7,6,5,4,3,2,1}
	bublesort(a[:])
	fmt.Println(a)
	selectSort(a[:])
	InsertSort(a[:])
	quickSort(a[:],0,len(a)-1)
}