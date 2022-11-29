package main

import (
	"fmt"
	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	var dataset = [6]string{"64kb", "128kb", "256kb", "512kb", "1024kb", "2048kb"}
	sh := shell.NewShell("URL:5001")
	for _, s := range dataset {
		startTime := time.Now()

		file, err := os.Open("dataset/" + s)
		cid, err := sh.Add(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		fmt.Println("added ", cid)
		elapsedTime := time.Since(startTime)

		fmt.Printf("실행시간: %s\n", elapsedTime)
	}
}
