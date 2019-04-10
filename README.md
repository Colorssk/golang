# golang
golang project

目录介绍：
1: bin 编译后的二进制文件  

2: pkg 编译后的包文件  

3： src 源码文件  

npm run build 自动编译src下面的项目（因为go是编译语言，非脚本语言  

自动根据go文件生成.exe文件(可执行文件)  

手动分界线---
上面的目录适用于个人


在github.com(编译时候自动生成的)新建colorsskproject（企业目录）
 colorsskproject作为总仓库
line1作为不同的业务组（比如前端业务组）(可以建立多个业务组)
line1下面就可以建立多个项目(比如hello)（每个项目下面又有多个模块）

企业目录的编译步骤：
go build github.com/colorsskproject/line1/hello
然后直接执行hello.exe即可
指定编译后的可执行二进制文件的目录：
go build -o bin/hello.exe github.com/colorsskproject/line1/hello

运行执行脚本：
go run  [运行的文件路径]（此时的路径是全部的：带src）


go install: 安装可执行文件到bin目录（go build之后的二进制文件再放入bin目录下面）
例子： go install github.com/colorsskproject/line1/hello
此时相当于执行了go build -o bin/hello.exe github.com/colorsskproject/line1/hello



go其他基本命令介绍：
go test 执行单元测试或压力测试
go env 显示go相关的环境变量
go fmt 格式化源代码: go fmt github.com/colorsskproject/line1/hello 类似于eslint规范自动校正

go的性能了解：

go的垃圾回收：有专门的cg线程，只需要new 分配内存，只要没有被引用了，就会被自动回收

go的特性：
天然并发： 从语言层面支持并发(c,c++,java是操作系统层面上的线程)
goruntine: 轻量级线程，成千上万的goroute成为可能

例子：
func calc(){
    // 大量的计算
}
func main(){
    go calc()//此时大量的计算会在一个新开辟的线程中运行而不会阻塞主线程
}


channel:
管道，类似unix/linux的pipe;多个goroute之间通过channel进行通信;支持类型

import(
    "strings"// strings.Split/strings.Contains/.HasPrefix(前缀判断)/.HasSuffix(后缀判断)/Index(第一次出现的位置)/LastIndex(最后出现的位置)
)








