package logger

import (
	"log"
	"os"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	Log          *log.Logger
	currentLevel = LevelInfo
)

// SetLevel changes the minimum log level.
func SetLevel(level Level) {
	currentLevel = level
}

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

}

func Debug(v ...any) {

	if currentLevel > LevelDebug {
		return
	}

	Log.Println(
		append([]any{"DEBUG:"}, v...)...,
	)
}

func Info(v ...any) {

	if currentLevel > LevelInfo {
		return
	}

	Log.Println(
		append([]any{"INFO:"}, v...)...,
	)
}

func Warn(v ...any) {

	if currentLevel > LevelWarn {
		return
	}

	Log.Println(
		append([]any{"WARN:"}, v...)...,
	)
}

func Error(v ...any) {

	if currentLevel > LevelError {
		return
	}

	Log.Println(
		append([]any{"ERROR:"}, v...)...,
	)
}

func Separator(title string) {

	Info(
		"========================================",
	)

	Info(title)

	Info(
		"========================================",
	)
}

func Section(title string) {

	Info("----------", title, "----------")
}
