package main

import (
	"fmt"
	"time"
	"unsafe"

	"syscall"
)

var (
	user32DLL    = syscall.NewLazyDLL("user32.dll")
	getCursorPos = user32DLL.NewProc("GetCursorPos")
)

type POINTER struct {
	x int32
	y int32
}

const sleepTime = 1000

func main() {
	var point POINTER
	unsafePointer := uintptr(unsafe.Pointer(&point))

	for {
		time.Sleep(sleepTime)
		successful, _, _ := getCursorPos.Call(unsafePointer)
		var lastX, lastY int32

		if successful == 1 &&
			(lastX != point.x || lastY != point.y) {
			fmt.Printf("[ %d : %d ]\n", point.x, point.y)
			lastX = point.x
			lastY = point.y
		}

	}

}
