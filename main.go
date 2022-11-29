package main

import (
	"fmt"
	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	//var dataset = [5]string{"128kb", "64mb", "128mb", "256mb", "512mb"}                                                                                                                                                                                                                     //, "1024kb", "2048kb"}
	var query_dataset = [5]string{"QmZFbKpxLdSBMTjvoqwknYn2bmbcRVynZX2CyfRPhqjHtL", "QmUpWVqNF2fh639AVe2s3b58SuuBkW39qtMgLDDU43Rt6s", "QmZFbKpxLdSBMTjvoqwknYn2bmbcRVynZX2CyfRPhqjHtL", "QmYKGy9y4kAteLU1sLUCnRMemfi4gsH2ea1Vdn9Y8Wr9UN", "QmVjuwPYXCfcaRWCCrXViSRDpvFEq8UqXqfqZGJDAhZRnk"} //, "1024kb", "2048kb"}

	sh := shell.NewShell("54.176.172.176:5001")
	time.Sleep(5 * time.Second)

	/*for _, s := range dataset {

		fmt.Println(s)
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
		time.Sleep(5 * time.Second)
	}*/

	for _, s := range query_dataset {

		fmt.Println(s)
		startTime := time.Now()

		err := sh.Get(s, s+".txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}

		elapsedTime := time.Since(startTime)

		fmt.Printf("실행시간: %s\n", elapsedTime)
		time.Sleep(5 * time.Second)
	}
}
