/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"foldcrypt/files"

	"github.com/spf13/cobra"
)

// shreadCmd represents the shread command
var shreadCmd = &cobra.Command{
	Use:                   "shread [OPTIONS]... [FILE... | DIR]",
	Short:                 "A command allows to shread file or directory",
	Example:               "  foldcrypt shread -i 5 /home/secret/file.txt \n  foldcrypt shread -d /home/secret/",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		nd, _ := cmd.Flags().GetBool("non-delete")
		ignore, _ := cmd.Flags().GetBool("ignore")
		iteration, _ := cmd.Flags().GetInt32("iteration")
		isDir, _ := cmd.Flags().GetBool("dir")

		if isDir {
			files.ShreadDir(args[0], int(iteration), !nd, ignore)
		} else {
			files.ShreadFiles(args, int(iteration), !nd, ignore)
		}

	},
}

func init() {
	shreadCmd.Flags().BoolP("non-delete", "n", false, "Does not delete the original file after shread")
	shreadCmd.Flags().BoolP("dir", "d", false, "Indicates the argument is a directory.")
	shreadCmd.Flags().Int32P("iteration", "i", 3, "number of iteration to shread")
	shreadCmd.Flags().Bool("ignore", false, "Ignore errors")
	rootCmd.AddCommand(shreadCmd)
}
