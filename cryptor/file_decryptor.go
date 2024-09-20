package cryptor

import (
	"fmt"
	"os"
	"path/filepath"

	"foldcrypt/cout"
	"foldcrypt/files"
	"foldcrypt/utiles"
)

func getfileName(file string) string {
	ext := filepath.Ext(file)
	filename := file[0 : len(file)-len(ext)]
	return filename
}

func decryptfile(cipherFileName string, delCipher bool, key string) error {
	filename := getfileName(cipherFileName)

	err := files.CreateFile(filename)
	if err != nil {
		cout.Error(fmt.Sprintf("%s unable to create output file", filename))
		fmt.Println(err)
		return err
	}

	err = Decrypt(cipherFileName, filename, []byte(key))
	if err != nil {
		cout.Error(fmt.Sprintf("%s unable to decrypt file", filename))
		os.Remove(filename)
		fmt.Println(err)
		return err
	}

	if delCipher {
		err = os.Remove(cipherFileName)
		if err != nil {
			fmt.Println(err)
		}
	}

	cout.Positive(fmt.Sprintf("%s file decrypted : %s", cipherFileName, filename))

	return nil

}

func DecryptDir(args []string, key string, deleteCipher bool, ignoreErr bool) {

	dirName := args[0]

	Files := args[1:]

	if len(Files) != 0 {
		for _, f := range Files {
			fpath := filepath.Join(dirName, f)

			if !files.IsExist(fpath) {
				cout.Error(fmt.Sprintf("%s file not found", f))
				if ignoreErr {
					continue
				} else {
					return
				}
			}

			if filepath.Ext(fpath) == ".encry" {
				err := decryptfile(fpath, deleteCipher, key)
				if err != nil {
					if !ignoreErr {
						return
					}
				}
			}
		}
		return
	}

	queue := utiles.NewQueue()
	queue.Push(dirName)
	var currentDir string

	for !queue.IsEmpty() {
		currentDir = queue.Pop()
		contents, err := files.ReadDir(currentDir)
		if err != nil {
			cout.Error(fmt.Sprintf("%s unable to open directory", dirName))
			fmt.Println(err)
			return
		}

		for _, content := range contents {

			if content.Name() == HashFile {
				continue
			}

			fpath := filepath.Join(currentDir, content.Name())

			if content.IsDir() {
				queue.Push(fpath)

			} else {

				if filepath.Ext(fpath) == ".encry" {
					err := decryptfile(fpath, deleteCipher, key)
					if err != nil {
						if !ignoreErr {
							return
						}
					}
				}

			}
		}
	}

}
