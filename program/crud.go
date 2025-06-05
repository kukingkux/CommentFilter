package main

import (
	"bufio"
	"fmt"
)

// Create, Read, Update, Delete (CRUD)

func readCommentsSubMenu(commentsArr *arrComments, commentsCount *int, originalOrderArr *arrComments, originalOrderCount *int, r *bufio.Reader) {
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
				
				fmt.Println("\nMain comment list telah di-reset ke urutan default")
			} else {
				fmt.Println("Tidak ada komentar untuk diurutkan.")
			}
		case 7: // Search comments
			if *commentsCount > 0 {
				searchComments(commentsArr, *commentsCount, r, &searchResultArr, &searchResultCount)
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

func displayComments(commentsList *arrComments, count int) {
	if len(commentsList) == 0 {
		fmt.Println("Tidak ada komentar.")
		return
	}

	fmt.Println("\n--- List Komentar ---")
	for i := 0; i < count; i++{
		c := commentsList[i]
		fmt.Printf("ID: %d, Sender: %s, Status: %s\nComment: %s\n------------------------\n",
			c.id, c.sender, statusToString(c.status), c.text)
	}
}

func createComment(commentsArr *arrComments, commentsCount *int, originalOrderArr *arrComments, originalOrderCount *int, nextID *int, r *bufio.Reader) {
	fmt.Println("\n--- Create New Comment ---")

	sender := getStringInput("Masukkan nama sender: ", r)
	commentText := getStringInput("Masukkan komentar: ", r)

	newComment := comment{
		id: *nextID,
		sender: sender,
		text: commentText,
		status: 0,
		isHidden: false,
	}

	commentsArr[*commentsCount] = newComment
	(*commentsCount)++

	originalOrderArr[*originalOrderCount] = newComment
	(*originalOrderCount)++

	(*nextID)++
	fmt.Println("Komentar berhasil ditambahkan!")
	fmt.Printf("ID: %d, Sender: %s, Comment: %s, Status: %s\n",
		newComment.id, newComment.sender, newComment.text, statusToString(newComment.status))
}

func findCommentByID(id int, arr *arrComments, count int) (int, *comment) {
	for i := 0; i < count; i++ {
		if arr[i].id == id {
			return i, &arr[i]
		}
	}
	return -1, nil
}

func editComment(commentsArr *arrComments, commentsCount *int, originalOrderArr *arrComments, originalOrderCount *int, r *bufio.Reader) {
	fmt.Println("\n --- Edit Comment ---")

	if *commentsCount == 0 {
		fmt.Println("Tidak ada komentar untuk diedit.")
		return
	}

	var targetID int
	fmt.Print("Masukkan ID komentar yang akan diedit: ")
	fmt.Scanln(&targetID)

	_, targetComment := findCommentByID(targetID, commentsArr, *commentsCount)
	if targetComment == nil {
		fmt.Printf("Komentar dengan ID %d tidak ditemukan.\n", targetID)
		return
	}

	fmt.Printf("Mengedit Comment ID: %d (Sender: %s, Status: %s)\n", targetComment.id, targetComment.sender, statusToString(targetComment.status))
	fmt.Printf("Komentar saat ini: %s\n", targetComment.text)

	newSender := getStringInput(fmt.Sprintf("Masukkan sender baru (current: %s, press Enter to keep): ", targetComment.sender), r)
	newCommentText := getStringInput(fmt.Sprintf("Masukkan komentar baru (current: %s, press Enter to keep): ", targetComment.text), r)

	changed := false
	if newSender != "" {
		targetComment.sender = newSender
		changed = true
	}
	if newCommentText != "" {
		targetComment.text = newCommentText
		changed = true
	}

	if changed {
		indexInOriginal, originalComment := findCommentByID(targetID, originalOrderArr, *originalOrderCount)
		if originalComment != nil {
			originalOrderArr[indexInOriginal] = *targetComment
		}
		fmt.Printf("Komentar dengan ID %d telah berhasil di-update.\n", targetID)
	} else {
		fmt.Println("Tidak ada perubahan yang dibuat.")
	}
}

func deleteComment(commentsArr *arrComments, commentsCount *int, originalOrderArr *arrComments, originalOrderCount *int) {
	fmt.Println("\n --- Delete Comment ---")
	
	if *commentsCount == 0 {
		fmt.Println("Tidak ada komentar untuk dihapus.")
	}

	var targetID int
	fmt.Print("Masukkan ID komentar yang akan dihapus: ")
	fmt.Scanln(&targetID)

	indexInComments, commentFound := findCommentByID(targetID, originalOrderArr, *originalOrderCount)
	if commentFound == nil {
		fmt.Printf("Komentar dengan ID %d tidak ditemukan.\n", targetID)
		return
	}

	for i := indexInComments; i < *commentsCount-1; i++ {
		commentsArr[*commentsCount-1] = comment{}
	}
	if *commentsCount > 0 {
		commentsArr[*commentsCount-1] = comment{}
	}
	(*commentsCount)--

	indexInOriginal, originalCommentFound := findCommentByID(targetID, originalOrderArr, *originalOrderCount)
	if originalCommentFound != nil {
		for i := indexInOriginal; i < *originalOrderCount-1; i++ {
			originalOrderArr[i] = originalOrderArr[i+1]
		}
		if *originalOrderCount > 0 {
			originalOrderArr[*originalOrderCount-1] = comment{}
		}
		(*originalOrderCount)--
	}
	fmt.Printf("Komentar dengan ID %d telah berhasil dihapus.", targetID)
}