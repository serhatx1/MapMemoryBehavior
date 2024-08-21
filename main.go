package main

import (
	"fmt"
	"runtime"
)

//func main() {
//	n := 10_000_000
//	m := make(map[int][128]byte)
//	printAlloc("Right after map allocation(without pointer)")
//	// Right after map allocation(without pointer): 0 MB
//
//	for i := 0; i < n; i++ {
//		m[i] = [128]byte{}
//	}
//	printAlloc("Right after 10 million elements added(without pointer)")
//	// Right after 10 million elements added(without pointer): 3694 MB
//
//	for i := 0; i < n; i++ {
//		delete(m, i)
//	}
//
//	runtime.GC()
//	printAlloc("After deleting 10 million elements and running GC(without pointer)")
//	// After deleting 10 million elements and running GC(without pointer): 2347 MB
//
//	runtime.KeepAlive(m)
//
//}

func main() {
	n := 10_000_000
	m := make(map[int]*[128]byte) // with pointers
	printAlloc("Right after map allocation(With pointer declare(with pointer declare)")
	//Right after map allocation(With pointer declare(with pointer declare): 0 MB
	for i := 0; i < n; i++ {
		m[i] = &[128]byte{}
	}
	printAlloc("Right after 10 million elements added(with pointer declare)")
	//Right after 10 million elements added(with pointer declare): 1701 MB
	for i := 0; i < n; i++ {
		delete(m, i)
	}
	runtime.GC()
	printAlloc("After deleting 10 million elements and running GC (with pointer)")
	//After deleting 10 million elements and running GC (with pointer): 306 MB
	runtime.KeepAlive(m)

}
func printAlloc(step string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: %d MB\n", step, m.Alloc/(1024*1024))
}
