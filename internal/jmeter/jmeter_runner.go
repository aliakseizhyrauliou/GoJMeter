package jmeter

import (
    "os"
    "os/exec"
)


func RunJMeterTest(userCount string) error {
	cmd := exec.Command("jmeter", "-n", "-t", userCount+"_users.jmx")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}