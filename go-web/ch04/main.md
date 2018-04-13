## Request

request 包括

- URL
- Header
- Body
- Form
- PostForm
- multipartForm

## Header 

Header 支持 get，add，delete，set 操作，Header 是一个 map 结构，key 是 string，value 是一个 slice，set 操作设置一个空的 slice，add 操作是在 key 对应的 value 中加一个新的值。

Header 的 get 操作返回一个字符串，直接取返回一个 map。

## Body

Body 是一个 io.ReadCloser Interface，有  Read 和 Close 方法，Read 接受一个 []byte ，返回 Body 的数据。

## Form

html form 有两种类型，enctype=application/x-www-form-urlencoded, multipart/form-data，不同的类型处理的 form 表单数据不同，前者是通过连字符&做数据拼接，后者是通过
```
----webkitformboudn
content-dispotion: form-data; name="first_name"

```

Go 中使用 ParseForm，ParseMultipartForm 解析 form 数据，然后从中获取表单数据。

Form 支持 urlencode 的方式，包括 url query 参数。

PostForm 仅仅包括 form 中的数据。

multiform 支持带文件的数据，返回两个 map ，一个是表单数据，一个是文件 map，multiform 返回的数据是一个 struct。

## Form value 

FormValue 支持不进行 parseForm 直接获取 form 数据，同理有 PostFormValue。FormValue 获取的是 form 中 key 的第一个 value，formvalue 只支持 enctype=application/x-www-form-urlencoded，他们是通过调用 ParseMultipartForm 实现的


## File form

Go 处理文件是通过 multiForm.File 实现的，同样也提供了 FormFile 来简化文件的处理

## Response

ResponseWriter 是一个 interface 包含有
- WriteHeader 设置 response status code
- Header 设置 response header
- Write  设置 response body，接受 byte

同样 Header().Set, Header().Add 代表了添加和新增

## Cookie

Cookie 这个 struct 存在 http 包中，并且除了使用 ResponseWriter.Header() 设置 cookie 之外，http 包还提供了 SetCookie 来简化对 cookie 的修改
