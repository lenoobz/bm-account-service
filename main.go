// Package main entry point of bm account service
package main

import (
	"log"

	"github.com/letrong/bm-account-service/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
