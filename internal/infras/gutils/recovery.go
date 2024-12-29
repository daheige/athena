package gutils

import "runtime/debug"

// CatchStack 捕获指定stack信息,一般在处理panic/recover中处理
// 返回完整的堆栈信息和函数调用信息
func CatchStack() []byte {
	return debug.Stack()
}
