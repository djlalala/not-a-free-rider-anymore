// author 董健 <joeyana@aliyun.com>
package main

import (
	"fmt"
	"unsafe"
)

/***
	Slice 的底层实现
		| 2.切片的数据结构
		| 3.创建切片
		| 4.切片扩容
		| 5.切片拷贝
***/
func main() {
	// 2.切片的数据结构
	// 切片的结构体由3部分构成，
	// Pointer 是指向一个数组的指针，len 代表当前切片的长度，cap 是当前切片的容量。cap 总是大于等于 len 的。
	fmt.Println("------------------------------------------------------------")
	fmt.Println("切片的数据结构:")
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	// 3.创建切片 TODO
	// nil切片和空切片 var slice []int
	// silce := make([]int, 0)
	// slice := []int{}
	fmt.Println("------------------------------------------------------------")
	fmt.Println("创建切片:")

	// 4.切片扩容,扩容策略 TODO
	// 如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。
	// 上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。
	// 一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。
	// 注意：扩容扩大的容量都是针对原来的容量而言的，而不是针对原来数组的长度而言的。
	fmt.Println("------------------------------------------------------------")
	fmt.Println("切片扩容策略:")
	sa := make([]int, 0, 1)
	ca := cap(sa)
	for i := 0; i < 2000; i++ {
		sa = append(sa, i)
		if n := cap(sa); n > ca {
			fmt.Printf("cap: %d -> %d\n", ca, n)
			ca = n
		}
	}

	// 5.切片拷贝
	// 由于 Value 是值拷贝的，并非引用传递，所以直接改 Value 是达不到更改原切片值的目的的，
	// 需要通过 &slice[index] 获取真实的地址。
	fmt.Println("------------------------------------------------------------")
	fmt.Println("切片拷贝:")
	var arr [3]int = [3]int{88, 99, 100}
	fmt.Println("arr[0]的真实地址:", &arr[0])
	s := arr[:]
	fmt.Println("&slice[0]的真实地址:", &s[0])
	s[0] = 1
	fmt.Println(arr, s)

}
