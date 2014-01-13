package main

import (
	"fmt"
	"github.com/yosssi/gocmd"
	"os/exec"
	"strings"
)

const (
	msgNothingFormat = "\nThere is nothing to format.\n"
)

func main() {
	output, err := gocmd.Pipe(exec.Command("find", "."), exec.Command("grep", "\\.go"))
	if err != nil {
		fmt.Println(msgNothingFormat)
		return
	}
	files := convOutputToStrings(output)
	targets := []string{}
	for _, file := range files {
		output, err := exec.Command("gofmt", "-l", file).Output()
		if err != nil {
			continue
		}
		target := convOutputToStrings(output)
		if len(target) == 0 {
			continue
		}
		targets = append(targets, target[0])
	}
	if len(targets) == 0 {
		fmt.Println(msgNothingFormat)
		return
	}
	fmt.Println("\nThe following go files have to be formatted:\n")
	for _, target := range targets {
		fmt.Println(target)
	}
	fmt.Println("")
	panic("There are the go files to be formatted.")
}

func convOutputToStrings(output []byte) []string {
	s := strings.Split(string(output), "\n")
	s = s[:len(s)-1]
	return s
}
