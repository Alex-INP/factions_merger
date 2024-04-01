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

func getFileManager() FilesManager {
	var cwd, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		pressEnterAndExit()
	}

	var result FilesManager = FilesManager{
		filepath.Join(cwd, MERGE_FILE_A_PATH),
		filepath.Join(cwd, MERGE_FILE_B_PATH),
		filepath.Join(cwd, RESULT_FILE_PATH),
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
		if err != nil {
			fmt.Println(err)
			pressEnterAndExit()
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s\n\n%s", AFFINITY_DELIMITER_A, AFFINITY_DELIMITER_B))

		createdAny = true
	}

	for _, folderPath := range [3]string{fm.dirPathA, fm.dirPathB, fm.resultDirPath} {
		_, err := os.Stat(folderPath)
		if os.IsNotExist(err) {
			err = os.Mkdir(folderPath, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				pressEnterAndExit()
			}
			createdAny = true
		}
	}
	if createdAny {
		fmt.Println("Необходимые папки и файлы созданы")
		pressEnterAndExit()
	}
}
