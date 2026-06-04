package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func (fo *FileOrganizer) initLog() error {
	file, err := os.OpenFile("organizer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	fo.logFile = file
	log.SetOutput(fo.logFile)
	return nil
}

func (fo *FileOrganizer) logSuccess(message string) {
	logMessage := fmt.Sprintf("[SUCCESS] %s", message)
	log.Println(logMessage)
}

func (fo *FileOrganizer) logError(message string) {
	logMessage := fmt.Sprintf("[ERROR] %s", message)
	log.Println(logMessage)
}

func (fo *FileOrganizer) Close() error {
	if fo.logFile != nil {
		if err := fo.logFile.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (fo *FileOrganizer) moveFile(sourcePath string, targetDir string) error {
	fileName := filepath.Base(sourcePath)
	resultDirPath := filepath.Join(fo.sourceDir, targetDir)
	if _, err := os.Stat(resultDirPath); err != nil {
		if err := os.MkdirAll(resultDirPath, 0750); err != nil {
		}
		return fmt.Errorf("error when create directory: %s", err)
	}
	resultFilePath := filepath.Join(resultDirPath, fileName)

	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	newFileName := fileNameWithoutExt + "_" + time.Now().Format("2006-01-02_15-04-05") + filepath.Ext(fileName)
	resultFilePath = filepath.Join(resultDirPath, newFileName)
	os.Rename(sourcePath, resultFilePath)
	fo.logSuccess("move completed")
	return nil
}

func main() {
	fo, err := NewFileOrganizer("/home/user/messy_folder")
	if err != nil {
		fmt.Println(err)
	}
	if err := fo.moveFile("/home/user/messy_folder/photo.jpg", "Images"); err != nil {
		fmt.Print(err)
	}
}
