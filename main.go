package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/omar391/sofinder/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <so-file-names-in-comma-sep-format> [architecture: x86_64|arm64] [distro: ubuntu|arch] [output-directory] [remove-container]")
		os.Exit(1)
	}

	soFileNames := os.Args[1]
	arch := "x86_64"
	if len(os.Args) > 2 {
		arch = os.Args[2]
	}
	distro := "ubuntu"
	if len(os.Args) > 3 {
		distro = os.Args[3]
	}
	outputDir := "./so_files"
	if len(os.Args) > 4 {
		outputDir = os.Args[4]
	}
	remove := false
	if len(os.Args) > 5 {
		remove = os.Args[5] == "true"
	}

	sg, err := pkg.NewSoFinder(arch, distro, outputDir, "", remove)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = sg.Collect(strings.Split(soFileNames, ",")...)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Completed successfully.")
}
