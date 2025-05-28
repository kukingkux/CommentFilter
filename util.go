package main

import (
	"bufio"
	"fmt"
	"strings"
)

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

func getStringInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func insertionSort(data *arrComments, count int, field string, ascending bool) {
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

			if ascending {
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

func binarySearch(arr arrComments, target int) (int, int) {
	iterations := 0
	low := 0
	high := len(arr) - 1

	for low <= high {
		iterations++
		mid := (low + high) / 2

		if arr[mid].id == target {
			return mid, iterations
		} else if target < arr[mid].id {
			high = mid -1
		} else {
			low = mid + 1
		}
	}
	return -1, iterations
}