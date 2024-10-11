//go:build unix

package cmd

import (
	"io/fs"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/twpayne/chezmoi/v2/internal/chezmoi"
)

const defaultEditor = "vi"

var defaultInterpreters = make(map[string]chezmoi.Interpreter)

func fileInfoUID(info fs.FileInfo) int {
	return int(info.Sys().(*syscall.Stat_t).Uid) //nolint:forcetypeassert
}

func darwinVersion() (map[string]any, error) {
	if runtime.GOOS != "darwin" {
		return nil, nil
	}
	darwinVersion := make(map[string]any)
	cmdOutput, err := exec.Command("sw_vers").Output()
	if err != nil {
		return nil, err
	}
	output := string(cmdOutput)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		key = strings.ToLower(key[:1]) + key[1:]
		value := strings.TrimSpace(parts[1])
		darwinVersion[key] = value
	}
	return darwinVersion, nil
}

func windowsVersion() (map[string]any, error) {
	return nil, nil
}
