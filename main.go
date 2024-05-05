package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MERGE_FILE_NAME = "user_empire_designs_v3.4.txt"
const AFFINITY_FILE_NAME = "affinity.txt"
const MERGE_FILE_A_PATH = "merge_a"
const MERGE_FILE_B_PATH = "merge_b"
const RESULT_FILE_PATH = "result"

const AFFINITY_DELIMITER_A = "--A--"
const AFFINITY_DELIMITER_B = "--B--"

func main() {
	var filesManager = getFilesManager()
	filesManager.setupFilesAndFolders()

	var parserA = parseDataFromFile(filesManager.filePathA)
	var parserB = parseDataFromFile(filesManager.filePathB)
	mergedMap := mergeFactions(parserA, parserB, filesManager)

	file, err := os.Create(filesManager.resultFilePath)
	if err != nil {
		fmt.Println(err)
		pressEnterAndExit()
	}
	defer file.Close()
	for factionName, factionData := range mergedMap {
		file.WriteString(factionName)
		file.WriteString(factionData)
	}

	fmt.Println("Мердж успешно завершен")
	pressEnterAndExit()
}

func parseDataFromFile(filePath string) Parser {
	raw_data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		pressEnterAndExit()
	}
	var parser Parser = GetParser(string(raw_data))
	parser.parseData()

	return parser
}

func mergeFactions(parserA Parser, parserB Parser, filesManager FilesManager) map[string]string {
	file, err := os.OpenFile(filesManager.affinityFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		pressEnterAndExit()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const stateA string = "A"
	const stateB string = "B"
	var affinityDataA []string
	var affinityDataB []string
	var delimiterState string
	for scanner.Scan() {
		var text string = scanner.Text()
		switch text {
		case AFFINITY_DELIMITER_A:
			delimiterState = stateA
		case AFFINITY_DELIMITER_B:
			delimiterState = stateB
		default:
			if text != "" {
				if delimiterState == stateA {
					affinityDataA = append(affinityDataA, text)
				} else if delimiterState == stateB {
					affinityDataB = append(affinityDataB, text)
				}
			}
		}
	}

	result := make(map[string]string)
	addAffiliateData(affinityDataA, parserA.json_data, result)
	addAffiliateData(affinityDataB, parserB.json_data, result)

	return result
}

func addAffiliateData(affinityData []string, parsedData map[string]string, result map[string]string) {
	for factionName, factionData := range parsedData {
		for _, n := range affinityData {
			if strings.Contains(factionName, n) {
				result[factionName] = factionData
				continue
			}
		}
	}
}

func pressEnterAndExit() {
	fmt.Println()
	fmt.Println("Press Enter to exit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	os.Exit(0)
}
