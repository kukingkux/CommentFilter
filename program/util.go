package main

import (
	"bufio"
	"fmt"
	"strings"
)

// insertionSort() mengurutkan array komentar secara ascending atau descending
// dapat menyorting berdasarkan id atau sender
// membutuhkan array komentar, jumlah data, field/properti untuk di-sort, dan boolean untuk ascending (true) atau descending (false)
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

// statusToString() mengubah status (int) menjadi string ("positive"/"negative"/"neutral")
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

// isHiddenToString() mengubah isHidden (bool) menjadi string ("Hidden"/"Visible")
func isHiddenToString(isHidden bool) string {
	switch isHidden {
	case true:
		return "Hidden"
	case false:
		return "Visible"
	default:
		return "Not Valid."
	}
}

// getStringInput() menampilkan prompt kepada user yang membaca seluruh baris teks
// termasuk spasi dan return berupa string yang telah dipangkas
func getStringInput(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

// getFirstNWords mengembalikan string pertama sejumlah N
func getFirstNWords(text string, n int) string {
	words := strings.Fields(text)
	if len(words) <= n {
		return text
	}

	return strings.Join(words[:n], " ")
}