package main

import (
	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strconv"
)

func main() {

	keyboard := keylogger.FindKeyboardDevice()

	logrus.Println("Found a keyboard at", keyboard)

	k, err := keylogger.New(keyboard)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer k.Close()

	poder := 1800
	poders := " "
	events := k.Read()
	for even := range events {
		switch even.Type {

		case keylogger.EvKey:

			if even.KeyString() == "F1" {
				if poder >= 200 {
					poder = poder - 200
				}
				poders = strconv.Itoa(poder)
				exec.Command("/bin/sh", "-c", "echo "+poders+" > /sys/class/backlight/intel_backlight/brightness").CombinedOutput()
			}
			if even.KeyString() == "F2" {
				if poder < 1800 {
					poder = poder + 200
				}
				poders = strconv.Itoa(poder)
				exec.Command("/bin/sh", "-c", "echo "+poders+" > /sys/class/backlight/intel_backlight/brightness").CombinedOutput()

			}

			break
		}
	}
}
