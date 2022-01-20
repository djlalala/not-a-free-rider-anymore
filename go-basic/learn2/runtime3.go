package main

import (
	"fmt"
	"runtime"
	"time"
)

/***
	通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
	Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。
	默认值是机器上的CPU核心数。
	例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

	Go语言中的操作系统线程和goroutine的关系：
	1.一个操作系统线程对应用户态多个goroutine。
	2.go程序可以同时使用多个操作系统线程。
	3.goroutine和OS线程是多对多的关系，即m:n。

	注:
	macOS:
	查看CPU信息：sysctl machdep.cpu
	machdep.cpu.cores_per_package: 8
	machdep.cpu.core_count: 8  // 核数为8
	machdep.cpu.logical_per_package: 8
	machdep.cpu.thread_count: 8 // cpu数量为8个 , 若16则使用了超线程技术：8核16线程
	machdep.cpu.brand_string: Apple M1


	Linux:
	# 查看物理CPU个数
	cat /proc/cpuinfo| grep "physical id"| sort| uniq| wc -l
	# 查看每个物理CPU中core的个数(即核数)
	cat /proc/cpuinfo| grep "cpu cores"| uniq
	# 查看逻辑CPU的个数
	cat /proc/cpuinfo| grep "processor"| wc -l
***/
func main() {
	// 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	// 将逻辑核心数设为2，此时两个任务并行执行。
	runtime.GOMAXPROCS(1) // 通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
	go a()
	go b()
	time.Sleep(time.Second)
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
