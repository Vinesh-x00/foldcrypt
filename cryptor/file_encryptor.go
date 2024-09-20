package cryptor

import (
	"fmt"
	"os"
	"path/filepath"

	"foldcrypt/cout"
	"foldcrypt/files"
	"foldcrypt/utiles"
)

var HashFile = ".foldcrypthash"

func encryptfile(fpath, cipherFileName string, delOrgin bool, key string, shred bool) error {

	cipherFile, err := files.CreateUniqueFile(cipherFileName)
	if err != nil {
		cout.Error(fmt.Sprintf("%s unable to create cipher file", fpath))
		return err
	}

	err = encrypt(fpath, cipherFile, []byte(key))
	if err != nil {
		cout.Error(fmt.Sprintf("%s unable to encrypt file", fpath))
		println(err.Error())
		os.Remove(cipherFileName)
		return err
	}

	if delOrgin {
		if shred {
			files.Shred(fpath, 3)
		}
		err = os.Remove(fpath)
		if err != nil {
			fmt.Println(err)
		}
	}

	cout.Positive(fmt.Sprintf("%s file encrypted : %s", fpath, cipherFileName))
	return nil
}

func EncryptDir(dirName string, key string, deleteOrgin bool, shred bool) {

	queue := utiles.NewQueue()
	queue.Push(dirName)

	var currentDir string

	for !queue.IsEmpty() {
		currentDir = queue.Pop()
		contents, err := files.ReadDir(currentDir)
		if err != nil {
			cout.Error(fmt.Sprintf("%s unable to open directory", dirName))
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
				cipherFileName := fmt.Sprintf("%s.encry", fpath)

				err := encryptfile(fpath, cipherFileName, deleteOrgin, key, shred)
				if err != nil {
					return
				}

			}
		}
	}

}

func ReEncryptDir(dirName string, key string, deleteOrgin bool, shred bool) {

	queue := utiles.NewQueue()
	queue.Push(dirName)

	var currentDir string

	for !queue.IsEmpty() {
		currentDir = queue.Pop()
		contents, err := files.ReadDir(currentDir)
		if err != nil {
			cout.Error(fmt.Sprintf("%s unable to open directory", dirName))
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

				if filepath.Ext(fpath) != ".encry" {

					cipherFileName := fmt.Sprintf("%s.encry", fpath)
					err := encryptfile(fpath, cipherFileName, deleteOrgin, key, shred)
					if err != nil {
						return
					}
				}

			}
		}
	}

}
