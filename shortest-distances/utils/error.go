package utils

import "log"

func IsFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
