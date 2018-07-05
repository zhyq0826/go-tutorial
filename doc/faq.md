## FAQ

### 为什么不能在主 goroutine 中直接使用 channel

```go
package main

import "fmt"

func main() {
    c := make(chan int)    
    c <- 1   
    fmt.Println(<-c)
}
```

c 是一个 unbufferd channel ，会一直会阻塞着不发送数据，除非在一个 goroutine 中已经有了它的接收者，在上面的代码中，由于 c<-1 是在接收之前执行的，c 没有看到它的接收者，而它本身又需要一直阻塞着，但是它又必须执行找到它的接收者，于是就发生了死锁。

When you create a channel in Go with ch := make(chan bool), an unbuffered channel for bools is created. What does this mean for your program? For one, if you read (value := <-ch) it will block until there is data to receive. Secondly anything sending (ch <- true) will block until there is somebody to read it. Unbuffered channels make a perfect tool for synchronizing multiple goroutines.


### why 两个 named type 不能相互赋值

https://groups.google.com/forum/#!topic/golang-nuts/4Db2z2dEhfc

That's correct.
The reasoning goes like this: If you take the trouble to name Pint and
Xint, it's because you want them to be distinct.  (This isn't C, where
a typedef is just an alias.)  But sometimes it makes sense to speak of
the structure as all you care about; consider things like the indexing
functions in the bytes package.  So we allow assignment in those
cases.

The motivating example in our thinking was something like type Point
struct { X, Y float }.   Another named type with the same fields
probably is a different idea or it would share the declaration; they
shouldn't be interchangeable.  Another way to say it is that if they
don't have the same methods (even potentially), they shouldn't be
assignable.  On the other hand a generic struct { X, Y int } is
talking just about the structure, not its interpretation, so it makes
sense to allow assignment between that type and Point.

Also, although I cannot reconstruct the history, the current rules
were arrived at largely through experience coupled with a desire for a
simple specification.  They aren't arbitrary.

underlying type 一样的 named type 不能赋值，既然是 named type，即使 underlying type 一样，肯定是为了作区分才分别写成两个 named type

-rob

###  go 安装依赖包一般会通过四种路径

```
1. github.com/
2. golang.org/
3. gopkg.in/
4. honnet.co/
```

比如我们可以通过 `go get github.com/xxx` 来下载安装包
下载好之后, 通过 `go install github.com/xxx` 来安装包

```
安装包会下载到 `$GOPATH/src` 文件中
安装后的执行文件在 `$GOPATH/bin` 文件
```

**常见错误**

当我们执行 `go get golang.org/x/tools/cmd/goimports` 会报错
*package golang.org/x/tools/cmd/goimports: unrecognized import path "golang.org/x/tools/cmd/goimports"*

这个问题会出现在高版本的 golang, 一般的解决办法是

```
# 创建文件夹  
mkdir $GOPATH/src/golang.org/x/ 
# 进入文件夹 
cd  $GOPATH/src/golang.org/x/
# 下载源码 
git clone https://github.com/golang/tools.git
# 安装
go install golang.org/x/tools/cmd/goimports
```

同理可解决其他处在 `golang.org/x/` 路径下的包

