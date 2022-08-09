package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func MultiFileVim(srcdir string, match_string string) {
	if (match_string == "") {
		match_string = "*"
	} else {
		match_string = "*" + match_string + "*"
	}
	cmdstr := fmt.Sprintf("find %s -type f -iname '%s' | fzf -m", srcdir, match_string)
	ci := exec.Command("bash", "-c", cmdstr)
	var stdoutBuf, stderrBuf bytes.Buffer
	ci.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	ci.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	err := ci.Run()
	out := strings.TrimSpace(stdoutBuf.String())
	files := strings.Split(out, "\n")

	cmdstr = fmt.Sprintf("vim -p %s", strings.Join(files, " "))
	ci = exec.Command("bash", "-c", cmdstr)
	ci.Stdin = os.Stdin
	ci.Stdout = os.Stdout
	ci.Stderr = os.Stderr
	err = ci.Run()
	if err != nil {
		fmt.Println(err)
	}
}
