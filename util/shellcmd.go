package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func ExecBuf(cmdstr string) (string, error) {
	ci := exec.Command("bash", "-c", cmdstr)
	var stdoutBuf, stderrBuf bytes.Buffer
	ci.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	ci.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := ci.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(stdoutBuf.String()), nil
}

func ExecInteractive(cmdstr string) (string, error) {
	ci := exec.Command("bash", "-c", cmdstr)
	ci.Stdin = os.Stdin
	ci.Stdout = os.Stdout
	ci.Stderr = os.Stderr
    err := ci.Run()
	if err != nil {
		fmt.Println(err)
	}
    return "", nil
}

func ExecAllOut(cmdstr string) (string, error) {
	ci := exec.Command("bash", "-c", cmdstr)
    out, err := ci.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}
