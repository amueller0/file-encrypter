package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func xorEncryptDecrypt(data, key []byte) []byte {
	output := make([]byte, len(data))
	keyLen := len(key)

	for i := range data {
		output[i] = data[i] ^ key[i%keyLen]
	}

	return output
}

func processFile(path string, outputPath string, key []byte) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	encryptedData := xorEncryptDecrypt(data, key)
	return os.WriteFile(outputPath, encryptedData, 0644)
}

func processDirectory(inputDir string, outputDir string, key []byte) error {
	return filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(inputDir, path)
		if err != nil {
			return err
		}

		outputPath := filepath.Join(outputDir, relativePath)
		if info.IsDir() {
			return os.MkdirAll(outputPath, os.ModePerm)
		}

		return processFile(path, outputPath, key)
	})
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ec.exe <path> <key>")
		return
	}

	inputPath := os.Args[1]
	// outputPath := os.Args[2]
	outputPath := inputPath
	key := []byte(os.Args[2])

	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		fmt.Printf("Failed to access input path: %v\n", err)
		return
	}

	if fileInfo.IsDir() {
		err = processDirectory(inputPath, outputPath, key)
	} else {
		err = processFile(inputPath, outputPath, key)
	}

	if err != nil {
		fmt.Printf("Failed to process: %v\n", err)
		return
	}

	fmt.Println("Operation completed successfully.")
}
