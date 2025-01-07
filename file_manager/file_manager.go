package file_manager

import (
	"bufio"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (fm *FileManager) WriteToFile(data string) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(data)

	if err != nil {
		return err
	}

	return nil
}

func New(inputFilePath string, outputFilePath string) *FileManager {
	fm := &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}

	return fm
}
