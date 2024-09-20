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

// unlockCmd represents the unlock command
var unlockCmd = &cobra.Command{
	Use:                   "unlock [OPTIONS]... [DIR] [FILE]... \n\nWith no FILE hole directory will be decrypted",
	Short:                 "Decrypts the specified directory and all subdirectories",
	Example:               "  foldcrypt unlock -i /home/secret file1.txt file2.txt",
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		nd, _ := cmd.Flags().GetBool("non-delete")
		ignore, _ := cmd.Flags().GetBool("ignore")
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

		fmt.Printf("Decrypting dir : %s\n", args[0])
		fmt.Print("Password ('C' for cancel): ")

		password := utiles.AskPasswd()
		fmt.Print("\n")

		if password == "$ERR$" {
			fmt.Println("Unable to get passwd")
		} else if password == "C" {
			fmt.Println("Decryption canceled")
		} else {
			if hashVerify || string(data) == utiles.GetHash(password) {
				cryptor.DecryptDir(args, password, !nd, ignore)
			} else {
				fmt.Println("INCORRECT passwd")
			}
		}

	},
}

func init() {
	unlockCmd.Flags().BoolP("non-delete", "n", false, "Does not delete the cipher file after decryption")
	unlockCmd.Flags().BoolP("ignore", "i", false, "Ignore decryption errors")
	unlockCmd.Flags().BoolP("unverilfy", "u", false, "Does not verify the password hash")
	rootCmd.AddCommand(unlockCmd)
}
