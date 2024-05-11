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
	// if len(arr) <= 3 {
	// 	println("please entre one word ")
	// 	return
	// } else if len(arr[1]) == 0 {
	// 	return
	// }
	split_backN := []string{}
	file:= ""
	if len(arr) == 1 {
		art.Check(arr[0])
		split_backN = strings.Split(arr[0], "\\n")
		file="standard"
	} else if len(arr) == 3 {
		art.Check(arr[1])
		split_backN = strings.Split(arr[1], "\\n")
		file=arr[2]
	} else {
		fmt.Println("sir t9Awad")
	}
	// var artDraw map[int][]string
	// if len(arr) == 1 {
	artDraw,_:= art.Art(file)
	// }


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
	if len(arr) == 1 {
		fmt.Println(resulat)
	} else if len(arr) == 3 {
		art.Write(resulat, &nFlag)
	}
	
}
