package main

/**
 * @Author kong
 * @Date 2024/3/15
 * @Description 栈空间不足逃逸
 */
// 当切片长度扩大到10000时就会逃逸。
// 实际上栈空间不足以存放当前对象时，或无法判断当前切片长度时将会分配到堆中。

func Slice() {
	s := make([]int, 10000)

	for index, _ := range s {
		s[index] = index
	}
}

func main() {
	Slice()
}

// .\02、栈空间不足逃逸.go:11:6: can inline Slice
// .\02、栈空间不足逃逸.go:19:6: can inline main
// .\02、栈空间不足逃逸.go:20:7: inlining call to Slice
// .\02、栈空间不足逃逸.go:12:11: make([]int, 10000) escapes to heap
// .\02、栈空间不足逃逸.go:20:7: make([]int, 10000) escapes to heap
