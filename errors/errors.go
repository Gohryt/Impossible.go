package errors

import (
	"log"
	"os"
)

func Panic(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Log(err error) {
	if err != nil {
		log.Println(err)
	}
}
