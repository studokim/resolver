package internal

import "log"

func HandleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Handle(err error) {
	if err != nil {
		log.Println(err)
	}
}
