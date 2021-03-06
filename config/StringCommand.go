package config

import (
	"os/exec"
	"regexp"
	"strings"
)

type StringCommand string

var reCmd = regexp.MustCompile("\\${.*}")

func (sc StringCommand) Process() string {
	return reCmd.ReplaceAllStringFunc(string(sc), func(s string) string {
		command := s[2:len(s)-1]
		cmd := exec.Command("sh", "-c", command)

		str, err := cmd.CombinedOutput()
		if err != nil {
			return "[ERR]"
		}

		return strings.ReplaceAll(string(str), "\n", "")
	})
}
