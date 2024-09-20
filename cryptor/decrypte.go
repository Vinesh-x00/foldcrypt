package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)

func Decrypt(inputPath, outputPath string, key []byte) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	//padding key
	paddedKey := Pad(key, aes.BlockSize)

	// Create the AES cipher
	block, err := aes.NewCipher(paddedKey)
	if err != nil {
		return err
	}

	// Read the IV from the input file
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(inputFile, iv); err != nil {
		return err
	}

	// Create the cipher stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Create a reader that decrypts from the input file
	reader := &cipher.StreamReader{S: stream, R: inputFile}

	// Copy the decrypted input file to the output file
	if _, err := io.Copy(outputFile, reader); err != nil {
		return err
	}

	return nil
}
