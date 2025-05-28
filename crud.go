package main

import (
	"bufio"
	"fmt"
)

// Create, Read, Update, Delete (CRUD)

func readCommentsSubMenu(commentsArr *arrComments, commentsCount *int, originalOrderArr *arrComments, originalOrderCount *int, reader *bufio.Reader) {
	var (
		displayArr arrComments
		displayCount int
	)

	var (
		searchResultArr arrComments
		searchResultCount int
	)

	for {
		fmt.Println("\nRead Comments Sub-Menu")
		fmt.Println("--------------------")
		fmt.Println("1. Display All Comments (Current Order of main list)")
		fmt.Println("2. Sort & Display by ID (Ascending) - Selection Sort")
		fmt.Println("3. Sort & Display by ID (Descending) - Selection Sort")
		fmt.Println("4. Sort & Display by Sender (Ascending) - Selection Sort")
		fmt.Println("5. Sort & Display by Sender (Descending) - Selection Sort")
		fmt.Println("6. Reset Main List to Default Order")
		fmt.Println("7. Search Comments (from current main list order)")
		fmt.Println("8. Back to Manage Comments Menu")
		fmt.Println("--------------------")
		
		var input int
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&input)

		if input >= 2 && input <= 5 {
			if *commentsCount > 0 {
				copy(displayArr[:*commentsCount], commentsArr[:*commentsCount])
				displayCount = *commentsCount
			} else {
				displayCount = 0
			}
		}

		switch input {
		case 1:
			fmt.Println("\n--- Menampilkan Main Comment List ---")
			displayComments(commentsArr, *commentsCount)
		case 2: // Sort By ID Ascending
			if displayCount > 0 {
				insertionSort(&displayArr, displayCount, "id", true)
				fmt.Println("Komentar yang diurutkan berdasarkan ID (Ascending): ")
				displayComments(&displayArr, displayCount)
			} else {
				fmt.Println("Tidak ada komentar untuk diurutkan.")
			}
		case 3: // Sort By ID Ascending
			if displayCount > 0 {
				insertionSort(&displayArr, displayCount, "id", false)
				fmt.Println("Komentar yang diurutkan berdasarkan ID (Descending): ")
				displayComments(&displayArr, displayCount)
			} else {
				fmt.Println("Tidak ada komentar untuk diurutkan.")
			}
		case 4: // Sort By Sender Ascending
			if displayCount > 0 {
				insertionSort(&displayArr, displayCount, "sender", true)
				fmt.Println("Komentar yang diurutkan berdasarkan ID (Ascending): ")
				displayComments(&displayArr, displayCount)
			} else {
				fmt.Println("Tidak ada komentar untuk diurutkan.")
			}
		case 5: // Sort By Sender Descending
			if displayCount > 0 {
				insertionSort(&displayArr, displayCount, "sender", false)
				fmt.Println("Komentar yang diurutkan berdasarkan ID (Descending): ")
				displayComments(&displayArr, displayCount)
			} else {
				fmt.Println("Tidak ada komentar untuk diurutkan.")
			}
		case 6: // Reset List to Default Order
			if *originalOrderCount > 0 {
				countToCopy := *originalOrderCount
				copy(commentsArr[:countToCopy], originalOrderArr[:countToCopy])
				*commentsCount = countToCopy
				
				fmt.Println("Main comment list telah di-reset ke urutan default")
			} else {
				fmt.Println("Tidak ada komentar untuk diurutkan.")
			}
		case 7: // Search comments
			if *commentsCount > 0 {
				searchComments(commentsArr, *commentsCount, reader, &searchResultArr, &searchResultCount)
			} else {
				fmt.Println("Tidak ada komentar.")
			}
		case 8:
			return
		default:
			fmt.Println("Input tidak valid.")
		}
	}
}