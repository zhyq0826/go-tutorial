## mgo.session

推荐的方式是每个 request life cycle 使用一个 mgo.session ，因此要从全局 mgo.session 获取一个 copy，mgo.session 提供了 copy 和 clone 操作来完成一个 session 的复制。

copy 复用的是 connections，而 clone 复用的 socket


## vscode gofmttool

vscode gofmtool 默认是 goreturns 会自动删除未使用的包，因为 vscode 默认会自动导入需要用的包路径

https://www.reddit.com/r/golang/comments/4vtoc4/noob_vscode_is_deleting_my_import_line/


## r.Request.Body 

r.Request.Body 是一个 buffer 读取完成之后就不能继续读了，处理 body 中的 json 时需要注意，在使用 json decoder 时注意 body 的读取时机，不然报 EOF 错误
