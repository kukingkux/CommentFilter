package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Create, Read, Update, Delete (CRUD)

func readCommentsSubMenu(commentsArr *arrComments, commentsCount *int, r *bufio.Reader) {
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
		fmt.Println("6. Search Comments (from current main list order)")
		fmt.Println("7. Back to Manage Comments Menu")
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
		case 6: // Search comments
			if *commentsCount > 0 {
				searchComments(commentsArr, *commentsCount, r, &searchResultArr, &searchResultCount)
			} else {
				fmt.Println("Tidak ada komentar.")
			}
		case 7:
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
		if !c.isHidden {
			fmt.Printf("ID: %d, Sender: %s, Status: %s\nComment: %s\n------------------------\n",
			c.id, c.sender, statusToString(c.status), c.text)
		}
		
	}
}

func createComment(commentsArr *arrComments, commentsCount *int, nextID *int, r *bufio.Reader) {
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

func editComment(commentsArr *arrComments, commentsCount *int, r *bufio.Reader) {
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

	if newSender != "" {
		targetComment.sender = newSender
	}
	if newCommentText != "" {
		targetComment.text = newCommentText
	}
}

func deleteComment(commentsArr *arrComments, commentsCount *int) {
	fmt.Println("\n --- Delete Comment ---")
	
	if *commentsCount == 0 {
		fmt.Println("Tidak ada komentar untuk dihapus.")
	}

	var targetID int
	fmt.Print("Masukkan ID komentar yang akan dihapus: ")
	fmt.Scanln(&targetID)

	indexInComments, commentFound := findCommentByID(targetID, commentsArr, *commentsCount)
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

	fmt.Printf("Komentar dengan ID %d telah berhasil dihapus.", targetID)
}

func searchComments(sourceArr *arrComments, sourceCount int, r *bufio.Reader, resultsDisplayArr *arrComments, resultsDisplayCount *int) {
	var (
		choiceSearch int
		targetStatus int
		choiceStatus int
		searchTerm string
	)
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
	_, err := fmt.Scanln(&choiceSearch)
	if err != nil {
		fmt.Println("Input invalid.")
		return
	}
	
	validStatusSearch := true
	
	switch choiceSearch {
	case 1:
		searchTerm = getStringInput("Masukkan nama sender: ", r)
	case 2:
		searchTerm = getStringInput("Masukkan kata kunci: ", r)
	case 3:
		fmt.Println("\nPilih status komentar: ")
		fmt.Println("1. Positif")
		fmt.Println("2. Negatif")
		fmt.Println("3. Netral")
		fmt.Print("Masukkan Pilihan (1-3): ")
		_, err := fmt.Scanln(&choiceStatus)
		if err != nil {
			fmt.Println("Input tidak valid.")
			return
		}
		
		switch choiceStatus {
		case 1:
			targetStatus = 1
		case 2:
			targetStatus = -1
		case 3:
			targetStatus = 0
		default:
			fmt.Println("Input tidak valid.")
			validStatusSearch = false
		}

		if !validStatusSearch {
			return
		}
	default:
		fmt.Println("Input tidak valid.")
		return
	}

	resultsFull := false

	for i := 0; i < sourceCount; i++ {
		c := sourceArr[i]
		match := false

		switch choiceSearch {
		case 1:
			if strings.Contains(strings.ToLower(c.sender), strings.ToLower(searchTerm)) {
				match = true
			}
		case 2:
			if strings.Contains(strings.ToLower(c.text), strings.ToLower(searchTerm)) {
				match = true
			}
		case 3:
			if validStatusSearch && c.status == targetStatus {
				match = true
			}
		}
		
		if match {
			if !resultsFull {
				if *resultsDisplayCount < NMAX {
					resultsDisplayArr[*resultsDisplayCount] = c
					(*resultsDisplayCount)++
				} else {
					fmt.Println("Batas hasil pencarian tercapai. Tidak semua yang cocok dapat ditampilkan.")
					resultsFull = true
				}
			}
		}
	}

	if *resultsDisplayCount > 0  {
		fmt.Println("\n\033[1m--- Hasil Pencarian ---\033[0m")
		for i := 0; i < *resultsDisplayCount; i++ {
			res := resultsDisplayArr[i]
			fmt.Printf("ID: %d, Sender: %s, Status: %s\nComment: %s\n------------------------\n",
				res.id, res.sender, statusToString(res.status), res.text)
		}
	} else {
		fmt.Println("Tidak ada komentar yang ditemukan.")
	}
}

func setVisibility(commentsArr *arrComments, commentsCount int) {
	for {
		fmt.Println("\n--- Atur Visibilitas Komentar (Tampilkan/Sembunyikan) ---")

		fmt.Println("Daftar Komentar:")
		for i := 0; i < commentsCount; i++ {
			c := commentsArr[i]
			if c.status == -1 {
				fmt.Printf("ID: %d,Sender: %s, Status: %s, Visibility: %s\nComment: %s\n------------------------\n",
				c.id, c.sender, statusToString(c.status), isHiddenToString(c.isHidden), c.text)
			}
		}

		var toggleID int
		fmt.Print("Masukkan ID komentar yang visibilitasnya ingin diubah (Enter untuk skip): ")
		_, err := fmt.Scanln(&toggleID)
		if err != nil {
			return
		}

		_, toggleComment := findCommentByID(toggleID, commentsArr, commentsCount)
		if toggleComment == nil {
			fmt.Printf("Komentar dengan ID %d tidak ditemukan.\n", toggleID)
			return
		}

		currentVisibility := "Show"
		if toggleComment.isHidden {
			currentVisibility = "Hidden"
		}
		fmt.Printf("Komentar ID %d (%s) saat ini %s.\n", toggleComment.id, getFirstNWords(toggleComment.text, 5), currentVisibility)
		fmt.Println("Pilih aksi:")
		fmt.Println("1. Tampilkan Komentar (Set Visible)")
		fmt.Println("2. Sembunyikan Komentar (Set Hidden)")
		fmt.Println("3. Batal")
		fmt.Print("Masukkan pilihan: ")

		var choice int
		_, err = fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Input tidak valid.")
			return
		}

		switch choice {
		case 1:
			if toggleComment.isHidden {
				toggleComment.isHidden = false
				fmt.Printf("Komentar ID %d sekarang ditampilkan.\n", toggleComment.id)
			} else  {
				fmt.Printf("Komentar ID %d sudah ditampilkan.\n", toggleComment.id)
			}
		case 2:
			if !toggleComment.isHidden {
				toggleComment.isHidden = true
				fmt.Printf("Komentar ID %d sekarang disembunyikan.\n", toggleComment.id)
				
			} else  {
				fmt.Printf("Komentar ID %d sudah disembunyikan.\n", toggleComment.id)
			}
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
			return
		}
	}
}