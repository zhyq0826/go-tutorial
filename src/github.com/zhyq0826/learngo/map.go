package main

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for _, v := range m {
		print(v)
	}

	v, ok := m["a"]
	print(v)
	print(ok)

	delete(m, "c")
	v1, ok1 := m["c"]
	print(v1)
	print(ok1)
}
