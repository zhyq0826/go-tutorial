package main

func main() {
	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	a := [3]int{1, 2, 3} //复合声明
	for i := 0; i < len(a); i++ {
		println(a[i])
	}

	b := [...]int{1, 2, 3} //复合声明
	for i := 0; i < len(b); i++ {
		println(b[i])
	}

	//invalid type for composite literal: int 数组声明 必须是 []int
	c := [3][2]int{[...]int{1, 2}, [...]int{1, 2}, [...]int{1, 2}}
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			println(c[i][j])
		}
	}

	sl := make([]int, 10)
	print(len(sl))

	slice := arr[0:5]
	print(cap(slice)) //cap 指向的实际的底层数组的长度 10
	print(len(slice)) //len 只表示当前类型的长度 5
	// slice[11] = 10 runtime error: index out of range slice 实际指向底层的 array

	s1 := append(slice, 5)
	s2 := append(slice, 5, 5, 5)
	// s3 := append(slice, arry...)
	print(s1)
	print(s2)
	// print(s3)
}
