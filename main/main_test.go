package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	arr := [][]string{
		{"hassane 10... (cap)\\nohamo", "hen.txt"},
		{"\"yassine", "yassine.txt"},
		{"https://learn.zone01oujda.ma/", "zone.txt"},
		{"ok", "ok.txt"},
		{"gorila", "out.txt"},
		{"hello", "hello.txt"},
		{"HELLO", "HELLO.txt"},
		{"HeLlo HuMaN", "HeLloHuMaN.txt"},
		{"1Hello 2There", "1Hello2There.txt"},
		{"Hello\\nThere", "hellothere.txt"},
		{"{Hello & There #}", "1234.txt"},
		{"hello There 1 to 2!", "helloThere1to2.txt"},
		{"MaD3IrA&LiSboN", "MaD3IrA&LiSboN.txt"},
		{"RGB", "RGB.txt"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ABCDEFGHIJKLMNOQPRSTUVWXYZ.txt"},
	}
	for _, test := range arr {
		PassOrNo(t, test[0], test[1])
	}
}

func PassOrNo(t *testing.T, test, filename string) {
	cmd := exec.Command("go", "run", ".", test)
	out, err := cmd.Output()
	// fmt.Print(string(out))
	if err != nil {
		// fmt.Println("errr 1")
		return
	}
	file, err := os.ReadFile("../.testfile/" + filename)
	 //fmt.Println(string(file))
	if err != nil {
		// fmt.Println("errr 2")
		return
	}
	if string(out) == string(file) {
		fmt.Println("test " + test + " --------------------> pass ")
	}
	if string(out) != string(file) {
		fmt.Println("test " + test + "-------------------> No pass ")
	}
}
