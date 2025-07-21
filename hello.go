//go:build windows

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/sys/windows"
)

// ANSI color codes for nicer output in the terminal.
const (
	CyanAnsi   = "\033[36m"
	GrayAnsi   = "\033[90m"
	YellowAnsi = "\033[33m"
	ResetAnsi  = "\033[0m"
)

// BeepType represents a type of beep sound that can be played using the MessageBeep function.
type BeepType struct {
	Name  string
	Value uint32
}

// beepTypes contains the different types of sounds that can be played by the MessageBeep function.
//
// See https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messagebeep.
var beepTypes = []BeepType{
	// A simple beep. If the sound card is not available, the sound is generated using the speaker.
	{"SimpleBeep", 0xFFFFFFFF},
	// See MB_ICONINFORMATION.
	{"MB_ICONASTERISK", 0x00000040},
	// See MB_ICONWARNING.
	{"MB_ICONEXCLAMATION", 0x00000030},
	// The sound specified as the Windows Critical Stop sound.
	{"MB_ICONERROR", 0x00000010},
	// See MB_ICONERROR.
	{"MB_ICONHAND", 0x00000010},
	// The sound specified as the Windows Asterisk sound.
	{"MB_ICONINFORMATION", 0x00000040},
	// The sound specified as the Windows Question sound.
	{"MB_ICONQUESTION", 0x00000020},
	// See MB_ICONERROR.
	{"MB_ICONSTOP", 0x00000010},
	// The sound specified as the Windows Exclamation sound.
	{"MB_ICONWARNING", 0x00000030},
	// The sound specified as the Windows Default Beep sound.
	{"MB_OK", 0x00000000},
}

var (
	user32      = windows.NewLazySystemDLL("user32.dll")
	messageBeep = user32.NewProc("MessageBeep")
)

// playRandomBeep plays a random beep sound from the predefined beep types and prints the result.
func playRandomBeep() {
	b := beepTypes[rand.Intn(len(beepTypes))]
	ret, _, err := messageBeep.Call(uintptr(b.Value))
	if ret != 0 {
		fmt.Printf("Played beep: %s\n", b.Name)
	} else {
		fmt.Printf("Failed to play beep: %s, error: %v\n", b.Name, err)
	}
}

func main() {
	// Print the program title, description, and usage instructions with ANSI color codes
	fmt.Println(CyanAnsi + "hello-world-go-beep" + ResetAnsi)
	fmt.Printf("%sA Hello, World! program in Go that plays random system beep sounds.%s\n\n", GrayAnsi, ResetAnsi)
	fmt.Printf("%sUsage:%s hello-world-go-beep [number_of_beeps_to_play]\n\n", GrayAnsi, ResetAnsi)

	// Get the number of beeps to play from the (single) command line argument
	flag.Parse()
	count, err := strconv.Atoi(flag.Arg(0))
	if err != nil || count <= 0 {
		// If no valid count is provided, use a random count between 1 and 100
		fmt.Printf("%sInvalid or no count provided, deciding how many beeps to play randomly.%s\n\n", YellowAnsi, ResetAnsi)
		count = rand.Intn(100) + 1
	}

	// Start playing the random system beep sounds
	fmt.Printf("Playing a random system beep sound %s%d times%s:\n", CyanAnsi, count, ResetAnsi)
	for range count {
		playRandomBeep()
		time.Sleep(1 * time.Second)
	}
}
