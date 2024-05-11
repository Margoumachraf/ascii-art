package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"art"
)

func main() {
	arr := os.Args[1:]

	

	nFlag := flag.String("output", "", "help message for flag output")

	flag.Parse()
	split_backN := []string{}
	file := ""
	if len(arr) == 1 {
		art.Check(arr[0])
		split_backN = strings.Split(arr[0], "\\n")
		file = "shadow"
	} else if len(arr) == 2 {
		art.Check(arr[0])
		split_backN = strings.Split(arr[0], "\\n")
		file = arr[1]
		fmt.Println(file)
	} else if len(arr) == 3 {
		art.Check(arr[1])
		split_backN = strings.Split(arr[1], "\\n")
		file = arr[2]
	} else {
		fmt.Println("sir t9Awad")
	}

	artDraw, _ := art.Art(file)

	if art.Back(split_backN) {
		for i := 0; i < len(split_backN)-1; i++ {
			fmt.Println()
		}
		return
	}
	var resulat string
	for _, word := range split_backN {
		resulat += art.PrintArt(word, artDraw)
	}
	if len(arr) == 3 {
		art.Write(resulat, &nFlag)
	} else {
		fmt.Println(resulat)
	}
}
