## mgo 安装 

gopkg.in 被墙了，去 https://gopm.io/download 输入 import 的包路径，然后下载，把源代码放到 src  目录下，类似 /Users/zhaoyongqiang/go-workspace/src/gopkg.in/mgo.v2， 然后进行 build and install


## gosublime 设置 path

```
{

    "env": {
        "GOPATH": "/Users/zhaoyongqiang/go-workspace"
    }
}
```

在 gosublime 中设置 gopath 帮助 gosublime code completion


## mgo 的使用

mgo API 很像 pymongo，使用 bson.M 完成 map 的转换，API 见 https://godoc.org/gopkg.in/mgo.v2


mgo 中 `collection := session.DB("bookmarkdb").C("bookmarks")` 是 concurrency-safe 的，可以在多个 goroutine 中使用，session 有三种可用的一致性 mode，eventually，monotonic，strong，默认是 strong，意味着所有的读和写都会使用 primary 节点来保证数据的一致性，可以通过 session setMode 来完成设置

## vscode 

vscode 的 go 环境和好用，但是由于 golang 官网被墙，最好先解决这个问题，然后 vscode 才能顺利安装需要的各种包

https://code.visualstudio.com/docs/languages/go