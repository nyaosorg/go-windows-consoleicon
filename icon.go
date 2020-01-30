package consoleicon

func SetFrom(fname string) (func(bool), error) {
	return setFrom(fname)
}

func SetFromExe() (func(bool), error) {
	return setFromExe()
}
