package main

import (
	"fmt"
	"strings"
)

// analyzeCommentSentiment() menghitung skor sentimen berdasarkan string teks.
// memecah input teks, kemudian cek bobot kata positif dan negatif
// dan mengatur skor berdasarkan kata tambahan negation, intensifier, atau diminisher
// return status yg telah ditentukan sentimennya (1 untuk positif dan -1 untuk negatif)
// dan skor dari sentimen
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

		for k := 0; k < posKeywordsCount && !isSentimenWord; k++ {
			if currentWord == posKeywords[k].word {
				wordScore = float64(posKeywords[k].score)
				isSentimenWord = true
			}
		}

		if !isSentimenWord {
			for k := 0; k < negKeywordsCount && !isSentimenWord; k++ {
				if currentWord == negKeywords[k].word {
					wordScore = float64(negKeywords[k].score)
					isSentimenWord = true
				}
			}
		}

		if isSentimenWord && i > 0 {
			previousWord := words[i-1]
			modifierApplied := false

			for k := 0; k < negationWordsCount && !modifierApplied; k++ {
				if previousWord == negationWords[k] {
					wordScore *= -1.0
					modifierApplied = true
				}
			}

			if !modifierApplied {
				for k := 0; k < intensifierCount && !modifierApplied; k++ {
					if previousWord == intensifiers[k].word {
						wordScore *= intensifiers[k].multiplier
						modifierApplied = true
					}
				}
			}

			if !modifierApplied {
				var diminisherFound bool = false
				for k := 0; k < diminisherWordCount && !diminisherFound; k++ {
					if previousWord == diminishers[k].word {
						wordScore *= diminishers[k].multiplier
						diminisherFound = true
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

// reviewComments() cek komentar secara otomatis
// cek keseluruhan array satu-per-satu mencari komentar dengan status netral
// kemudian memanggil analyzeCommentSentimen() untuk menentukan nilai status yang baru
// setelah review otomatis, user secara manual dapat menentukan komentar negatif untuk disembunyikan
func reviewComments(
	commentsArr *arrComments, commentsCount *int,
	posKeywords *arrPKeywords, posKeywordsCount int,
	negKeywords *arrNKeywords, negKeywordsCount int,
	negationWords *arrNegationWords, negationWordsCount int,
	intensifiers *arrIntensifierWords, intensifierCount int,
	diminishers *arrDiminisherWords, diminisherWordsCount int,
) {
	fmt.Println("\n=== Starting Comment Review Process ===")
	fmt.Println("\n--- List Komentar ---")
	fmt.Printf("\033[1m%-4s %-12s %-12s %-10s\033[0m\n", "ID", "Old Status", "New Status", "Score")
	fmt.Println("-----------------------")
	analyzedCount := 0
	changedCount := 0

	insertionSort(commentsArr, *commentsCount, "id", true)
	for i := 0; i < *commentsCount; i++ {
		if commentsArr[i].status == 0 {
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
				fmt.Printf("%-4d %-12s %-12s %-10.2f \nComment: %s...\n------------------------\n",
					commentsArr[i].id, statusToString(originalStatus),
					statusToString(newStatus), score, getFirstNWords(commentsArr[i].text, 5))
			}
		}
	}

	setVisibility(commentsArr, *commentsCount)	
	fmt.Printf("--- Review Process Complete. Analyzed: %d neutral comments. Status changed for: %d comments. --\n",
	analyzedCount, changedCount)
}

