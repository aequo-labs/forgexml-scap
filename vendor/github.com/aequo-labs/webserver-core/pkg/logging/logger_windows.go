//go:build windows
// +build windows

package logging

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc/eventlog"
)

type windowsLogger struct {
	eventLog *eventlog.Log
}

func (w *windowsLogger) Close() error {
	return w.eventLog.Close()
}

func (w *windowsLogger) Log(level LogLevel, message string) error {
	switch level {
	case DEBUG, INFO:
		return w.eventLog.Info(1, message)
	case WARN:
		return w.eventLog.Warning(2, message)
	case ERROR, FATAL:
		return w.eventLog.Error(3, message)
	default:
		return fmt.Errorf("unknown log level: %v", level)
	}
}

func eventSourceExists(source string) bool {
	key := fmt.Sprintf("SYSTEM\\CurrentControlSet\\Services\\EventLog\\Application\\%s", source)
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	return true
}

func newSysLogger(source string) (SysLogger, error) {
	// Check if the event source already exists
	if !eventSourceExists(source) {
		// Attempt to create the event source only if it doesn't exist
		err := eventlog.InstallAsEventCreate(source, eventlog.Error|eventlog.Warning|eventlog.Info)
		if err != nil {
			return nil, fmt.Errorf("failed to install event source: %v", err)
		}
	}

	// Open the event log
	el, err := eventlog.Open(source)
	if err != nil {
		return nil, fmt.Errorf("failed to open event log: %v", err)
	}

	return &windowsLogger{eventLog: el}, nil
}
