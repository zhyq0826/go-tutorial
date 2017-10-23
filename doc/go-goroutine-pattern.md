title: Go 中的并发模式
date: 2017-09-22 17:54:56
tags:
categories:
---


# 1. Timeout Out Connections

利用 select 读取多个 channel 的特性，可以实现简单的超时处理

```go
package main

import (
    "fmt"
    "time"
)

func timeout(t chan bool) {
    time.Sleep(5000000000)
    t <- true
}

func readString(s chan string) {
    // s <- "string"
}

//timeout 超时设置
func main() {
    t := make(chan bool)
    s := make(chan string)

    go readString(s)
    go timeout(t)

    select {
    case msg := <-s:
        fmt.Println("recevied", msg)
    case <-t:
        fmt.Println("time out")
    }
}

```


# Aliased xor Mutable 

Go 中通过共享对象来实现并发中消息的传递，你经常会不小心写出多个 goroutine 同时修改一个 object 的 代码，为了避免这样的问题出现，请遵循以下的规则
- 通过 channel 传递指针也应该传递指针的拥有者
- 避免通过参数向新开启的 goroutine 中传递调用者可能修改的对象的指针

总结来说就是尽量避免两个 goroutine 同时修改对象。


# Share Memory by Communicating



