package helper

import "log"

func PanicIfError(err error, msg string) {
	if err != nil {
		log.Panicf("Error: %s. Message: %s", err, msg)
	}
}
