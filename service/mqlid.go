package service

import (
	"strings"
)

const (
	characters = "ABCDEF0123456789"
	idLength   = 8
	totalIDs   = 4294967296 // 16^8
	batchSize  = 4000
)

func generateBatchIDs(startIndex int, chars []rune, length int, batchSize int) []string {
	result := make([]string, 0, batchSize)
	currentIndex := 0
	var generate func(string, int)
	generate = func(prefix string, remainingLength int) {
		if len(result) >= batchSize {
			return
		}
		if remainingLength == 0 {
			if currentIndex >= startIndex {
				result = append(result, prefix)
			}
			currentIndex++
			return
		}
		for _, char := range chars {
			if len(result) >= batchSize {
				return
			}
			generate(prefix+string(char), remainingLength-1)
		}
	}

	generate("", length)
	return result
}

func SaveMQLID() {
	chars := []rune(characters)
	totalGenerated := 0

	for totalGenerated < totalIDs {
		ids := generateBatchIDs(totalGenerated, chars, idLength, batchSize)
		if len(ids) == 0 {
			break
		}
		SetMongoDB(strings.Join(ids, ","), "")
		totalGenerated += len(ids)
	}
}
