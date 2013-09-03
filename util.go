package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func execCmd(input string) {
	inputs := strings.Split(input, " ")
	name := inputs[0]
	args := inputs[1:]

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
