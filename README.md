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


// 

interface(数据类型)
var a interface{}
var b int32
// 因为a是interface类型，所以可以往里面存储任何数据类型
a = b// 成立

单引号就表示一个字符，字符串使用双引号

//自定义包package 生成.a可执行文件（静态文件）
进入当前包 然后执行go install 才会生成静态库
静态库生成目录在：go在项目的pkg文件夹下，具体包名下

引用方式:除了src之外的绝对路径（路径是包的路径，不是可执行文件的路径）

一个包中可以有0个或多个init函数，在程序启动时自动调用

执行顺序: 先执行全局变量初始化-> 然后再执行init函数-> 最后执行main函数

在包导入的顺序中，最后导入的包先执行（先完成全局变量初始化和init函数）

如果_标识符加载包的前面表示导入这个包但是没引用这个包中的属性和方法(包中的init方法会执行)

// switch char {
    case 'a','A':
      // doing
    case 'e','E'：
      //doing
}

结构体
通过struct实现面向对象，struct是用户自定义的类型
type User struct {
    Username string
    Sex string
    Age int
    AvataUrl string
    int
    string//  匿名字段
    *Address// 匿名结构体// 这样变相实现继承
}

初始化1:
     var user User
     user.Age = 18
     // 匿名函数初始化
     user.int = 100
     user.Address = &Address{
         a:'a',
         b:'b'
     }
     // 或者直接
     user.a = 'a'// 直接给Address中的属性赋值，同理访问的时候可以直接调用里面结构体中的属性
     ...
初始化2：
    var user User = User {
        "Username": "user01",
        "Age": 18,// 也可以部分初始化
    }

// 结构体类型的指针
 var user *User  = &User{}//&User{}和new(User)一样都是返回一个结构体的地址
 //或者
 var user *User  = &User{
     "Username":"user01"
 }
 // 或者
 user:= new(User)
 user.Username="user01"// 这里是个语法糖相当于(*user).Username="user01"


 指针时候访问： user相当于*user

 结构体内存布局:
 占用连续的内存空间

 go语言中struct没有构造函数和析构函数但是可以自己实现一套


 匿名冲突(外层和内层结构体有相同的匿名字段后者署名字段)
 此时初始化需要具体指向内部结构体:user.Address =  new(Address)
 访问时候同理: user.Address.a = "a"

 tag是结构体中的元信息,可以在运行的时候通过反射的机制读取出来

 type User struct{
     Username string `json:"username",db:"user_name"`
     Sex string `json:"sex"`// key-value格式
 }

json序列化
包：encoding/json
data,_:=json.Marshal(User)//data是字符数组 string(data)转化为字符串
此时输出就是小写的username,用途json序列化的时候传参需要大写

var a int
输入：fmt.Scanf("%d\n",&a)

go中的方法定义：是在函数前面加上一个接受者，这样就知道这个方法属于哪个类型了

type A struct {

}

func (a A) Test(s string){//A是Test的接受者，因此A这个对象就有Test方法
}

//也可以为当前包内定义的任何类型增加方法
type Integer int

func (i Integer)Print(){
    
}
func main(){
    var a Integer
    a = 1000
    a.Print()

    //或者
    var b int =200
    a =Integer(b)// 就是Integer结构对象，相当于初始化了
    a.pritn()
}

// 为了能够让结构体是引用类型：
 func (p *People) Set(name string,country string){
    p.Name = name
    p.Country = country
 }

 //调用的时候  初始化之后
 p1 := People{
		Name:   "aa",
		Contry: "China",
	}
   //(&p1).Set("a","b")//此时修改之后的值就会改变
   当然上面是底层会帮做的所以正常就可
   p1.Set("a","b")


// 结构体和json序列化 json(文件夹)

格式化输入
var a int
fmt.Scanf("%d%v%d",&a,&b,&c)// 一定要用地址   输入的时候要加空格
单个输入需要  fmt.Scanf("%d\n",&a)
fmt.Scan(&a,&b,&c)// 此时支持换行输入
fmt.Scanln(&a,&b,&c)// 此时一定要在同一行输入

// 从字符串中获取输入
str:="88 asd 10.0"
fmt.Sscanf(str,"%d%s%f\n",&a,&b,&c)
str1:="88 asd\n\n 10.0"
fmt.Sscan(str,&a,&b,&c)// 此时支持换行输入
fmt.Sscanln(str,&a,&b,&c)// 此时一定要在同一行输入

   
// 从文件中输入
fmt.Fscanf(os.Stdin,"%d%s%f\n",&a,&b,&c)
fmt.Fscan(os.Stdin,&a,&b,&c)
...同理
fmt.Fprintln(os.Stdout,a,b,c)

//buffer流
import bufio
reader:=bufio.NewReader(os.Stdio)
var str string
str,_=reader.ReadString('\n')//标识符是读取到换行符

//控制台传参(字符串的切片)
main(){
    if len(os.Args)>1{
        fmt.Println(os.Args[0])
        for index,v := range os.Args{
            if index =  0{// 第零个参数不考虑
                continue
            }
        }
    }
}


cli包的使用
import github.com/urfave/cli
(在项目目录下运行 go get github.com/urfave/cli)

框架：控制台程序的框架
func main(){
app:=cli.NewApp()
app.Name='app'

app.Usage = "fight"
app.Action = func(c *cli.Context) error{
    if c.NArgs>0//表示传过来的参数的个数 
    fmt.Print("hello world")
    // 业务代码  c.Args().Get(0)//获取第一个参数或者c.Args()[0]
    return nil
}
app.Run(os.Args)

}


// flag
  var lang string
  app.Flags = []cli.Flag{
      cli.StringFlag{
          Name: 'lang,l',后面是短选项  xxx.exe --lang arg   传入参数arg
          Value:'aa',
          Usage:'aa',
          Destination:&lang,
      },
      cli.BoolFlag{
          Name: ''
          Usage:'',
          Destination:,
      }
  }

文件读写:
//只读
 file,err = os.Open("")


defer file.Close()
var buf[128]byte
n,err:=file.Read(buf[:])
//使用
string(buf[0:n])


或者把读取的结果放入切片
var content []byte
for{
    n,err:=file.Read(buf[:])
    if err == io.EOF{//文件读取结束
    break
    }
    if err!=nil{//
        return
    }
    content  = append(content,buf[0:n]...)//把一个切片添加到另一个切片需要展开
}

//使用缓存buffer
 file,err = os.Open("")


defer file.Close()
reader:=bufio.NewReader(file)
for{
    line,err:=reader.ReadString("\n")
     if err == io.EOF{//文件读取结束
    break
    }
    if err!=nil{//
        return
    }
}

// 文件读取
import io/ioutil
content,err:=ioutil.ReadFile("")//cotent是字节切片

fmt.Print(string(content))

//读取压缩文件
compress/gzip

fi,err:=os.Open("...gz")
fz,err:=gzip.NewReader(fi)


var buf [128]byte
n,err:=fz.Read(buf)

//文件写入
file,err:=os.OpenFile("./test.data",os.O_CREATE|os.O_CREATE|os.O_WRONLY,0666)//创建|如果有覆盖...

defer file.Close()
file.Write([]byte("hello world"))

writeat
writestring


//bufio方式写入
op,err:=bufio.NewWrite(file)
op.WriteString("....")
op.Flush()//刷新从缓冲区写入磁盘

// ioutil写入
content,err:= ioutil.WriteFile("test.dat",[]byte("aaa"),0755)//写入的是字节数组


//拷贝文件
 b,err:=os.Open("")
 a,err:=is.OpenFile("",os.O_WRONLY|os.O_CREATE,0644)
 io.Copy(a,b)//b文件写入a


//defer
x:=5
defer func(){
 x=1
}
return x
x赋值
defer指令
ret指令

但是 
func func()(y int){
    x:=5
    defer func(){
        x+=1//改的局部变量,y没有改变
    }()
    return x
}
先x赋值y
derfer执行，值类型，x变化，y没变
return  y 值是5

同理: defer func(x int){

}(x)//x作为新参传入的是副本


接口：
 type Animaml interface{
     Talk()
     Eat()
     Name()String
 }
 type Dog struct {

 }
 func (d Dog) Talk(){

 }
 func(d Dog)Eat(){
    fmt.Println("aa")
 }
 func(d Dog)Name(){

 }
 func(d Dog)Name()string{
     
 func main(){
     var d Dog
     var a Animal
     a = d//d这个接口对象需要包含接口中的所有方法

     a.Eat()// aa
 }


 //空接口
 var a interface{}
 var b int = 100
 a = b

 //类型断言
 func test(a interface{}){  a Animal//参数可以传具体的接口，下面再判断
     s,ok  := a.(int)// 此时a就断言只能是int类型
     if !ok{
         fmt.Println("")     }
     // type switch
     switch a.(type){//这里一定是type
        case string:
        ...
        case int:
        default:
        ...
     }
     
 }

 // 指针类型的接口
 var a animal// 接口也要申明
 var d *Dog = &Dog{}
 a = d
 a.Eat()//==*(&Dog).Eat()

 但是如果是指针类型实现的这个接口上面就会报差
 func (d *Dog)Name() string{

 }

 //接口嵌套
 type a interface{
 }
 type b interface{
 }
 type c interface{
     a
     b
 }


 //获取路径最后的文件名:
 path.Base("...")

 //获取cpu执行到的位置
 import "runtime"


 pc.file,line,ok:=runtime.Caller(2)//参数2表示栈的深度

if ok{
    file = file//执行到的文件
    funcName = runtime.FuncForPc(pc).Name()//执行到的函数
    lineNo = line// 行号
}

反射：
内置包： reflect
获取类型信息： reflect.TypeOf
获取值信息： reflect.ValueOf// 这里拿到的仅仅是一个副本所以直接用v.SetInt()设值会报错

t := reflect.TypeOf(a)// 返回的是结构体，所以之后用Kind()方法具体判断
k:=t.Kind()
switch k {
    case reflect.Int64:
    ...
    case reflect.String:
    ...
    case reflect.Chan:
    ...
    case reflect.Map:
    ...
    case reflect.Struct:
    ...
    case reflect.Ptr:
}

同理：
v:=reflect.VlaueOf(x)
v.Type(),
v.Kind()
v.Float()// 获取在值信息

通过反射设置值:
v.SetInt(100)
v.SetFloat(6.9)

注意点：分装成方法的时候传值需要时地址: &x
而指针中的赋值是这样的：
var b *int = new(int)
*b = 100
所以：
case reflect.Ptr:
    v.Elem().SetFloat(6.8)

对于结构体来说：

type Student struct {
    Name string
    ......
}

func main(){
    var s Student
    v:=reflect.VlaueOf(s)//     v.NumField()   field:=v.Field(i)    field.Interface()//转化成接口类型  结构体元素值
    t:=v.Type()  t.Field(i).Name// 结构体变量名字
    kind:=t.kind()
    switch kind {
    case reflect.Int64:
    ...
    case reflect.String:
    ...
    case reflect.Struct:
}


结构体赋值：
var s Student
v:=reflect.ValueOf(&s)

v.Elem().Field(0).SetString("123")
v.Elem().FiledByName("Sex").SetInt(2)


//访问本地服务器：

命令： godoc -http=:9090


配置库项目：解析ini配置文件
[server]
ip=10.230.2.2
port = 8080

[mysql]
username = root
passwd = root
database = test
host=192.168.1.1
port = 3838



单元测试的试用也在iniconfig文件中
field.Tag.Get("ini")//获取tag别名

字符串索引:
index := strings.Index(line, "=")

.SetString()
.Interface()//转化为%v任意类型

写入文件：
"io/ioutil"
ioutil.WriteFile(filename, result, 0755)


// 单元测试和压力测试:
testing包

测试包：在一个包下面test测试文件中可以是随意调用其他文件中的函数
单元测试：必须以:_test.go结尾，单元测试函数必须是Test开头可以是TestAdd(Testh后面的英文首字母大写)且只有一个参数: *Testing.T；
t.Fatalf();
t.Logf()
基准测试或者压力测试:必须是 Benchmark开头，参数是*Testing.B

压力测试案例：line19; 调用  go test -bench  
go test -bench   BenchmarkAdd //表示仅仅执行这个函数
go test -bench  . //执行所有压力测试用例
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 10
		b := 20
		Add(a, b)// 反复测试调用这个函数
	}
}

测试某个测试用例： go test -run ...

调试工具delve:
安装：
go get github.com/derekparker/delve/cmd/dlv
默认安装到GOPATH的bin目录下
设置环境变量

delve命令：
dlv debug  包或者源代码的路径
然后设置断点：
b main.main // b断点开头 
c // 执行程序
main.main  文件名字
p a // 显示a的值
next继续执行

print b //打印b的值

r// 重新执行程序
s代表单步调试

quit//退出

监听正在运行的项目：
dlv attach <pid>
然后设置断点:
b main.go:16// 在16行设置断点


多线程设置断点： 
dlv debug .\main.go
b productSudu// 进入某个线程调试
goroutines//查看现在有多少线程
goroutine <线程编号>
bt//查看具体线程堆栈

并发和并行；
并发： 两队人用一个咖啡机，一个段时间多个队列同时执行，但是实际上只有一个任务执行
并行： 两队人用两个咖啡机，某个时刻是多个任务同时进行的

用户态的线程叫做协程(goroutine)


goroutine创建： linethread->goroutine


runtime(文件linethread->runtime)
通过runtime包进行多核设置
GOMAXPROCS设置当前程序运行时占用的CPU核数
NumCPU获取当前系统的核数


channel是个队列是个容器（队列）

linethread->channel

var 变量名 chan 数据类型


用channel实现goruntine的同步
linethread->goroutine_sync

单向channel:申明 var c  <-  chan int//代表管道c只能读取数据
var c chan   <-  int // 代表只能插入数据


for range管道：linethread->for_range_chan


重要： 队列会发生死锁的两种情况：
1： 队列是空的取出来的时候回发生死锁
2： 队列设置了大小，插入的元素个数超过限制，发生死锁


channel的长度和容量：  cap(ch)// 容量  len(ch)// 长度


// 重要：
等待一组goroutine的结束：
1： 用带缓存区的channel:  linethread->waitgroup1(exitchannel沟通所有线程)
更好的方式：linethread->waitgroup
2：用sync包中的WaitGroup

workerpool:linethread->workerpool（正经的生产消费者模式）


select语法:（实现同时监听多个管道）
linethread->select// 实现多个channel同时读写
也有用select写入管道linethread->select


线程安全：
多个goroutine操作同一个资源（临界区），需要互斥锁
linethread->mutex


读写锁（读多写少）linethread->rw_lock
场景： 当一个线程写的时候其他线程都等待
       当一个线程读的时候，除了写的线程等待，读的线程均可运行

原子操作：（如果有大量的goroutine加锁会产生平凡的上下文切换）       
加锁代价比较耗时，所以使用原子操作：
针对基本数据类型
原子操作的方法有：
1:加减操作：
AddInt32(addr *int32, data int64)
AddInt64;
AddUnit32;
AddUnit64
2:比较病交换
CompareAndSwapInt32(add *int32,old,new int32);
CompareAndSwapInt64;
CompareAndSwapPointer
读取操作：
LoadInt32(addr *int32);
LoadInt64;
LoadPointer
3：写入操作:
StroreInt32(addr *int32);
StroreInt64;
StrorePointer
4:交换操作：
SwapInt32(addr *int32,val int32);
SwapInt64;
SwapPointer

服务端的tcp连接  linethread->tcp_server
客户端的tcp连接  linethread->tcp_client

http客户端连接   linethread->http_client(发送http请求)

udp实时性比较好  test

设置http服务器   linethread->http_server  设置路由和监听端口等

表单处理：linethread->form_proc
html模板：web->template

{{}}// 替换元素；
{{.}}// 标识当前对象
{{.Filenam}}// 访问当前对象属性

模板中的if：
{{if gt .Age 18}}
{{else}}

判断:
if not .condition
end

if and .condition1 .condition2
end

if or .condition1 .condition2
end

if eq .condition1 .condition2
end
if ne .condition1 .condition2(不等于)
end

if lt .condition1 .condition2
end

if le .condition1 .condition2(小于等于)
end

if ge .condition1 .condition2(大于等于)

模板for web-> template_for
range .

web服务器平滑升级(优雅重启)
文件继承(只支持linux不支持windows),子进程继承父进程:
startChild(){
    args := []string{"-child"}
    cmd := exec.Command
}


jason格式：
web->json_example


xml和结构体之间的转换 web->xml_example
encoding/xml


msgpack数据交换格式：二进制的json协议  web->msgpack_example
需要安装包： go get github.com/vmihailenco/msgpack

protobuf 数据格式：google推出，二进制（属于中间语言）
需要：idl编写，生成指定语言的代码，序列化和反序列化
idl中定义
枚举类型：
enum EnumAllowingAlias{
    UNKNOW = 0;
    STARTED = 1;
    RUNNUING = 2;
}

结构体类型：
message Person {
    int32 id = 1;
    string name = 2;// 后面的数字表示标识号
    repeated Phone phones = 3;//repeated表示可重复的(相当于数组/切片)
}

使用：
安装protoc编译器，解压后拷贝到GOPATH/bin目录下
下载路径：https://github.com/google/protobuf/releases
安装golang代码插件: go get -u github.com/golang/protobuf/protoc-gen-go

写完dle后在当前目录生成go的代码: protoc --go_out=./address/ .\person.proto    //前面是生成存放的路径，后面是目标文件 


依赖管理工具： godep:
下载: go get github.com/tools/godep
输入godep有帮助信息说明安装成功
godep save
生成godep(版本信息)文件和vendor（第三方包源码）文件


golang中提供mysql驱动接口和管理：
github.com/go-sql-driver/mysql

创建表例子:
CREATE TABLE 'user' {
'id' bingint(20) NOT_NULL AUTO_INCREMENT,
'name' varchar(20) DEFAULT '',
PRIMARY KEY ('id')
}ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

提供方法：web->mysql_example
单行查询：Db.QueryRow
多行查询：Db.Query


sql预处理：(命令部分和数据部分分开来单独发)：防范sql注入
DB.Prepare(sql string)(*sql.Stmt,error)
Stmt.Query()
DB.Prepare(sql string)(*sql.Stmt,error)
Stmt.Exec()

mysql事务处理步骤：
Db.Begin()// 开启事务
Db.Commit()// 提交一个事务
Db.Rollback()// 回滚一个事务


性能更好的第三方库： sqlx库: web->sqlx_example
支持: mysql，postgresql,oracle,sqlite；
查询  sqlx.DB.Get和sql.DB.Select
更新，插入和删除： sql.DB.Exec
事务：sql.DB.Begin(),sql.DB.Commit(),sql.DB.Rollback()

自己拼装sql会有sql注入的漏洞
......
