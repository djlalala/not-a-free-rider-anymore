package main

import (
	"fmt"
	"math"
)

/***
	1.基本类型介绍
	2.
***/
func main() {
	// 1. 基本类型介绍
	// 类型	长度(字节)	默认值	说明
	// bool	1	false
	// byte	1	0	uint8
	// rune	4	0	Unicode Code Point, int32
	// int, uint	4或8	0	32 或 64 位
	// int8, uint8	1	0	-128 ~ 127, 0 ~ 255，byte是uint8 的别名
	// int16, uint16	2	0	-32768 ~ 32767, 0 ~ 65535
	// int32, uint32	4	0	-21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名
	// int64, uint64	8	0
	// float32	4	0.0
	// float64	8	0.0
	// complex64	8
	// complex128	16
	// uintptr	4或8		以存储指针的 uint32 或 uint64 整数
	// array			值类型
	// struct			值类型
	// string		""	UTF-8 字符串
	// slice		nil	引用类型
	// map		nil	引用类型
	// channel		nil	引用类型
	// interface		nil	接口
	// function		nil	函数

	// 2. 字符串转义符
	// 		\r	回车符（返回行首）
	// 		\n	换行符（直接跳到下一行的同列位置）
	// 		\t	制表符
	// 		\'	单引号
	// 		\"	双引号
	// 		\	反斜杠

	// 3.类型转换
	var x, y = 3, 4
	var z int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	z = int(math.Sqrt(float64(x*x + y*y)))
	fmt.Println(z)

	// 4.byte和rune类型
	// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	// rune类型，代表一个 UTF-8字符
	var a = '中'
	var b = 'x'
	fmt.Println("字符用单引号包裹:", a, b)
	fmt.Println("字符用单引号包裹:", string(a), string(b))

}
