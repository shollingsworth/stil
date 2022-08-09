package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func FuzzyCopy(src, dest string) (string, error) {
	cmdstr := fmt.Sprintf("find %s -type d -maxdepth 4 | fzf", src)
	ci := exec.Command("bash", "-c", cmdstr)
	var stdoutBuf, stderrBuf bytes.Buffer
	ci.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	ci.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := ci.Run()
	if err != nil {
		return "", err
	}
	srcdir := strings.TrimSpace(stdoutBuf.String())
	if srcdir == "" {
		return "", nil
	}
	cmdstr = fmt.Sprintf("cp -a %s %s", srcdir, dest)
	out, err := exec.Command("bash", "-c", cmdstr).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
