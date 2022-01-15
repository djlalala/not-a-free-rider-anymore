// author 董健 <joeyana@aliyun.com>
package main

/***
	Map实现原理
***/
func main() {
	// TODO
	// 当往map中存储一个kv对时，通过k获取hash值，
	// hash值的低八位和bucket数组长度取余，定位到在数组中的那个下标，
	// hash值的高八位存储在bucket中的tophash中，用来快速判断key是否存在，
	// key和value的具体值则通过指针运算存储，当一个bucket满时，通过overfolw指针链接到下一个bucket。
}
