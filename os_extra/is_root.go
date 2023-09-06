package osextra

import (
	"os"
	"os/exec"
	"os/user"
	"runtime"
)

func IsRoot() bool {
	if runtime.GOOS == "windows" {
		// TODO: write check for Windows
		return true
	}
	// Note: ignoring errors
	currentUser, _ := user.Current()
	return currentUser.Uid == "0"
}

func MakeRoot() error {
	if IsRoot() {
		return nil
	}
	if runtime.GOOS == "windows" {
		// TODO: write running runas on Windows
		return nil
	}
	cmd := exec.Command("sudo", os.Args...)
	return cmd.Run()
}
