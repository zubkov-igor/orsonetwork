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
	DebugLog     *log.Logger
	currentLevel = LevelInfo
)

// SetLevel changes the minimum log level.
func SetLevel(level Level) {
	currentLevel = level
}

func Init() {

	file, err := os.OpenFile(
		"info.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		panic(err)
	}

	debugFile, err := os.OpenFile(
		"debug.info",
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

	DebugLog = log.New(
		debugFile,
		"",
		log.Ldate|log.Ltime,
	)
}

func Debug(v ...any) {

	DebugLog.Println(
		append([]any{"DEBUG:"}, v...)...,
	)
}

func Info(v ...any) {

	if currentLevel <= LevelInfo {

		Log.Println(
			append([]any{"INFO:"}, v...)...,
		)
	}

	DebugLog.Println(
		append([]any{"INFO:"}, v...)...,
	)
}

func Warn(v ...any) {

	if currentLevel <= LevelWarn {

		Log.Println(
			append([]any{"WARN:"}, v...)...,
		)
	}

	DebugLog.Println(
		append([]any{"WARN:"}, v...)...,
	)
}

func Error(v ...any) {

	if currentLevel <= LevelError {

		Log.Println(
			append([]any{"ERROR:"}, v...)...,
		)
	}

	DebugLog.Println(
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
