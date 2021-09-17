package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func validate(s string) (bool, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return false, errors.New("Harus terdapat String")
	}
	return true, nil
}
func main() {
	var input string
	var yt string
	var x = true
	for x {
		fmt.Print("Inputkan sebuah kata: ")
		fmt.Scanln(&input)
		if valid, err := validate(input); valid {
			reg, err := regexp.Compile("[^a-zA-Z]+")

			if err != nil {
				log.Fatal(err)
			}

			newInput := reg.ReplaceAllString(input, "")

			fmt.Println(reverse(newInput))
		} else {
			fmt.Println(err.Error())
		}
		fmt.Print("\nLanjut memesan? (Y/T): ")
		fmt.Scanln(&yt)

		yt = strings.ToUpper(yt)
		if yt == "T" {
			os.Exit(0)
		}
	}

}
