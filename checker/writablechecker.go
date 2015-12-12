package checker

import (
	"os/exec"
	"strings"
	"errors"
	"fmt"
)

func CheckIfPathWritable(directory string) error {
	cmd := exec.Command(
		"bin/adb/adb.exe",
		"shell",
		"echo",
		"writable",
		">",
		directory + "checkwritable.log")

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	if strings.Contains(string(output), "Read-only file system") {
		return errors.New(
			fmt.Sprintf("directory %v to store pcap file not writable", directory))
	}

	return nil
}
