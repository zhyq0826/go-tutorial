package main

type Ptr *int
type Map map[int]string
type MapMap Map

func main() {
	var p *int
	var mm Map
	var mmm MapMap
	var m1 map[int]string = mm
	var m2 map[int]string = mmm
	var ptr Ptr = p
	print(ptr)
	print(m1)
	print(m2)
}
