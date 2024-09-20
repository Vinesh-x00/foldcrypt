package files

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ReadDir(root string) ([]fs.FileInfo, error) {
	f, err := os.Open(root)
	if err != nil {
		return nil, err
	}

	fileInfo, err := f.Readdir(0)
	if err != nil {
		return fileInfo, err
	}

	return fileInfo, nil
}

func IsRegular(file os.File) bool {
	stat, err := file.Stat()
	if err != nil {
		return false
	}
	return stat.Mode().IsRegular()
}

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func CreateUniqueFile(fileName string) (string, error) {
	baseName := fileName
	ext := filepath.Ext(fileName)
	nameWithoutExt := fileName[:len(fileName)-len(ext)]

	for i := 0; i < 100; i++ {
		//ignored i = 0, to create file
		if i > 0 {
			fileName = fmt.Sprintf("%s_%d%s", nameWithoutExt, i, ext)
		}

		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if err == nil {
			file.Close()
			return fileName, nil
		}

		//if file exist , cotinue the loop
		if !os.IsExist(err) {
			return "", err
		}
	}

	return "", fmt.Errorf("failed to create  file afterattempts: %s", baseName)
}

func CreateFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err == nil {
		file.Close()
		return nil
	} else {
		return err
	}
}

func IsExist(fileName string) bool {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
