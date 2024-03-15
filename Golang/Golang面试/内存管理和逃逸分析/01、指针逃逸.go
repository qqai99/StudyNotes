package main

import "fmt"

/**
 * @Author kong
 * @Date 2024/3/15
 * @Description 指针逃逸
 */

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) // 局部变量s 逃逸到堆上
	s.Name = name
	s.Age = age

	return s
}

func main() {
	fmt.Println(StudentRegister("kong", 24))
}

// .\01、指针逃逸.go:26:29: inlining call to StudentRegister
// .\01、指针逃逸.go:26:13: inlining call to fmt.Println
// .\01、指针逃逸.go:16:22: leaking param: name
// .\01、指针逃逸.go:18:10: new(Student) escapes to heap
// .\01、指针逃逸.go:26:13: ... argument does not escape
// .\01、指针逃逸.go:26:29: new(Student) escapes to heap
