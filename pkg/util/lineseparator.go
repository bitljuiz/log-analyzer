package util

import "runtime"

func CheckLineSeparator() rune {
	if runtime.GOOS == "windows" {
		return '\r'
	}

	return '\n'
}

func LineSeparator() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}
