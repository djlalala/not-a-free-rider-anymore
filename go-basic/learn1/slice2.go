// author 董健 <joeyana@aliyun.com>
package main

import (
	"fmt"
	tools "go-basic/tools"
	"os"
	// tools "go-basic/tools"
)

func init() {
	tools.Hello()
	fmt.Println("------------------------------------------------------------")
	dir, _ := os.Getwd()
	fmt.Println("当前路径:", dir)
	fmt.Printf("当前位置: %v() \n", tools.GetFunctionName(main))
	fmt.Println("Slice底层实现:")
}

/***
	Slice底层实现
		| 1.切片和数组
		| 2.切片的数据结构
		| 3.创建切片
		| 4.切片扩容
		| 5.切片拷贝

***/
func main() {

	/***
	Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据。
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("切片和数组:")
	fmt.Println("1.传数组")
	arrayA := [2]int{100, 200}
	var arrayB [2]int
	arrayB = arrayA
	fmt.Printf("arrayA : %p , %v, %d, %d\n", &arrayA, arrayA, len(arrayA), cap(arrayA))
	fmt.Printf("arrayB : %p , %v, %d, %d\n", &arrayB, arrayB, len(arrayA), cap(arrayB))
	printArrayInfo1(arrayA)
	fmt.Println("结论:数组赋值和函数传参，都会复制整个数组。所以三个对应的内存地址也不同。")

	/***
		假想每次传参都用数组，那么每次数组都要被复制一遍。
		如果数组大小有 100万，在64位机器上就需要花费大约 800W 字节，即 8MB 内存。
		这样会消耗掉大量的内存。于是乎有人想到，函数传参用数组的指针。

		结论:
		就算是传入10亿的数组，也只需要再栈上分配一个8个字节的内存给指针就可以了。
		这样更加高效的利用内存，性能也比之前的好。
		不过传指针会有一个弊端，从打印结果可以看到，第一行和第三行指针地址都是同一个，
		万一原数组的指针指向更改了，那么函数里面的指针指向都会跟着更改。
	***/
	fmt.Println("2.传数组指针")
	array1 := [2]int{100, 200}
	var array2 *[2]int
	array2 = &array1
	fmt.Printf("array1 : %p , %v\n", &array1, array1)
	fmt.Printf("array2 : %p , %v\n", &array2, *array2)
	printArrayInfo2(&array1)
	// 模拟改变原数组指针
	*array2 = [2]int{88, 99}
	fmt.Printf("array1 : %p , %v\n", &array1, array1)
	fmt.Printf("array2 : %p , %v\n", &array2, *array2)
	printArrayInfo2(&array1)
	fmt.Println("结论:就算是传入10亿的数组，也只需要再栈上分配一个8个字节的内存给指针就可以了。\n这样更加高效的利用内存，性能也比之前的好。\n不过传指针会有一个弊端，从打印结果可以看到，第一行和第三行指针地址都是同一个，\n万一原数组的指针指向更改了，那么函数里面的指针指向都会跟着更改。")

	/***
		3.传切片,切片的优势也就表现出来了。
		用切片传数组参数，既可以达到节约内存的目的，也可以达到合理处理好共享内存的问题。
		切片的指针和原来数组的指针是不同的。

		结论:
		把第一个大数组传递给函数会消耗很多内存，采用切片的方式传参可以避免上述问题。
		切片是引用传递，所以它们不需要使用额外的内存并且比使用数组更有效率。
	***/
	fmt.Println("3.传切片")
	arrayX := [2]int{100, 200}
	arrayY := arrayX[:]
	fmt.Printf("arrayX : %p , %v\n", &arrayX, arrayX)
	fmt.Printf("arrayY : %p , %v\n", &arrayY, arrayY)
	printArrayInfo3(&arrayY)
	fmt.Println("结论:用切片传数组参数，既可以达到节约内存的目的，也可以达到合理处理好共享内存的问题。切片的指针和原来数组的指针是不同的。")
	/***
		总的结论:
		1.传数组，都是拷贝
		2.传指针，只需一个指针单内存大小，但原数组的指针指向更改，函数里面的指针也会改变。
		3.传切片，切片是引用传递，不占内存，且切片的指针和原来数组的指针是不同的，不存在2种的弊端。
		所以，数组传参，请通过传切片的方式传递!
	***/
}

// 传数组
func printArrayInfo1(x [2]int) {
	fmt.Printf("funcArray : %p , %v, %d, %d\n", &x, x, len(x), cap(x))
}

// 传数组指针
func printArrayInfo2(x *[2]int) {
	fmt.Printf("funcPtr : %p , %v\n", x, *x)
	(*x)[1] += 100
}

// 传切片
func printArrayInfo3(x *[]int) {
	fmt.Printf("funcSlice : %p , %v\n", x, *x)
	(*x)[1] += 100
}
