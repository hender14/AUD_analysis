package main

import (
	"app/infrastructure"
	"log"
)

// main関数
func main() {
	fs, err := infrastructure.NewDB()
	if err != nil {
		log.Fatalf("Listen and serve failed. %s\n", err)
	}
	r := infrastructure.NewRouting(fs)
	r.Run()
}