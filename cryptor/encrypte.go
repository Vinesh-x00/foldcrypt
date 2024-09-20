package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func encrypt(inputPath, outputPath string, key []byte) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_APPEND, os.ModePerm)
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

	// Create a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Write the IV to the output file
	if _, err := outputFile.Write(iv); err != nil {
		return err
	}

	// Create the cipher stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Create a writer that encrypts to the output file
	writer := &cipher.StreamWriter{S: stream, W: outputFile}

	// Copy the input file to the encrypted output file
	if _, err := io.Copy(writer, inputFile); err != nil {
		return err
	}

	return nil
}

func Pad(input []byte, blockSize int) []byte {
	r := len(input) % blockSize
	pl := blockSize - r
	for i := 0; i < pl; i++ {
		input = append(input, byte(pl))
	}
	return input
}
