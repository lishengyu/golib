package slog

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

const (
	LogLevelInfo  = "INFO"
	LogLevelWarn  = "WARN"
	LogLevelError = "ERROR"
)

var syslogger *syslog.Writer

func init() {
	InitSyslog("PcapReplay")
}

// InitSyslog 初始化系统日志
func InitSyslog(appName string) error {
	var err error
	syslogger, err := syslog.New(syslog.LOG_NOTICE|syslog.LOG_DAEMON, appName)
	if err != nil {
		return err
	}

	multiWriter := io.MultiWriter(os.Stdout, syslogger)
	log.SetOutput(multiWriter)
	return nil
}

func Err(message string) {
	log.Printf("[%s] %s\n", LogLevelError, message)
}

func Warn(message string) {
	log.Printf("[%s] %s\n", LogLevelWarn, message)
}

func Info(message string) {
	log.Printf("[%s] %s\n", LogLevelInfo, message)
}
