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

待续。。。



