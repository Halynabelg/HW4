package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myfile, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)

	}
}
