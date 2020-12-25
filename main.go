package main

import "log"

func main() {
	cmd := NewBcryptGenerateCmd()
	if err := cmd.Execute(nil); err != nil {
		log.Fatalln(err)
	}
}
