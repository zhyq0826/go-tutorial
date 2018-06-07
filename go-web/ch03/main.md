## 自定义 server 配置

go 通过 net/http 包提供 web 功能，http.Server struct 提供了关于 server 的额外的配置，见 `go doc net/http.Server`

## http.handler

`go doc net/http.Handler`

handler 是一个有 ServeHTTP 方法的 interface 类型，接收两个参数 HTTPResponseWriter interface 和 Request pointer。

```go
ServeHTTP(http.ResponseWriter, *http.Request)
```

handler 一般是一个实现了 Handler interface 的实例，比如一个实现了 Handler 的 struct。

> An http.Handler wrapper is a function that has one input argument and one output argument, both of type http.Handler.

## ServerMux

Go Doc 解 `go doc net/http.ServeMux`

> ServeMux is an HTTP request multiplexer. It matches the URL of each incoming
request against a list of registered patterns and calls the handler for the
pattern that most closely matches the URL

也就是 ServeMux 是一个满足 handler 的 struct，它包含一个方法是 `func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)`

## http.Handle

同时是 ServeMux `func (mux *ServeMux) Handle(pattern string, handler Handler)
` 方法，该方法的主要作用就是注册 pattern 和 handler 到 defaultServeMux，本质上是 ServeMux 中的 Handle。

注意 handler 是必须实现了 ServeHTTP 的 instance

## http.HandlerFunc 

是一个类型，用以把普通函数转换成一个 Handler，也就是具有 ServeHTTP 的函数

## handler function

handler function 是行为表现和 Handler 一样的 function， 他们的 ServeHTTP 函数参数签名是一样的

```go
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

```

handler function 的作用就是把一个 函数转换成一个 handler，并且把该 handler 注册到 serverMux 
```
mux.Handle(pattern, HandleFunc(handler))
```

## chain handler 链式调用 handler

```go

func (handler http.Handler) http.Handler {
    return http.HandlerFunc(func (ResponseWriter, *Request){
        
    })
}
```



