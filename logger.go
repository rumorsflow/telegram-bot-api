package tgbotapi

import (
	"fmt"
	"go.uber.org/zap"
)

type logger struct {
	l *zap.Logger
}

func (l *logger) Println(v ...any) {
	if len(v) > 0 {
		switch m := v[0].(type) {
		case error:
			l.l.Error("telegram bot api error", zap.Error(m))
		case string:
			l.l.Info(m)
		}
	}
}

func (l *logger) Printf(format string, v ...any) {
	l.l.Info(fmt.Sprintf(format, v...))
}
