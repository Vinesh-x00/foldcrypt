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

// lockCmd represents the lock command
var lockCmd = &cobra.Command{
	Use:                   "lock [OPTIONS]... [DIR]",
	Short:                 "Encrypts the specified directory and all subdirectories using the AES cipher",
	Example:               "  foldcrypt lock -n /home/secret",
	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("non-delete")
		s, _ := cmd.Flags().GetBool("non-shread")

		if !files.IsDirectory(args[0]) {
			fmt.Printf("%s not directory", args[0])
			return
		}

		hashFilePath := filepath.Join(args[0], HashFile)

		if !files.IsExist(hashFilePath) {
			err := files.CreateFile(hashFilePath)
			if err != nil {
				fmt.Println("Unable to create hash file")
				fmt.Println(err)
				return
			}
		}

		file, err := os.OpenFile(hashFilePath, os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Encrypting dir : %s\n", args[0])
		fmt.Print("Password ('C' for cancel): ")

		password := utiles.AskPasswd()
		fmt.Print("\n")

		if password == "$ERR$" {
			fmt.Println("Unable to get passwd")
		} else if password == "C" {
			fmt.Println("Encryption canceled")
		} else {
			hashedPasswd := utiles.GetHash(password)
			file.Write([]byte(hashedPasswd))
			file.Close()
			cryptor.EncryptDir(args[0], password, !v, !s)
		}

	},
}

func init() {
	lockCmd.Flags().BoolP("non-delete", "n", false, "Does not delete the original file after encryption")
	lockCmd.Flags().Bool("non-shread", false, "Does not shred files for deletion")
	rootCmd.AddCommand(lockCmd)
}
