package consoleicon

func SetConsole(fname string) (func(bool), error) {
	return setConsole(fname)
}

func SetConsoleExe() (func(bool), error) {
	return setConsoleExe()
}
