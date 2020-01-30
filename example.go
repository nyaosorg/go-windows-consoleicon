// +build example

package main

import (
	"os"

	"github.com/zetamatta/go-windows-consoleicon"
)

const icon_restore = true

func main() {
	closer, err := consoleicon.SetConsole(`C:\Windows\System32\notepad.exe`)
	if err != nil {
		return
	}
	defer closer(icon_restore)

	var dummy [1]byte
	os.Stdin.Read(dummy[:])
}
