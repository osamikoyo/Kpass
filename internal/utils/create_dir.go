package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDir() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil{
		return "", fmt.Errorf("cant get user home dir: %w", err)
	}

	dirpath := filepath.Join(homedir, ".kpass")
	_ = os.Mkdir(dirpath, 0755)

	return filepath.Join(dirpath, "main.db"), nil
}