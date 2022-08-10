package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func MultiFileVim(srcdir string, match_string string) {
	if match_string == "" {
		match_string = "*"
	} else {
		match_string = "*" + match_string + "*"
	}
	cmdstr := fmt.Sprintf("find %s -type f -iname '%s' | fzf -m", srcdir, match_string)
	out, err := ExecBuf(cmdstr)
	if out == "" {
		fmt.Println("No files selected")
		return
	}
	files := strings.Split(out, "\n")

	cmdstr = fmt.Sprintf("vim -p %s", strings.Join(files, " "))
	ci := exec.Command("bash", "-c", cmdstr)
	ci.Stdin = os.Stdin
	ci.Stdout = os.Stdout
	ci.Stderr = os.Stderr
	err = ci.Run()
	if err != nil {
		fmt.Println(err)
	}
}
