package main

import (
	"fmt"
	"go-basic/tools"
)

/**
	1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
	2. 切片的长度可以改变，因此，切片是一个可变的数组。
	3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
	4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
	5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
	6. 如果 slice == nil，那么 len、cap 结果都等于 0。
**/
var intArr []int = []int{222, 333, 444, 555, 666, 777}
var slice []int = intArr[4:]

// var slice []int = intArr[:2]
// var slice []int = intArr[1:2]

func main() {
	fmt.Printf("intArr是:%v\n", intArr)

	fmt.Printf("slice是:%v\n", slice)

	/***
		创建切片
		var slice []type = make([]type, len)
		slice  := make([]type, len)
		slice  := make([]type, len, cap)
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("创建切片:")
	var slice0 []int = make([]int, 20)
	fmt.Printf("当前创建的切片是:%v\n", slice0)
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。其中的8就是索引的index
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值，len不能大于cap
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))

	fmt.Println("------------------------------------------------------------")
	fmt.Println("使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。\n还可用指针直接访问底层数组，退化成普通数组操作。")
	s := []int{0, 1, 2, 3}
	fmt.Println("初始s数组是:", s)
	p := &s[2] // *int, 获取底层数组元素指针。
	*p += 100
	fmt.Println("当前s数组是:", s)

	/***
		[][]T 指的是 元素类型为 []T
		比如这里 [][]int 类型为 []int
	***/
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println("[][]int的data是:", data)

	/***
		可直接修改 struct array/slice 成员。
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("可直接修改 struct array/slice 成员:")
	d := [5]struct {
		x int
		// y string
	}{}
	fmt.Println("原始的d是:", d)
	// s := d[:]
	d[1].x = 10
	// s[2].x = 20
	fmt.Println("现在的d是:", d)
	fmt.Printf("%p, %p\n", &d, &d[0])

	/***
		用append内置函数操作切片（切片追加）
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("用append内置函数操作切片（切片追加）:")
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	// d := append(c, 7)
	// fmt.Printf("slice d : %v\n", d)
	e := append(c, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)

	/***
		append ：向 slice 尾部添加数据，返回新的 slice 对象。
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("append ：向 slice 尾部添加数据，返回新的 slice 对象:")
	ss1 := make([]int, 0, 5)
	fmt.Printf("原来的ss1的内存地址:%p\n", &ss1)
	fmt.Printf("原来的ss1的值:%v\n", ss1)
	ss2 := append(ss1, 1)
	fmt.Printf("现在的ss2的内存地址:%p\n", &ss2)
	fmt.Printf("现在的ss2的值:%v\n", ss2)
	fmt.Println(ss1, ss2)

	/***
		超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满:")
	ssdata := [...]int{0, 1, 888, 3, 4, 5, 10: 0}
	sss := ssdata[2:4] // 因为从2开始，所以sss的cap是11-2=9
	fmt.Println("原始的ssdata数组是:", ssdata, len(ssdata), cap(ssdata))
	fmt.Println("原始的sss的slice是:", sss, len(sss), cap(sss))
	// sss[0] = 8888 // 此时我们通过slice去改变数组中的值，引用改值

	// sss = append(sss, 100, 200, 300) // 一次 append 两个值，超出 sss.cap 限制。
	// sss = append(sss, 100, 200, 300, 1, 2, 3, 4, 5,6) // 一次 append 两个值，超出 sss.cap 限制。
	sss = append(sss, 100, 200, 300, 1, 2, 3, 4, 5) // 一次 append 两个值，超出 sss.cap 限制。

	/***
		结论1.当前append的长度小于slice指针继续可以在数组上往后移动的长度的时候，slice追加会覆盖原来arr
		结论2.当前append的长度大于slice指针继续可以在数组上往后移动的长度的时候，slice会重新分配底层数组，与原数组无关
		通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。
		或初始化足够长的 len 属性，改用索引号进行操作。及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。
	***/
	fmt.Println("append之后的ssdata是:", ssdata, len(ssdata), cap(ssdata)) // 重新分配底层数组，与原数组无关。
	fmt.Println("append之后的sss是:", sss, len(sss), cap(sss))             // 重新分配底层数组，与原数组无关。
	fmt.Println("结论1.当slice的append的长度小于slice指针继续可以在数组上往后移动的长度的时候，slice追加会覆盖原来arr")
	fmt.Println("结论2.当slice的append的长度大于slice指针继续可以在数组上往后移动的长度的时候，slice会重新分配底层数组，与原数组无关")
	// 这里比对底层数组起始指针。
	fmt.Printf("slice中的值为:%v的指针,对应的内存地址:%p\n", sss[0], &sss[0])
	fmt.Printf("array中的值为:%v的指针,对应的内存地址:%p\n", ssdata[2], &ssdata[2])
	fmt.Println("由此可见:当前slice追加的长度大于指针可往右移动的长度时，slice会重新分配底层数组")

	/***
		slice中cap重新分配规律
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("slice中cap重新分配规律:")
	sa := make([]int, 0, 1)
	ca := cap(sa)
	for i := 0; i < 50; i++ {
		sa = append(sa, i)
		if n := cap(sa); n > ca {
			fmt.Printf("cap: %d -> %d\n", ca, n)
			ca = n
		}
	}

	/***
		切片拷贝
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("切片拷贝:")
	sss1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice sss1 : %v\n", sss1)
	sss2 := make([]int, 10)
	fmt.Printf("slice sss2 : %v\n", sss2)
	copy(s2, s1)
	fmt.Printf("执行copy(s2, s1)后的sss1 : %v\n", sss1)
	fmt.Printf("执行copy(s2, s1)后的sss2 : %v\n", sss2)

	//1.不同类型的切片无法复制
	//2.如果s1的长度大于s2的长度，将s2中对应位置上的值替换s1中对应位置的值
	//3.如果s1的长度小于s2的长度，多余的将不做替换
	// s1 := []int{1, 2, 3}
	// s2 := []int{4, 5}
	// s3 := []int{6, 7, 8, 9}
	// copy(s1, s2)
	// fmt.Println(s1) //[4 5 3]
	// copy(s2, s3)
	// fmt.Println(s2) //[6 7]

	sss3 := []int{1, 2, 3}
	fmt.Println("原始-slice sss3 : ", sss3, len(sss3), cap(sss3))
	sss3 = append(sss3, 4, 5, 6)
	fmt.Println("现在-slice sss3 : ", sss3, len(sss3), cap(sss3))
	sss3 = append(sss3, sss2...)
	fmt.Println("继续1-slice sss3 : ", sss3, len(sss3), cap(sss3))
	sss3 = append(sss3, 4, 5, 6)
	fmt.Println("继续2-slice sss3 : ", sss3, len(sss3), cap(sss3))

	/***
		字符串和切片
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("字符串和切片:")
	str := "hello-world"

	sliceX := str[0:5]
	fmt.Println("sliceX:", sliceX)
	sliceY := str[6:]
	fmt.Println("sliceY:", sliceY)

	str2 := "Hello world"
	bytes := []byte(str2)
	bytes[0] = 'D'
	fmt.Println(bytes)
	fmt.Println(bytes[2])
	fmt.Println(bytes[3])
	str2 = string(bytes)
	fmt.Println("newStr:", str2)

	// 中文
	str3 := "你好，世界！hello world！" //中文字符需要用[]rune(str)
	runes := []rune(str3)
	runes[3] = '够'
	runes[4] = '浪'
	runes[12] = 'g'
	runes = runes[:14]
	str = string(runes)
	fmt.Println(str)

	// test
	fmt.Println("------------------------------------------------------------")
	fmt.Println("slice全量等于arr的一个测试:")
	testArr := []int{1, 2, 3}
	testSlice := testArr[:]
	fmt.Println("testArr是:", testArr, len(testArr), cap(testArr), &testArr, &testArr[0])
	fmt.Println("testSlice是:", testSlice, len(testSlice), cap(testSlice), &testSlice, &testSlice[0])
	fmt.Println("结论:整个数据/长度/容量/内存地址/首个指针 都一致")

	/***
		golang slice data[:6:8] 两个冒号的理解
		常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为4（6-9）
		另一种写法： data[:6:8] 相当于data[0:6:8]每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8
		a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("slice两个冒号的理解:")
	sliceOO := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := sliceOO[7:8]
	fmt.Println("d1:", d1, len(d1), cap(d1))
	d2 := sliceOO[:7:8]
	fmt.Println("d2:", d2, len(d2), cap(d2))
	d3 := sliceOO[0:7:8]
	fmt.Println("d3:", d3, len(d3), cap(d3))
	d4 := sliceOO[2:7:8]
	fmt.Println("d4:", d4, len(d4), cap(d4))

	fmt.Println("------------------------------------------------------------")
	fmt.Println("数组or切片转字符串:")
	aIntArr := []int{1, 2, 3}
	aIntSlice := aIntArr[:]
	fmt.Println("初始数组和切片是:", aIntArr, aIntSlice)
	aIntArrString := tools.IntArrayToString(aIntArr)
	aIntSliceString := tools.IntSliceToString(aIntSlice)
	fmt.Println("数组转的字符串:" + aIntArrString)
	fmt.Println("切片转的字符串:" + aIntSliceString)

}
