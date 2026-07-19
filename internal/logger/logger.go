package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {

	file, err := os.OpenFile(
		"orsonetwork.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		panic(err)
	}

	Log = log.New(
		file,
		"",
		log.Ldate|log.Ltime,
	)

	Log.Println(
		"LOGGER STARTED",
	)
}