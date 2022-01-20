# 学习问题记录
* slice 的append是具体操作的 (已结束)
    a := []int{1,2,3}
    s := arr[:]
    s = append(a,4,5,6)
    // 继续append
    这个时候的s,len,cap,指针位置是怎么变化的?

* Slice 创建切片/切片扩容/切片拷贝 的原理

* Map实现原理

* var s []int = []int{1,2,3} 和 s:=[]int{1,2,3}的区别是什么
思考记录,需要实际了解更多:
> s:=这种方式,后面是什么类型都可，并且默认s能够接收到类型
> var s Tpye =这种方式,需要type与后者对的上

* select流程控制