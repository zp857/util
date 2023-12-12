package db

import (
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Writer struct {
	logger.Writer
	logZap bool
}

func NewWriter(w logger.Writer, logZap bool) *Writer {
	return &Writer{
		Writer: w,
		logZap: logZap,
	}
}

func (w *Writer) Printf(message string, data ...interface{}) {
	if w.logZap {
		zap.L().Sugar().Named("[gorm]").Infof(message, data...)
	} else {
		w.Writer.Printf(message, data...)
	}
}
