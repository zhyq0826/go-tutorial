package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string
}

// Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，如果要访问当前对象的字段通过{{.FieldName}},但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的),否则在渲染的时候就会报错，请看下面的这个例子：

//如果模板中输出{{.}}，这个一般应用于字符串对象，默认会调用fmt包输出字符串的内容。

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	// t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
	p := Person{UserName: "wecatch"}
	t.Execute(os.Stdout, p)

	t2 := template.New("string example")
	t2, _ = t2.Parse("hello {{.}}") //代表当前对象，一般会应用 fmt 输出字符串
	t2.Execute(os.Stdout, "abcdefg")

}
