package main

import (
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

