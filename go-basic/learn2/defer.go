package main

import (
	"go-basic/tools"
)

func main() {
	defer tools.Ending("1") // defer 是栈 ,先进后出
	defer tools.Ending("2") // defer 是栈 ,先进后出
}
