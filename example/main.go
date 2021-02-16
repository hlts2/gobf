package main

import (
	"context"
	"log"

	"github.com/hlts2/gobf"
)

func main() {
	bf, err := gobf.New(
		gobf.WithNumber(true),
		gobf.WithUpper(true),
		gobf.WithLower(true),
		gobf.WithSize(4),
		gobf.WithConcrencyLimit(1000000),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start to search pattern: s3Df")

	err = bf.Do(context.Background(), func(pattern string) {
		if pattern == "s3Df" {
			log.Printf("find: %s\n", pattern)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
