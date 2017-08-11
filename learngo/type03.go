package main

type NewString string

func (n *NewString) name() {

}

//两个 named type 不能相互赋值
func main() {
	var my string = "a"
	var you NewString = my //cannot use my (type string) as type NewString in assignment
}
