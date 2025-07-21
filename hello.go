//go:build windows

package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var (
	user32      = windows.NewLazySystemDLL("user32.dll")
	messageBeep = user32.NewProc("MessageBeep")
)

func main() {
	fmt.Println("Hello, World!")
	ret, test, err := messageBeep.Call(0x00000010)

	fmt.Printf("MessageBeep returned: %d, test: %d, error: %v\n", ret, test, err)
}
