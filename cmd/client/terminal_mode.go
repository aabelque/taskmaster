package main

import (
	"os"
	"os/exec"
)

func set_raw_mode() {
	raw_mode := exec.Command("/bin/stty", "raw")
	raw_mode.Stdin = os.Stdin
	_ = raw_mode.Run()
}

func unset_raw_mode() {
	raw_mode_off := exec.Command("/bin/stty", "-raw")
	raw_mode_off.Stdin = os.Stdin
	_ = raw_mode_off.Run()
}
