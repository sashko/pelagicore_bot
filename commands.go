package main

import (
	"log"
	"os/exec"
)

func handleHelp() string {
	str := "I accept the following commands:\n\n" +
		"/help\t    print available commands"

	return str
}

func handleScreenLock() string {
	cmd := exec.Command("gnome-screensaver-command", "--lock")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)

		return err.Error()
	}

	return "Screen has been LOCKED"
}

func handleScreenUnlock() string {
	cmd := exec.Command("gnome-screensaver-command", "--deactivate")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)

		return err.Error()
	}

	return "Screen has been UNLOCKED"
}
