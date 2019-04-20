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

var a[8] int数组  值类型   arr:=[...]int{1,2,3}// 让编译器自己确定长度
var a []int切片（可以用数组的所有操作）   引用类型 直接用==null/nil 判断   并且通过数组转化成切片之后，切片值改变，数组的值也变化
切片和数组都要初始化之后才能够赋值
a:= [5]int{1,2,34}
var b []int = a[1,2]// 切片左闭右开  (slice) [1:]/[:5]/[:](包含所有元素)
或者：
 定义：  c:= [] int{6,7,8}// 不限长度


 make方式创建切片
 //cap() 切片容量
 i:=make([]int,5,10)// 第一个5表示切片的长度，后面一个表示切片的容量, 底层数组长度是10超过5个以后不能直接i[6]操作但是可以往里面添加数据： append(i,11)， append之后的值是放在5个元素后面的
 append假如查过容量，底层的长度就会加倍扩容

 a:= [...]string{"1","2","3","4"}
 b:= a[1:3]
 那么b的容量就是3，从第二个开始到最后

 密码生成工具：文件夹:password

 


 


 切片拼接

 veggies:=[]string{"pa","as","aa"}
 fruites:=[]string{"aa","bb","cc"}
 元素后面的拓展符号表示参数
 food:=append(veggies,fruits...)

 copy(a1,a2)// 把第二个切面拷贝到第一个切片当中（会覆盖）但是不会扩容，也就是a1多大只会拿多大的数据


 make和new的区别
 make为内建类型slice,map,和channel分配内存
 
 new用于各种类型的内存分配,new返回是一个指针比如int和数组
 // 文件夹：makeandnew


 排序包
 “sort”

 sort.Ints(a[:])// 对切片升序排序
 sort.Strings(a[:])// 对字符串进行排序
 // 同理 .Float64s()

指针类型存储地址又叫引用类型
申明指针：
var b int32
b = 156
var a *int32
a = &b

操作指针变量:  获取值 *a

map结构(key-value),map默认初始化为nil，需要使用make分配map内存(初始化):make(map[string]int)
 var a map[string]int
 var b map[int]string

 见map文件夹(map是引用类型)

 注意：map的扩容是有消耗的(涉及到扩容和拷贝)

 判断map中的指定的key是否存在  value,ok := map[key]// 存在就会获得value并且ok为true

 遍历map用for key,value := range a{

 }
 删除map中的函数

 delete(a,"key1")

 长度len(a)


 ////
 定义map类型的切片（切片的类型是map）//有点像数组对象
 var mapslice []map[string]int
 mapslice = make([]map[string]int,5)// 定义二维“数组”的时候也要先初始化再使用
 mapslice[0] = make(map[string]int,10)//相当于要先把数组元素先声明类型再使用
 mapslice[0]["a"] = 10


 // 同理存在对象数组(map中的元素是切片)

var s map[string][]int = make(map[string][]int,5)
// 同理需要判断map元素是否存在不存在，需要初始化
key:="aa"
value,ok:= s[key]
if !ok{
    //不存在，一定要初始化(初始化切片)
    s[key] = make([]int,0,10)
}
value = append(value,100)
s[key] = value
