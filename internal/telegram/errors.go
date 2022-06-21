package telegram

import "log"

func logError(err error) {
	log.Print(err)
	panic(err)
}
