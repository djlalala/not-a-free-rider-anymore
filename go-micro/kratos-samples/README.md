# Kratos 框架学习

## 环境准备
* protobuf安装
```
# refer https://www.jianshu.com/p/67f64307d268
1. github 下载protobuf-3.12.4包
2. sudo -i 切换root用户
3. cd /Users/xxx/xxx/protobuf/protobuf-3.12.4
4. ./configure --prefix=/usr/local/protobuf 设置编译目录
5. make 编译
6. make install 
7. .bash_profile 配置环境变量 
export PROTOBUF=/usr/local/protobuf 
export PATH=$PROTOBUF/bin:$PATH 
```