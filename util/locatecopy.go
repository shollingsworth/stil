package util

import "fmt"

func LocateCopy(search string) error {
	cmdstr := fmt.Sprintf("locate %s | fzf", search)
	out, err := ExecBuf(cmdstr)
	if err != nil {
		return err
	}
	if out == "" {
		return nil
	}
	cmdstr = fmt.Sprintf("cp %s .", out)
	_, err = ExecAllOut(cmdstr)
	fmt.Println(cmdstr)
	return err
}
