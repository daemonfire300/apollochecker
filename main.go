package main

import (
	"fmt"
	"log"
	"os"

	"github.com/daemonfire300/apollochecker/pkg/client"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Not enough command line arguments")
		return
	}
	infoText, err := client.Status(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	fmt.Println(infoText)
}
