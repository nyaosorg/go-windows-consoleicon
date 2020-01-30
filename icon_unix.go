// +build !windows

package consoleicon

func setConsoleExeIcon() (func(bool), error) {
	return func(bool) {}, nil
}
