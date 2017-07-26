## blog

- https://www.goinggo.net/post.html

## packages

- [Understanding Golang Packages](https://thenewstack.io/understanding-golang-packages/)
- https://www.golang-book.com/books/intro/11
- http://www.amghost.cc/go-import-understand/

## interface

- https://research.swtch.com/interfaces
- http://jordanorelli.tumblr.com/post/32665860244/how-to-use-interfaces-in-go
- https://gobyexample.com/interfaces
- https://golang.org/doc/effective_go.html#interfaces
- https://my.oschina.net/xinxingegeya/blog/711916
- https://my.oschina.net/xlplbo/blog/199630

## type

- http://www.integralist.co.uk/posts/golang-webserver.html
- https://thenewstack.io/understanding-golang-type-system/

## pointer

- http://spf13.com/post/go-pointers-vs-references/

## map 

- https://blog.golang.org/go-maps-in-action

## nil

- https://golang.org/ref/spec#The_zero_value
- https://newfivefour.com/golang-null-nil.html
- https://www.gmarik.info/blog/2016/understanding-golang-nil-value/
- https://www.reddit.com/r/golang/comments/4injtt/understanding_gos_nil_value/
- http://www.tapirgames.com/blog/golang-nil
- https://katcipis.github.io/2016/09/17/fun-with-nil-interfaces.html

> If you want an interface value of nil, there's only two ways to get it, either assign it to 'nil' explicitly, or use the zero value of a variable of that type, i.e. var myErr error.

## gotchas-and-common-mistakes-in-go-golang

http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/

## go project structer


- https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project

## web

- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md


## web socket

- https://dinosaurscode.xyz/go/2016/07/17/go-websockets-tutorial/
- http://www.alexedwards.net/blog/golang-response-snippets#json

## go standlibary

- https://github.com/polaris1119/The-Golang-Standard-Library-by-Example


## go web examples

- https://gowebexamples.github.io/

## go examples

- https://gobyexample.com/


## chrome pdf
chrome.exe --save-as-pdf=<url> --output-file=<path-to-file> --pdf-format=A3 --pdf-orienation=Landscape --pdf-margins=10,10,10,10


## go defer

- https://blog.golang.org/defer-panic-and-recover
- https://www.goinggo.net/2013/06/understanding-defer-panic-and-recover.html

## channel

- http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
- http://colobu.com/2016/04/14/Golang-Channels/
- https://github.com/a8m/go-lang-cheat-sheet#concurrency
- https://stackoverflow.com/questions/18660533/why-using-unbuffered-channel-in-the-the-same-goroutine-gives-a-deadlock
- https://golang.org/doc/effective_go.html#concurrency
- http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/
- https://abronan.com/introduction-to-goroutines-and-go-channels/


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




