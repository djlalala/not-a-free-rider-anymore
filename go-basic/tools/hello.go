// author 董健 <joeyana@aliyun.com>
package tools

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func init() {
	fmt.Printf("这是tools包的 init 方法\n")
}

func Ending(str string) {
	fmt.Println("程序执行结束:", str)
}

func Hello() {
	fmt.Printf("OUTPUT:别做伸手党! 这是 %v() 方法\n", GetFunctionName(Hello))
}

// 数组转string
func IntArrayToString(arr []int) string {
	return strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", ",", -1)
}

// Slice转string
func IntSliceToString(slice []int) string {
	return strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1)
}

func GetFunctionName(i interface{}, seps ...rune) string {
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}
