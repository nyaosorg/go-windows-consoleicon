//go:build !windows
// +build !windows

package consoleicon

func setFrom(fname string) (func(bool), error) {
	return func(bool) {}, nil
}

func setFromExe() (func(bool), error) {
	return func(bool) {}, nil
}
