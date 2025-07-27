package checker_test

import "unsafe"

func g1() {
	var a int
	if a == 10 {
	}
}

func g2() bool {
	f := func() interface{} { return nil }
	return f() != nil
}

func _() {
	const foo = 10
	_ = 15 != 10
	_ = 15 == foo

	type myArray struct {
		data [10]int
	}
	var arr myArray
	var i int
	_ = len(arr.data) == i
	_ = i == len(arr.data)

	_ = unsafe.Sizeof(0) == 0

	var c byte
	if c <= '0' && '9' <= c {
		// character range ok
	}
	if c >= '0' && '9' <= c {
		// character range ok
	}
}
