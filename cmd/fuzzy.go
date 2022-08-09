package cmd

import (
	"fmt"

	"github.com/shollingsworth/stil/util"
	"github.com/spf13/cobra"
)

var FuzzyCmd = &cobra.Command{
	Use:     "fuzzy",
	Short:   "Fuzzy search",
	Aliases: []string{"f"},
}

var FuzzySubVimMultiCmd = &cobra.Command{
	Use:     "vim-multifile [directory] [match string]",
	Short:   "vim multiple files",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			util.MultiFileVim(args[0], "")
		} else {
			util.MultiFileVim(args[0], args[1])
		}
	},
}

var FuzzySubCopyDirCmd = &cobra.Command{
	Use:     "copydir [source] [destination]",
	Short:   "Copy directory",
	Aliases: []string{"cd"},
	Run: func(cmd *cobra.Command, args []string) {
		src := args[0]
		dst := args[1]
		out, err := util.FuzzyCopy(src, dst)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(out)
	},
}

func init() {
	FuzzySubCopyDirCmd.Args = cobra.MinimumNArgs(2)
	FuzzySubVimMultiCmd.Args = cobra.MinimumNArgs(1)
	FuzzyCmd.AddCommand(FuzzySubVimMultiCmd)
	FuzzyCmd.AddCommand(FuzzySubCopyDirCmd)
	rootCmd.AddCommand(FuzzyCmd)
}