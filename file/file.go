package file

import (
	"log"
	"os"
)

func Setup() {
	src := "./storage"
	_, err := os.Stat(src)

	if os.IsNotExist(err) {
		err := os.MkdirAll(src, os.ModePerm)
		if err != nil {
			log.Fatalf("file setup create file err: %v", err)
		}
	}
}
