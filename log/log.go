package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Logger struct {
	l      *log.Logger
	config *Config
}

func New(config *Config) *Logger {
	l := &Logger{}
	l.config = config
	l.l = log.New(os.Stdout, fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.appName, l.config.reference), 0)
	return l
}

func (l *Logger) GetReference() string {
	return l.config.reference
}

// Debug log with Println format
func (l *Logger) Debug(v ...interface{}) {
	if l.config.level > DEBUG {
		return
	}

	l.l.Println(l.formatLog("DEBUG", v...)...)
}

// Info log with Println format
func (l *Logger) Info(v ...interface{}) {
	if l.config.level > INFO {
		return
	}

	l.l.Println(l.formatLog("INFO", v...)...)
}

func (l *Logger) Warn(v ...interface{}) {
	if l.config.level > WARN {
		return
	}

	l.l.Println(l.formatLog("WARN", v...)...)
}

func (l *Logger) Error(v ...interface{}) {
	if l.config.level > ERROR {
		return
	}

	l.l.Println(l.formatLog("ERROR", v...)...)
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.config.level > FATAL {
		return
	}

	l.l.Println(l.formatLog("FATAL", v...)...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.config.level > DEBUG {
		return
	}

	format, v = l.formatLogf("DEBUG", format, v...)
	l.l.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.config.level > INFO {
		return
	}

	format, v = l.formatLogf("INFO", format, v...)
	l.l.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.config.level > WARN {
		return
	}

	format, v = l.formatLogf("WARN", format, v...)
	l.l.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.config.level > ERROR {
		return
	}

	format, v = l.formatLogf("ERROR", format, v...)
	l.l.Printf(format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.config.level > FATAL {
		return
	}

	format, v = l.formatLogf("FATAL", format, v...)
	l.l.Printf(format, v...)
}

// Format the log to contain the log levels
func (l *Logger) formatLog(logType string, v ...interface{}) []interface{} {
	var n []interface{}
	n = append(n, "["+logType+"] ")
	file, line := l.getFileLine(3)

	n = append(n, file+":"+strconv.Itoa(line)+":")
	n = append(n, v...)
	return n
}

// Format the log to contain the log levels
func (l *Logger) formatLogf(logType string, format string, v ...interface{}) (string, []interface{}) {
	var n []interface{}
	prefix := "[%s] "
	n = append(n, logType)
	file, line := l.getFileLine(3)

	prefix += "%s:%d: "
	format = prefix + format
	n = append(n, file)
	n = append(n, line)
	n = append(n, v...)
	return format, n
}

func (l *Logger) getFileLine(n int) (string, int) {
	_, file, line, _ := runtime.Caller(n)
	// If you want the short path not the full file path, you can uncomment everything below
	if l.config.filePathSize == SHORT {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
	}

	return file, line
}
