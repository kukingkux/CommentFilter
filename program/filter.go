package main

import (
	"fmt"
	"strings"
)

func analyzeCommentSentiment(
	commentText string,
	posKeywords *arrPKeywords, posKeywordsCount int,
	negKeywords *arrNKeywords, negKeywordsCount int,
	negationWords *arrNegationWords, negationWordsCount int,
	intensifiers *arrIntensifierWords, intensifierCount int,
	diminishers *arrDiminisherWords, diminisherWordCount int,
) (newStatus int, sentimenScore float64) {

	words := strings.Fields(strings.ToLower(commentText)) // Memisahkan tiap kata dan membuatnya jadi huruf kecil
	totalSentimenScore := 0.0
	n := len(words)

	for i := 0; i < n; i++ {
		currentWord := words[i]
		wordScore := 0.0
		isSentimenWord := false

		for k := 0; k < posKeywordsCount; k++ {
			if currentWord == posKeywords[k].word {
				wordScore = float64(posKeywords[k].score)
				isSentimenWord = true
				break
			}
		}

		if !isSentimenWord {
			for k := 0; k < negKeywordsCount; k++ {
				if currentWord == negKeywords[k].word {
					wordScore = float64(negKeywords[k].score)
					isSentimenWord = true
					break
				}
			}
		}

		if isSentimenWord && i > 0 {
			previousWord := words[i-1]
			modifierApplied := false

			for k := 0; k < negationWordsCount; k++ {
				if previousWord == negationWords[k] {
					wordScore *= -1.0
					modifierApplied = true
					break
				}
			}

			if !modifierApplied {
				for k := 0; k < intensifierCount; k++ {
					if previousWord == intensifiers[k].word {
						wordScore *= intensifiers[k].multiplier
						modifierApplied = true
						break
					}
				}
			}

			if !modifierApplied {
				for k := 0; k < diminisherWordCount; k++ {
					if previousWord == diminishers[k].word {
						wordScore *= diminishers[k].multiplier
						modifierApplied = true
						break
					}
				}
			}
		}
		totalSentimenScore += wordScore
	}

	finalStatus := 0
	if totalSentimenScore > 0.5 {
		finalStatus = 1
	} else if totalSentimenScore < -0.5 {
		finalStatus = -1
	}

	return finalStatus, totalSentimenScore
}

func reviewComments(
	commentsArr *arrComments, commentsCount *int,
	originalOrderArr *arrComments, originalOrderCount *int,
	posKeywords *arrPKeywords, posKeywordsCount int,
	negKeywords *arrNKeywords, negKeywordsCount int,
	negationWords *arrNegationWords, negationWordsCount int,
	intensifiers *arrIntensifierWords, intensifierCount int,
	diminishers *arrDiminisherWords, diminisherWordsCount int,
) {
	fmt.Println("\n--- Starting Comment Review Process ---")
	analyzedCount := 0
	changedCount := 0

	for i := 0; i < *commentsCount; i++ {
		analyzedCount++
		originalStatus := commentsArr[i].status
		newStatus, score := analyzeCommentSentiment(commentsArr[i].text,
			posKeywords, posKeywordsCount,
			negKeywords, negKeywordsCount,
			negationWords, negationWordsCount,
			intensifiers, intensifierCount,
			diminishers, diminisherWordsCount,
		)

		if newStatus != originalStatus {
			commentsArr[i].status = newStatus
			changedCount++
			fmt.Printf("Comment ID %d: Text=\"%s...\", Old Status: %s, New Status: %s, Score: %.2f\n",
				commentsArr[i].id, getFirstNWords(commentsArr[i].text, 5), statusToString(originalStatus),
				statusToString(newStatus), score)

			idxInOriginal, originalComment := findCommentByID(commentsArr[i].id, originalOrderArr, *originalOrderCount)
			if originalComment != nil {
				originalOrderArr[idxInOriginal].status = newStatus
			}
		}
	}
	fmt.Printf("--- Review Process Complete. Analyzed: %d neutral comments. Status changed for: %d comments. --\n",
	analyzedCount, changedCount)
}

