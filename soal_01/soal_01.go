package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	MataUang("145.000")
	// MataUang("2050")
}

// Function Soal  Nomor 01
func MataUang(uang string) {
	var res int
	var pecahan = [10]int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	rep := strings.Replace(uang, ".", "", -1)

	//Pembulatan pecahan keatas
	if rep[len(rep)-2:] != "00" {
		r := strings.Replace(rep[len(rep)-2:], rep[len(rep)-2:], "00", -1)
		bulat := rep[:len(rep)-2] + r
		v, _ := strconv.Atoi(bulat)
		res = v + 100
	} else {
		v, _ := strconv.Atoi(rep)
		res = v
	}

	//Tukar uang
	var temp, x int
	temp = res
	var toStr string

	for i := 0; i < len(pecahan); i++ {
		x = temp / pecahan[i]
		temp = temp - (pecahan[i] * x)

		if pecahan[i] >= 1000 {
			if x != 0 {
				toStr = strconv.Itoa(pecahan[i])
				fmt.Printf("\nRp. %+v.000 : %+v", toStr[:len(toStr)-3], x)
			}

		} else if pecahan[i] < 1000 {
			if x != 0 {
				toStr := strconv.Itoa(pecahan[i])
				fmt.Printf("\nRp. %+v : %+v", toStr, x)
			}
		}
	}
}
