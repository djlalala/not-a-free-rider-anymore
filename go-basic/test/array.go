package main

import "fmt"

/***
	1. 数组：是同一种数据类型的固定长度的序列。
    2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
    3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
    4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
    for i := 0; i < len(a); i++ {
    }
    for index, v := range a {
    }
    5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
    6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
    7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
    8.指针数组 [n]*T，数组指针 *[n]T。
***/
func main() {
	var a [3]int = [3]int{123, 234, 345}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for index, v := range a {
		fmt.Print(index)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println("")
	}

	fmt.Println(a, a, a)

	// var val [一维][二维][三维]类型
	var mutiArr [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	fmt.Println(mutiArr)
	fmt.Println("当前多维数组长度是:", len(mutiArr))
	fmt.Println("当前多维数组0元素长度是:", len(mutiArr[0]))

	// 多维数组遍历
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}

	// 数组拷贝和传参
	// &和*用于指针和值的转换
	// 可以发现当前printArr的参数为指针类型的int数组，当调用该方法，传递内存地址，方法中改变了值，
	// 回到调用的main方法中内存中的值同样也发生了改变。
	var arr1 [5]int
	printArr(&arr1)
	// 打印内存地址
	fmt.Printf("%p", &arr1)
	fmt.Println("\n", arr1)
}

func printArr(arr *[5]int) {
	fmt.Println("printArr方法开始执行:", arr)
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
