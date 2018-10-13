package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"sync/atomic"  //原子操作
)
/*包
线程同步import("sync")

互斥锁 var mu sync.Mutex
读写锁 var mu sync.RWMutex
*/
var lock sync.Mutex
var rwlock sync.RWMutex
func testMap() {
	/*用两个goroutine 修改map的值，这样虽然不会报错，但是存在互斥性有竞争关系，在bulid的时候
	加上-race参数可以启用数据竞争检测，加上-race参数 go build -race xx.go完之后,再次执行生成的可执行文件
	对于有竞争关系的变量，如果不加锁就会报错。
	*/
	var a map[int]int
	a = make(map[int]int,5)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[10] = 10
	for i := 0; i < 2; i++{
		go func(b map[int]int){
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

}
func testRWloack() {
	var a map[int]int
	a = make(map[int]int,5)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[10] = 10
	for i := 0; i < 2; i++{
		go func(b map[int]int){
			rwlock.Lock()
			b[8] = rand.Intn(100)
			rwlock.Unlock()
		}(a)
	}
	for i := 0; i < 100; i++{
		go func(b map[int]int) {
			rwlock.RLock()
			fmt.Println(a)
			rwlock.RUnlock()
			atomic.AddInt32(&count,1)
		}(a)
	}
	time.Sleep(time.Second*100)
	fmt.Println(atomic.LoadInt32(&count))
}
func main() {
	//testMap()
	testRWloack()
}