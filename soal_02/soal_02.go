package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("\nSoal nomor 2 : \nResult =>  %+v", FormatWord("telkom", "telecom"))
	fmt.Printf("\nSoal nomor 2 : \nResult =>  %+v", FormatWord("telkom", "tlkom"))
}

// Function Soal Nomor 02
func FormatWord(w1, w2 string) bool {
	var res []string
	var temp string

	if len(w2) < len(w1) {
		temp = w2
		w2 = w1
		w1 = temp
	} else if len(w1) == len(w2) {
		r, _ := regexp.Compile("[" + w1 + "]+")
		res = r.FindAllString(w2, 2)
		toStr := strings.Join(res, "")
		var result []int
		var t int

		for i := 0; i < len(toStr); i++ {
			if string(toStr[i]) != " " {
				t++
				result = append(result, t)
			} else if i == len(toStr)-1 {
				result = append(result, t)
			}
		}

		if len(result) == len(w2)-1 {
			return true
		}
		return false
	}

	r, _ := regexp.Compile("[" + w2 + "]+")
	res = r.FindAllString(w1, -1)
	toStr := strings.Join(res, "")

	if len(toStr) == len(w1) {
		return true
	} else {
		return false
	}
}
