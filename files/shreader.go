package files

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"foldcrypt/cout"
	"foldcrypt/utiles"
	"os"
	"path/filepath"
)

func overwrite(path string) error {
	// variable declaration
	var err error

	fd, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	// use buffered io for performance
	wr := bufio.NewWriter(fd)
	randbuff := make([]byte, 4096)

	fileStat, err := fd.Stat()
	if err != nil {
		return err
	}

	// write random data in blocks of 4K when possible
	for sizeleft := fileStat.Size(); sizeleft > 0; sizeleft -= 4096 {
		if sizeleft < 4096 {
			// last block smaller than 4K sometimes
			randbuff = make([]byte, sizeleft)
		}
		rand.Read(randbuff)
		wr.Write(randbuff)
	}

	wr.Flush()
	fd.Close()
	return nil
}

func Shred(path string, iteration int) error {
	var err error
	for i := 0; i < iteration; i++ {
		err = overwrite(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func ShreadDir(dirName string, iteraion int, delFile, ignore bool) {
	queue := utiles.NewQueue()
	queue.Push(dirName)

	var currentDir string

	for !queue.IsEmpty() {
		currentDir = queue.Pop()
		contents, err := ReadDir(currentDir)

		if err != nil {
			cout.Error(fmt.Sprintf("%s unable to open directory", dirName))
			return
		}

		for _, content := range contents {

			fpath := filepath.Join(currentDir, content.Name())

			if content.IsDir() {
				queue.Push(fpath)

			} else {

				err := Shred(fpath, iteraion)
				if err != nil {
					cout.Error(fmt.Sprintf("%s unable to shread file", fpath))
					fmt.Println(err)
					if !ignore {
						return
					}
				} else {
					cout.Positive(fmt.Sprintf("%s shreaded", fpath))
				}

				if delFile {
					os.Remove(fpath)
				}

			}
		}
	}

}

func ShreadFiles(fileNames []string, iteration int, delFile, ignore bool) {
	for _, file := range fileNames {
		err := Shred(file, iteration)
		if err != nil {
			cout.Error(fmt.Sprintf("%s unable to shread file", file))
			if !ignore {
				return
			}
		} else {
			cout.Positive(fmt.Sprintf("%s shreaded", file))
		}

		if delFile {
			os.Remove(file)
		}
	}
}
