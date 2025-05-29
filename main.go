package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type comment struct {
	id int
	sender string
	comments string
	status int
}

const NMAX = 100000

type arrComments [NMAX]comment

func main() {
	var data arrComments

	data = arrComments{
		{id: 1, sender: "Gus", comments: "Yo: Gurt", status: 0},
		{id: 2, sender: "Gus", comments: "Mul: Yo\nYo: No 😂✌️", status: 0},
		{id: 3, sender: "Gus", comments: "Gurt: Yo", status: 0},
		{id: 4, sender: "Gus", comments: "Gurt: Yo", status: 0},
		{id: 5, sender: "Gus", comments: "Gurt: Yo", status: 0},
		{id: 6, sender: "Gus", comments: "Gurt: Yo", status: 0},
		{id: 7, sender: "Gus", comments: "Gurt: Yo", status: 0},
		{id: 8, sender: "Gus", comments: "Gurt: Yo", status: 0},
		{id: 9, sender: "Gus", comments: "Gurt: Yo", status: 0},
	}

	fmt.Println(data)
}

func mainMenu() {
	var input int
	// var (
	// 	comments arrComments
	// 	commentsCount = 0
	// )

	// var (
	// 	originalOrder arrComments
	// 	originalOrderCount = 0
	// )

	// var nextCommentID = 1

	fmt.Println("\nMain Menu")
	fmt.Println("--------------------")
	fmt.Println("1. Manage Komentar")
	fmt.Println("2. Perlihatkan Statistik Komentar")
	fmt.Println("3. Review Komentar")
	fmt.Println("4. Exit")
	fmt.Println("--------------------")
	fmt.Print("Masukkan pilihan: ")

	for {
		fmt.Scan(&input)
		switch input {
		case 1:
			fmt.Println("\nMenuju ke Manage Komentar~")
			// manageCommentsMenu(&comments, &commentsCount, &originalOrder, &originalOrderCount, &nextCommentID)
		case 2:
			fmt.Println("\nMenuju ke Perlihatkan Statistik Komentar~")
			// showStatistics()
		case 3:
			fmt.Println("\nMenuju ke Review Comments~")
			// reviewComments()
		case 4:
			fmt.Println("\nExiting program. . . Sampai Bertemu Lagi, Atmin :D")
			os.Exit(0) // Exits the program
		default:
			fmt.Println("\nInput Tidak Valid. Tolong pilih opsi valid (1-4).")
		}
		fmt.Println()
	}
}

func manageCommentsMenu(commentsArr *arrComments, originalOrderData *arrComments, nextID *int) {
	for {
		fmt.Println("\nManage Comments")
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
		case 2:
		case 3:
		case 4:
		case 5:
			return
		default:
			fmt.Println("Input Tidak Valid. Tolong pilih opsi valid (1-5).")
		}
		fmt.Println()
	}
}

func displayComments(commentsList *arrComments, count int) {
	if len(commentsList) == 0 {
		fmt.Println("Tidak ada komentar.")
		return
	}

	fmt.Println("\n--- List Komentar ---")
	for i := 0; i < count; i++{
		c := commentsList[i]
		fmt.Printf("ID: %d, Sender: %s, Status: %d\nComment: %s\n------------------------\n",
			c.id, c.sender, c.status, c.comments)
	}
}

func searchComments(sourceArr *arrComments, sourceCount int, r *bufio.Reader, resultsDisplayArr *arrComments, resultsDisplayCount *int) {
	var input int
	*resultsDisplayCount = 0

	if sourceCount == 0 {
		fmt.Println("Tidak ada komentar.")
		return
	}

	fmt.Println("\nSearch By: ")
	fmt.Println("1. Sender")
	fmt.Println("2. Comment Text")
	fmt.Println("3. Status (Positif, Negatif, Netral)")
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scanln(&input)

	
	for i := 0; i < sourceCount; i++ {
		c := sourceArr[i]
		match := false
		switch input {
		case 1:
			searchTerm := getStringInput("Masukkan kata kunci: ", r)
			if strings.Contains(strings.ToLower(c.sender), strings.ToLower(searchTerm)) {
				match = true
			}
		case 2:
			searchTerm := getStringInput("Masukkan kata kunci: ", r)
			if strings.Contains(strings.ToLower(c.comments), strings.ToLower(searchTerm)) {
				match = true
			}
		case 3:
			var targetStatus int
			validStatusSearch := true
			fmt.Println("\nPilih status komentar: ")
			fmt.Println("1. Positif")
			fmt.Println("2. Negatif")
			fmt.Println("3. Netral")
			fmt.Print("Masukkan Pilihan: ")
			fmt.Scanln(&targetStatus)
			
			if targetStatus != 1 && targetStatus != 2 && targetStatus != 3 {
				fmt.Println("Input tidak valid.")
				validStatusSearch = false
				return
			}

			if validStatusSearch && c.status == targetStatus {
				match = true
			}
		default:
			fmt.Println("Input tidak valid.")
			return
		}

		if match && *resultsDisplayCount < NMAX {
			resultsDisplayArr[*resultsDisplayCount] = c
			(*resultsDisplayCount)++
		}
	}

	if *resultsDisplayCount > 0  {
		fmt.Println("\n--- Hasil Pencarian ---")
		for i := 0; i < *resultsDisplayCount; i++ {
			res := resultsDisplayArr[i]
			fmt.Printf("ID: %d, Sender: %s, Status: %s\nComment: %s\n------------------------\n",
				res.id, res.sender, statusToString(res.status), res.comments)
		}
	} else {
		fmt.Println("Tidak ada komentar yang ditemukan.")
	}
}

// func commentFilters(data arrComments) int {
// 	var (
// 		indexPositive float64 = 3/3
// 		indexNeutral float64 = 2/3
// 		indexNegative float64 = 1/3
// 	)

// 	var res int



// 	return res
// }

