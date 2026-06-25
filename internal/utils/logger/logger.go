package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type DailyWriter struct {
	mu          sync.Mutex
	currentDate string
	file        *os.File
}

func NewDailyWriter() (*DailyWriter, error) {
	w := &DailyWriter{}

	if err := w.rotateIfNeeded(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *DailyWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.rotateIfNeeded(); err != nil {
		return 0, err
	}

	return w.file.Write(p)
}

func (w *DailyWriter) rotateIfNeeded() error {
	now := time.Now()

	date := now.Format("2006-01-02")

	if date == w.currentDate && w.file != nil {
		return nil
	}

	if w.file != nil {
		_ = w.file.Close()
	}

	logDir := filepath.Join(
		"logs",
		now.Format("2006"),
		now.Format("01"),
		now.Format("02"),
	)

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(
		filepath.Join(logDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		return err
	}

	w.file = file
	w.currentDate = date

	return nil
}

func (w *DailyWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		return w.file.Close()
	}

	return nil
}

func Setup() (*DailyWriter, error) {
	dailyWriter, err := NewDailyWriter()
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(
		os.Stdout,
		dailyWriter,
	)

	log.SetOutput(multiWriter)

	gin.DefaultWriter = multiWriter
	gin.DefaultErrorWriter = multiWriter

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return dailyWriter, nil
}
