## template

go 提供了标准的 html/template ，ParseFiles 支持多个文件，模板执行的过程是首先进行模板解析，然后再执行数据绑定，也就是 parase 和 execute 的过程

template 还支持比较简单的写法，直接使用 Must, Must 会处理 ParseFiles 报出来的 Error


## 多个 files

多个 file parse execute 只会执行一个，默认是第一个，如果要执行其他的，需要显示指定 filename, name 就是文件的名称

## actions

template 支持多种 actions

- Conditional actions
- Iterator actions
- Set actions
- Include action
- Data actions .


## if action



## range action


 range 支持 array, slice, map, or channel


## set action


set action 可以改变 dot 的上下文


## include

include 可以包含另一个 template，注意 parseFile 必须要包括需要 include 的 file，include file 的 data 可以传递进去


## Arguments, variables, and pipelines

$ 符号可以声明一个 variable

```
{{ range $key, $value := . }}
 The key is {{ $key }} and the value is {{ $value }}
{{ end }}
```

pipeline 类似 jinja2 的 pipeline
```
{{ p1 | p2 | p3 }}
```

## func

func 类似于 jinja2 的标签，事先绑定一个 func 到 template 实现更强大的功能

## context awareness

