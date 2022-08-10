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
	cmdstr = fmt.Sprintf("cp -i %s .", out)
	_, err = ExecInteractive(cmdstr)
	fmt.Println(cmdstr)
	return err
}
