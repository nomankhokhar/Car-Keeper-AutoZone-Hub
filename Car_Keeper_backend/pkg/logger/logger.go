package logger

import "log"

func Info(message string) {
	log.Println("[INFO]", message)
}

func Error(message string) {
	log.Println("[ERROR]", message)
}

func Debug(message string) {
	log.Println("[DEBUG]", message)
}

func Fatal(message string) {
	log.Fatal("[FATAL]", message)
}
func Success(message string) {
	log.Println("[SUCCESS]", message)
}
