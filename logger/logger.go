package logger

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.New(log.Writer(), "", log.LstdFlags)}
}

func (l *Logger) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		l.Printf("\033[44m[%s]\033[0m %s %s - %v", r.Method, r.URL.Path, duration, http.StatusOK)
	})
}

func (l *Logger) Info(status int, path string, start time.Time) {
	duration := time.Since(start)
	l.Printf("\033[42m[%d]\033[0m %s - %v", status, path, duration)
}

func (l *Logger) Error(status int, path string, err error) {
	l.Printf("\033[41m[%d]\033[0m | PATH: \033[33m%s\033[0m | ERROR: %v", status, path, err)
}
