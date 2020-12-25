package main

import (
	"log"
	"os"

	"github.com/btwiuse/bcrypt/pkg/bcrypt"
)

func main() {
	if err := bcrypt.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
