package help

import "os"

type print func()

func Check(help bool, fn print) {
	if help {
		fn()
		os.Exit(0)
	}
}
