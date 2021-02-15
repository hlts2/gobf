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
	)
	if err != nil {
		log.Fatal(err)
	}

	err = bf.Do(context.Background(), func(pattern string) {
		if pattern == "s3sx" {
			log.Printf("find: %s\n", pattern)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
