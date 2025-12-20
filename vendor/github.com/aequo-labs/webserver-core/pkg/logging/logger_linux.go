//go:build linux
// +build linux

package logging

import (
	"log/syslog"
)

type linuxLogger struct {
	syslog *syslog.Writer
}

func (l *linuxLogger) Close() error {
	return l.syslog.Close()
}

func (l *linuxLogger) Log(level LogLevel, message string) error {
	switch level {
	case DEBUG:
		return l.syslog.Debug(message)
	case INFO:
		return l.syslog.Info(message)
	case WARN:
		return l.syslog.Warning(message)
	case ERROR:
		return l.syslog.Err(message)
	case FATAL:
		return l.syslog.Crit(message)
	default:
		return l.syslog.Info(message)
	}
}

func newSysLogger(tag string) (SysLogger, error) {
	syslogWriter, err := syslog.New(syslog.LOG_INFO|syslog.LOG_USER, tag)
	if err != nil {
		return nil, err
	}
	return &linuxLogger{syslog: syslogWriter}, nil
}
