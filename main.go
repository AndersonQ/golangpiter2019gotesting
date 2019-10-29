package main

import (
	"log"
	"os/exec"

	"github.com/AndersonQ/golangpiter2019gotesting/pwd"
)

func main() {
	cmd := exec.Command("pwd")
	out, _ := cmd.Output()
	log.Printf("main pwd: %s", string(out))
	pwd.PWD()
}
