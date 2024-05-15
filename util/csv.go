package util

// csv関連の操作

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func ReadFile(filepath string) (string, error){
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
			fmt.Println("File reading error", err)
			return "", err
	}

	return buf.String(), nil
}
