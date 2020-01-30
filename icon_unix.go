// +build !windows

package consoleicon

func setConsole(fname string) (func(bool), error) {
	return func(bool) {}, nil
}

func setConsoleExe() (func(bool), error) {
	return func(bool) {}, nil
}
