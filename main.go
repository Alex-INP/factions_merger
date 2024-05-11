package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MERGE_FILE_NAME = "user_empire_designs_v3.4.txt"
const AFFINITY_FILE_NAME = "affinity.txt"
const MERGE_FILE_A_DIR = "merge_a"
const MERGE_FILE_B_DIR = "merge_b"
const RESULT_FILE_DIR = "result"

const AFFINITY_DELIMITER_A = "--A--"
const AFFINITY_DELIMITER_B = "--B--"

func main() {
	var filesManager = getFilesManager()
	filesManager.setupFilesAndFolders()

	var parserA = parseDataFromFile(filesManager.filePathA)
	var parserB = parseDataFromFile(filesManager.filePathB)

	resultA := make(map[string]string)
	resultB := make(map[string]string)
	affinityDataA, affinityDataB := getAffinityData(filesManager)
	addAffiliateData(affinityDataA, parserA.json_data, resultA)
	addAffiliateData(affinityDataB, parserB.json_data, resultB)

	writeResultFile([2]map[string]string{resultA, resultB}, filesManager)

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

func getAffinityData(filesManager FilesManager) ([]string, []string) {
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
	return affinityDataA, affinityDataB
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

func writeResultFile(resultData [2]map[string]string, filesManager FilesManager) {
	file, err := os.Create(filesManager.resultFilePath)
	if err != nil {
		fmt.Println(err)
		pressEnterAndExit()
	}
	defer file.Close()
	for _, item := range resultData {
		for factionName, factionData := range item {
			file.WriteString(factionName)
			file.WriteString(factionData)
		}
	}
}

func pressEnterAndExit() {
	fmt.Println()
	fmt.Println("Нажмите Enter чтобы выйти")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	os.Exit(0)
}
