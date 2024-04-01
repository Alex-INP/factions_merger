package main

import (
	"regexp"
)

type Parser struct {
	raw_data  string
	json_data map[string]string
}

func (p *Parser) parseData() {
	re := regexp.MustCompile(`".+"=`)
	var nameIndexes = re.FindAllIndex([]byte(p.raw_data), -1)
	var contentIndexes [][2]int

	var pointerFirst int
	var pointerSecond int = 1
	for pointerSecond < len(nameIndexes) {
		itemFirst, itemSecond := nameIndexes[pointerFirst], nameIndexes[pointerSecond]
		contentIndexes = append(contentIndexes, [2]int{itemFirst[1], itemSecond[0] - 1})
		pointerFirst += 1
		pointerSecond += 1
	}
	contentIndexes = append(
		contentIndexes,
		[2]int{
			nameIndexes[len(nameIndexes) - 1][1],
			len(p.raw_data),
		},
	)

	for ind, indexes := range nameIndexes {
		factionName := p.raw_data[indexes[0]: indexes[1]]
		factionContent := p.raw_data[contentIndexes[ind][0]: contentIndexes[ind][1]]
		p.json_data[factionName] = factionContent
	}
}

func GetParser(raw_data string) Parser {
	return Parser{raw_data, make(map[string]string)}
}
