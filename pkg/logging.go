package pkg

import (
    "log"
    "os"
)

// представляет нашу структуру логгера
type Logger struct {
    *log.Logger
}

// создает и возвращает новый экземпляр Logger
func NewLogger() *Logger {
    return &Logger{
        Logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

// логирует сообщения с уровнем Debug
func (l *Logger) Debug(msg string) {
    l.SetPrefix("DEBUG: ")
    l.Println(msg)
}

// логирует сообщения с уровнем Info
func (l *Logger) Info(msg string) {
    l.SetPrefix("INFO: ")
    l.Println(msg)
}

// логирует сообщения с уровнем Error
func (l *Logger) Error(msg string) {
    l.SetPrefix("ERROR: ")
    l.Println(msg)
}
