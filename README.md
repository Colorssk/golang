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
    "strings"// strings.Split/strings.Contains/.HasPrefix(前缀判断)/.HasSuffix(后缀判断)/Index(第一次出现的位置)/LastIndex(最后出现的位置)/join(a[]string,sep string)
)

len表示的是字节数
要判断字符数:
rune类型用来表示utf8字符，一个rune字符由1个或多个byte组成

时间和日期类型

time包：time.Time类型, 用来表示时间
time.Now()获取当前时间
time.Now().Unix()当前时间戳

time.Duration 表示纳秒
常量: 
Nanosecond Duration = 1
Microsecond = 1000*Nanosecond
Millisecond = 1000*Microsecond
Second = 1000*~
...

格式化时间：time.Now().Format("2006/01/02 15:04:05")// 要把2006作为模版


if的其他用法：  if num:=10;num%2==0 {

}

for的不同的用法：
for no,i:=10,1;i<=10&&no<=19;i,no = i+1,no+1{

}
for{// 无限循环

}

switch a:=2,a {
    case 1:
    ...
    case 2:
    ...
    // 其他用法
    case 'a','b','c','d':
    ...
    default:
    ...
}
// 或者
a:=2
switch  {
    case a>1:// 此时每个都是一个条件判断 加上fallthrough之后会穿透到下一个代码执行
    ...
    case a<1:
    ...
    default:
    ...
}
// function前面的形参参数类型不写会继承后面的类型
fun add(a,b int)(sum int,sub int){
    sum := a+b
    sub := a-b
    return
}

func main() {
    sub,_:=add(2,3)//_表示不需要的返回值，申明不使用不会报错
}
// 传参：可变参数
func calc_v1(b ...int)(sub in,sum int){// 表示可以传一个参数b可以不传也可以传多个
    return
}


defer语句
defer用来定义语句，只有在函数返回之前才会执行




