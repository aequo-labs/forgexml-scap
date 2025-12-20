//go:build darwin
// +build darwin

package logging

import (
	"log/syslog"
)

type darwinLogger struct {
	syslog *syslog.Writer
}

func (d *darwinLogger) Close() error {
	return d.syslog.Close()
}

func (d *darwinLogger) Log(level LogLevel, message string) error {
	switch level {
	case DEBUG:
		return d.syslog.Debug(message)
	case INFO:
		return d.syslog.Info(message)
	case WARN:
		return d.syslog.Warning(message)
	case ERROR:
		return d.syslog.Err(message)
	case FATAL:
		return d.syslog.Crit(message)
	default:
		return d.syslog.Info(message)
	}
}

func newSysLogger(tag string) (SysLogger, error) {
	// On macOS, syslog messages are typically sent to /var/log/system.log
	// The LOG_USER facility is commonly used for user-level applications
	syslogWriter, err := syslog.New(syslog.LOG_INFO|syslog.LOG_USER, tag)
	if err != nil {
		return nil, err
	}
	return &darwinLogger{syslog: syslogWriter}, nil
}
