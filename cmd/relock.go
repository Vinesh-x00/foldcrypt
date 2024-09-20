/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"foldcrypt/cryptor"
	"foldcrypt/files"
	"foldcrypt/utiles"
)

// relockCmd represents the relock command
var relockCmd = &cobra.Command{
	Use:                   "relock [OPTIONS]... [DIR]",
	Short:                 "Rencrypts directories and subdirectories when new files are added to locked directories",
	Example:               "  foldcrypt relock -n /home/secret",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("non-delete")
		s, _ := cmd.Flags().GetBool("non-shread")
		hashVerify, _ := cmd.Flags().GetBool("unverilfy")

		if !files.IsDirectory(args[0]) {
			fmt.Printf("%s not directory", args[0])
			return
		}

		hashFilePath := filepath.Join(args[0], HashFile)

		data, err := os.ReadFile(hashFilePath)
		if err != nil {
			fmt.Println("Unable to read hash file")
			return
		}

		fmt.Printf("Recrypting dir : %s\n", args[0])
		fmt.Print("Password ('C' for cancel): ")

		password := utiles.AskPasswd()
		fmt.Print("\n")

		if password == "$ERR$" {
			fmt.Println("Unable to get passwd")
		} else if password == "C" {
			fmt.Println("Encryption canceled")
		} else {
			if hashVerify || string(data) == utiles.GetHash(password) {
				cryptor.ReEncryptDir(args[0], string(password), !v, !s)
			} else {
				fmt.Println("INCORRECT passwd")
			}
		}
	},
}

func init() {
	relockCmd.Flags().BoolP("non-delete", "n", false, "Does not delete the original file after encryption")
	relockCmd.Flags().Bool("non-shread", false, "Does not shred files for deletion")
	relockCmd.Flags().BoolP("unverilfy", "u", false, "Does not verify the password hash")
	rootCmd.AddCommand(relockCmd)
}
