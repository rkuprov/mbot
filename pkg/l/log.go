package l

import (
	"log"
	"os"
	"sync"
)

var logger MLogger

type MLogger struct {
	slog *log.Logger
	flog *log.Logger
}

func NewLogger() func() error {
	logger.slog = newSlog()
	flog, done := newFlog()
	logger.flog = flog

	return done
}

func newSlog() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}
func newFlog() (*log.Logger, func() error) {
	f, err := os.OpenFile("mbot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	flogger := log.New(f, "", log.LstdFlags|log.Lshortfile)

	return flogger, f.Close
}

func Log(s string) {
	m := sync.Mutex{}
	m.Lock()
	logger.slog.Println(s)
	logger.flog.Println(s)
	m.Unlock()
}
