package main

import (
	"fmt"
	"os"
)

var DefaultRules = map[string]string{
	".jpg":  "Images",
	".jpeg": "Images",
	".png":  "Images",
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".txt":  "Documents",
	".mp3":  "Music",
	".wav":  "Music",
	".mp4":  "Video",
	".avi":  "Video",
	".zip":  "Archives",
	".rar":  "Archives",
}

type FileOrganizer struct {
	sourceDir      string
	rulesMap       map[string]string
	processedFiles int
	logFile        *os.File
}

func NewFileOrganizer(sourceDir string) (*FileOrganizer, error) {
	if sourceDir == "" {
		return nil, fmt.Errorf("sourceDir is empty")
	}
	fileInfo, err := os.Stat(sourceDir)
	if err != nil {
		return nil, fmt.Errorf("path not found: %w", err)
	}
	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("path is not directory")
	}

	return &FileOrganizer{
		sourceDir: sourceDir,
		rulesMap:  DefaultRules,
	}, nil
}

func main() {
	sourceDir := ".test"
	_, err := NewFileOrganizer(sourceDir)
	if err != nil {
		fmt.Printf("Ошибка: %s", err)
	}
	fmt.Printf("FileOrganizer создан для директори: %s\n", sourceDir)
}
