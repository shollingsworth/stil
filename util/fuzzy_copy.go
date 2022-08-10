package util

import (
	"fmt"
)

func FuzzyCopy(src, dest string) (string, error) {
	cmdstr := fmt.Sprintf("find %s -type d -maxdepth 4 | fzf", src)
	srcdir, err := ExecBuf(cmdstr)
	if err != nil {
		return "", err
	}
	if srcdir == "" {
		return "", nil
	}
	cmdstr = fmt.Sprintf("cp -i -a %s %s", srcdir, dest)
	out, err := ExecInteractive(cmdstr)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
