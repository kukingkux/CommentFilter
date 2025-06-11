package main

import (
	"bufio"
	"fmt"
	"os"
)

type arrComments [NMAX]comment
type arrPKeywords [MAX_P_KEYWORDS]keywordScore
type arrNKeywords [MAX_N_KEYWORDS]keywordScore
type arrNegationWords [MAX_NEGATION_WORDS]string
type arrIntensifierWords [MAX_INTENSIFIER_WORDS]modifierWord
type arrDiminisherWords [MAX_DIMINISHER_WORDS]modifierWord

func main() {
	var (
		dataComments arrComments
		dataPKeywords arrPKeywords
		dataNKeywords arrNKeywords
		dataNegationWords arrNegationWords
		dataIntensifierWords arrIntensifierWords
		dataDiminisherWords arrDiminisherWords
	)

	var (
		commentsCount int
		posKeywordCount int
		negKeywordCount int
		negationWordCount int
		intensifierWordCount int
		diminisherWordCount int
	)

	var nextCommentID int = 10

	dataComments = arrComments{
		{id: 2, sender: "Fufu", text: "Mul: Yo\nYo: No 😂✌️", status: 0},
		{id: 5, sender: "Mulyo", text: "Saya senang dengan fufu namun saya kecewa dengan fafa", status: 0},
		{id: 3, sender: "Akira", text: "Gurt, lukisanmu keren banget", status: 0},
		{id: 6, sender: "Akira", text: "Yo, ini jelek banget woi", status: 0},
		{id: 1, sender: "Gus", text: "Yo: Gurt, lu keren parah tapi ga nonchalant", status: 0},
		{id: 8, sender: "Ye", text: "kelakuan Kanye sangat buruk", status: 0},
		{id: 4, sender: "Gus", text: "Yang dia lakukan sangat pintar tapi juga berbahaya", status: 0},
		{id: 9, sender: "Kanye", text: "Hebat bisa bikin beat se keren ini", status: 0},
		{id: 7, sender: "Kanye", text: "NYOKAP GUA DIRUMAH MASAK PEPES TAHU 🗣️ 🗣️ 🔥🔥🔥🔥🔥🔥", status: 0},
	}
	commentsCount = 9

	dataPKeywords = arrPKeywords{
		{word: "bagus", score: 3},
		{word: "baik", score: 2},
		{word: "suka", score: 4},
		{word: "senang", score: 4},
		{word: "bahagia", score: 5},
		{word: "hebat", score: 4},
		{word: "mantap", score: 3},
		{word: "keren", score: 3},
		{word: "pintar", score: 3},
	}
	posKeywordCount = 9

	dataNKeywords = arrNKeywords{
		{word: "buruk", score: -3},
		{word: "jelek", score: -3},
		{word: "benci", score: -4},
		{word: "kecewa", score: -4},
		{word: "parah", score: -5},
		{word: "menyedihkan", score: -4},
		{word: "masalah", score: -2},
		{word: "tidak suka", score: -4},
		{word: "berbahaya", score: -3},
	}
	negKeywordCount = 9

	dataNegationWords = arrNegationWords{
		"tidak",
		"bukan",
		"jangan",
		"tak",
		"ngga",
		"engga",
		"ga",
	}
	negationWordCount = 7

	dataIntensifierWords = arrIntensifierWords{
		{word: "sangat", multiplier: 1.5},
		{word: "amat", multiplier: 1.5},
		{word: "benar-benar", multiplier: 1.7},
		{word: "sekali", multiplier: 1.3},
		{word: "banget", multiplier: 1.3},
	}
	intensifierWordCount = 5

	dataDiminisherWords = arrDiminisherWords{
		{word: "agak", multiplier: 0.7},
		{word: "rada", multiplier: 0.7},
		{word: "sedikit", multiplier: 0.6},
		{word: "kurang", multiplier: 0.5},
	}
	diminisherWordCount = 4

	for {
		fmt.Println("\n====================")
		fmt.Println("Main Menu")
		fmt.Println("--------------------")
		fmt.Println("1. Manage Komentar")
		fmt.Println("2. Perlihatkan Statistik Komentar")
		fmt.Println("3. Review Komentar")
		fmt.Println("4. Exit")
		fmt.Println("--------------------")
		fmt.Print("Masukkan pilihan: ")

		input := 0
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("\nMenuju ke Manage Komentar~")
			manageCommentsMenu(&dataComments, &commentsCount, &nextCommentID)
		case 2:
			fmt.Println("\nMenuju ke Perlihatkan Statistik Komentar~")
			showStatistics(&dataComments, commentsCount)
		case 3:
			fmt.Println("\nMenuju ke Review Comments~")
			reviewComments(&dataComments, &commentsCount, &dataPKeywords, posKeywordCount,&dataNKeywords, 
				negKeywordCount,&dataNegationWords, negationWordCount,&dataIntensifierWords, intensifierWordCount, &dataDiminisherWords, diminisherWordCount)
		case 4:
			fmt.Println("\nExiting program. . . Sampai Bertemu Lagi, Atmin :D")
			os.Exit(0) // Exits the program
		default:
			fmt.Println("\nInput Tidak Valid. Tolong pilih opsi valid (1-4).")
		}
		fmt.Println()
	}
}

// manageCommentsMenu() digunakan untuk menampilkan sub-menu manage komentar
// akan terus looping hingga user memilih kembali ke main menu
func manageCommentsMenu(commentsArr *arrComments, commentsCount *int, nextID *int) {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n====================")
		fmt.Println("Manage Comments")
		fmt.Println("--------------------")
		fmt.Println("1. Read Comments (List, Sort, Search)")
		fmt.Println("2. Create Comment")
		fmt.Println("3. Edit Comment (by ID)")
		fmt.Println("4. Delete Comment (by ID)")
		fmt.Println("5. Kembali ke Main Menu")
		fmt.Println("--------------------")

		var input int
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			readCommentsSubMenu(commentsArr, commentsCount, r)
		case 2:
			createComment(commentsArr, commentsCount, nextID, r)
		case 3:
			editComment(commentsArr, commentsCount, r)
		case 4:
			deleteComment(commentsArr, commentsCount)
		case 5:
			return
		default:
			fmt.Println("Input Tidak Valid. Tolong pilih opsi valid (1-5).")
		}
		fmt.Println()
	}
}

// showStatistics() menghitung dan menampilkan statistik komentar
// menampilkan jumlah komentar dan statusnya serta
// persentase per-kategori status (positif, negatif, netral)
func showStatistics(commentsArr *arrComments, commentsCount int) {
	fmt.Println("\n--- Statistik Komentar---")

	if commentsCount == 0 {
		fmt.Println("Tidak ada komentar.")
		return
	}

	posCount := 0
	negCount := 0
	neuCount := 0

	for i := 0; i < commentsCount; i++ {
		switch commentsArr[i].status {
		case 1:
			posCount++
		case -1:
			negCount++
		case 0:
			neuCount++
		}
	}

	total := float64(commentsCount)
	percentPos := 0.0
	percentNeg := 0.0
	percentNeu := 0.0

	if total > 0 {
		percentPos = (float64(posCount) / total) * 100
		percentNeg = (float64(negCount) / total) * 100
		percentNeu = (float64(neuCount) / total) * 100
	}

	fmt.Printf("Total Komentar Keseluruhan: %d\n", commentsCount)
	fmt.Println("----------------------------------------")
	fmt.Printf("Jumlah Komentar Positif  : %d (%.2f%%)\n", posCount, percentPos)
	fmt.Printf("Jumlah Komentar Negatif  : %d (%.2f%%)\n", negCount, percentNeg)
	fmt.Printf("Jumlah Komentar Netral   : %d (%.2f%%)\n", neuCount, percentNeu)
	fmt.Println("----------------------------------------")
}