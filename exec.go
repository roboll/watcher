package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/roboll/watcher/handlers"
)

var (
	execCommand = os.Getenv("EXEC_COMMAND")
)

func init() {
	handlers.Register("exec", Exec)
}

func Exec(path string) error {
	log.Printf("exec: %s", path)
	args := strings.Split(execCommand, " ")

	switch len(args) {
	case 0:
		return errors.New("exec: EXEC_COMMAND is required")
	case 1:
		out, err := exec.Command(args[0]).CombinedOutput()
		log.Println(string(out))
		return err
	default:
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		log.Println(string(out))
		return err
	}
}
