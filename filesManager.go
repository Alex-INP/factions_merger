package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FilesManager struct {
	dirPathA         string
	dirPathB         string
	resultDirPath    string
	filePathA        string
	filePathB        string
	resultFilePath   string
	affinityFilePath string
}

func getFilesManager() FilesManager {
	var cwd, err = os.Getwd()
	handleErrorIfAny(err)

	var result FilesManager = FilesManager{
		filepath.Join(cwd, MERGE_FILE_A_DIR),
		filepath.Join(cwd, MERGE_FILE_B_DIR),
		filepath.Join(cwd, RESULT_FILE_DIR),
		"",
		"",
		"",
		filepath.Join(cwd, AFFINITY_FILE_NAME),
	}
	result.filePathA = filepath.Join(result.dirPathA, MERGE_FILE_NAME)
	result.filePathB = filepath.Join(result.dirPathB, MERGE_FILE_NAME)
	result.resultFilePath = filepath.Join(result.resultDirPath, MERGE_FILE_NAME)

	return result
}

func (fm *FilesManager) setupFilesAndFolders() {
	var createdAny bool = false

	_, err := os.Stat(fm.affinityFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(fm.affinityFilePath)
		handleErrorIfAny(err)
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s\n\n%s", AFFINITY_DELIMITER_A, AFFINITY_DELIMITER_B))

		createdAny = true
	}

	for _, folderPath := range [3]string{fm.dirPathA, fm.dirPathB, fm.resultDirPath} {
		_, err := os.Stat(folderPath)
		if os.IsNotExist(err) {
			err = os.Mkdir(folderPath, os.ModePerm)
			handleErrorIfAny(err)
			createdAny = true
		}
	}
	if createdAny {
		fmt.Println("Необходимые папки и файлы созданы")
		pressEnterAndExit()
	}
}
