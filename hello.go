//go:build windows

package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

// BeepType defines the sound to be played by the MessageBeep function.
//
// See https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messagebeep.
type BeepType uint32

const (
	// A simple beep. If the sound card is not available, the sound is generated using the speaker.
	SimpleBeep BeepType = 0xFFFFFFFF
	// See MB_ICONINFORMATION.
	MB_ICONASTERISK BeepType = 0x00000040
	// See MB_ICONWARNING.
	MB_ICONEXCLAMATION BeepType = 0x00000030
	// The sound specified as the Windows Critical Stop sound.
	MB_ICONERROR BeepType = 0x00000010
	// See MB_ICONERROR.
	MB_ICONHAND BeepType = 0x00000010
	// The sound specified as the Windows Asterisk sound.
	MB_ICONINFORMATION BeepType = 0x00000040
	// The sound specified as the Windows Question sound.
	MB_ICONQUESTION BeepType = 0x00000020
	// See MB_ICONERROR.
	MB_ICONSTOP BeepType = 0x00000010
	// The sound specified as the Windows Exclamation sound.
	MB_ICONWARNING BeepType = 0x00000030
	// The sound specified as the Windows Default Beep sound.
	MB_OK BeepType = 0x00000000
)

// A function that generated a random BeepType and plays it using the MessageBeep function.
// It returns the result of the call and any error encountered.
func playRandomBeep() (ret uintptr, err error) {
	// Generate a random BeepType.
	beepType := BeepType(0x00000010) // For example, using MB_ICONERROR.

	// Call the MessageBeep function with the generated BeepType.
	ret, _, err = messageBeep.Call(uintptr(beepType))
	if err != nil && err.Error() != "The operation completed successfully." {
		return 0, fmt.Errorf("MessageBeep failed: %w", err)
	}

	return ret, nil
}

var (
	user32      = windows.NewLazySystemDLL("user32.dll")
	messageBeep = user32.NewProc("MessageBeep")
)

func main() {
	fmt.Println("Hello, World!")
	ret, test, err := messageBeep.Call(0x00000010)

	fmt.Printf("MessageBeep returned: %d, test: %d, error: %v\n", ret, test, err)
}
