package main

import (
	"bufio"
	"fmt"
	"strings"
)

func insertionSort(data *arrComments, count int, field string, asc bool) {
	if count <= 1 {
		return
	}

	for i := 1; i < count; i++ {
		key := data[i]
		j := i - 1
		
		shifting := true
		for j >= 0 && shifting {
			compareResult := 0

			switch field {
				case "id":
					if key.id < data[j].id {
						compareResult = -1
					} else if key.id > data[j].id {
						compareResult = 1
					}
				case "sender":
					keySender := strings.ToLower(key.sender)
					dataJSender := strings.ToLower(data[j].sender)
					if keySender < dataJSender {
						compareResult = -1
					} else if keySender > dataJSender {
						compareResult = 1
					}
				default:
					fmt.Println("Sorting Error due to Unknown Field.")
					return
			}

			if asc {
				if compareResult == -1 {
					data[j+1] = data[j]
					j--
				} else {
					shifting = false
				}
			} else { // Descending
				if compareResult == 1 {
					data[j+1] = data[j]
					j--
				} else {
					shifting = false
				}
			}
		}
		data[j+1] = key
	}
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

func statusToString(status int) string {
	switch status {
	case 1:
		return "Positive"
	case -1:
		return "Negative"
	default:
		return "Neutral"
	}
}

func getStringInput(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func getFirstNWords(text string, n int) string {
	words := strings.Fields(text)
	if len(words) <= n {
		return text
	}

	return strings.Join(words[:n], " ")
}