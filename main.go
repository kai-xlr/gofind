package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dirFlag := flag.String("dir", ".", "Directory path to operate in (default is current directory)")

	flag.Parse()
	fmt.Println("Directory to operate in:", *dirFlag)

	if _, err := os.Stat(*dirFlag); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", *dirFlag)
	} else {
		fmt.Printf("Directory exists: %s\n", *dirFlag)
	}
}
