package files

import "runtime"

var (
	root string
)

func init() {
	if runtime.GOOS == "windows" {
		root = "C:"
	}

	if runtime.GOOS == "linux" {
		root = "/"
	}
}
