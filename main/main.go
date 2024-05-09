package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	// "regexp"
	"art"
)

func main() {
	arr := os.Args[1:]

	// str,_ := regexp.MustCompile(`--n=([^ ]+) "([^"]+)" ([^\s]+)`)

	nFlag := flag.String("output", "", "help message for flag output")

	flag.Parse()
	if len(arr) != 3 {
		println("please entre one word ")
		return
	} else if len(arr[1]) == 0 {
		return
	}
	art.Check(arr[1])
	artDraw, err := art.Art("standard")
	if err != nil {
		fmt.Println(err)
		return
	}
	split_backN := strings.Split(arr[1], "\\n")
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
	art.Write(resulat, &nFlag)
}
