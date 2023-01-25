package logger

import (
	"log"
	"strings"
)

const (
	ERROR   = "[Error]: "
	SUCCESS = "[Success]: "
)

func logger(typeLog string) {
	log.SetFlags(0)
	log.SetPrefix(typeLog)
}

func LogError(msg ...string) {
	logger(ERROR)
	log.Fatalln(strings.Join(msg[:], " "))
}

func LogSuccess(msg ...string) {
	logger(SUCCESS)
	log.Println(strings.Join(msg[:], " "))
}
