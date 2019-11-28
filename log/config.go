package log

import (
	"github.com/google/uuid"
)

type Config struct {
	// The supported log levels are as follows
	// DEBUG < INFO < WARN < ERROR < FATAL
	// If a log level is specified all logs with level below the specified level are ignored
	// Eg. If INFO is selected, All DEBUG logs are ignored
	// If ERROR is selected all logs except ERROR and FATAL are ignored
	level Level

	// Size of the file to be printed, there are two possible values FULL, SHORT
	// SHORT - Only the file name is displayed
	// FULL - File name along with full file path is specified
	// SHORT is used by default
	filePathSize int

	// Log Reference (context) ID to be added to each log
	// This can be used to search relevent logs for the context
	reference string

	// Name of the App sending the log
	appName string
}

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
	FATAL

	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"

	SHORT = iota
	FULL

	FilePathSizeShort = "SHORT"
	FilePathSizeFull  = "FULL"
)

// Creates Log Config using Application name.
// By default Reference is random uuid,
// Level is INFO,
// File Path Size is SHORT,
func NewConfig(appName string) *Config {
	ref := uuid.New()

	return &Config{
		reference:    ref.String(),
		level:        INFO,
		filePathSize: SHORT,
		appName:      appName,
	}
}

// Setting the log level to Log Config. Use the log level string enum
// ( log.LevelDebug (or) log.LevelInfo (or) log.LevelWarm (or) log.LevelError (or) log.LevelFatal)
func (c *Config) SetLevel(level string) {
	switch level {
	case LevelDebug:
		c.level = DEBUG
	case LevelInfo:
		c.level = INFO
	case LevelWarn:
		c.level = WARN
	case LevelError:
		c.level = ERROR
	case LevelFatal:
		c.level = FATAL
	default:
		c.level = INFO
	}
}

// Setting the log level to Log Config. Use the log level enum
// ( log.DEBUG (or) log.INFO (or) log.WARN (or) log.ERROR (or) log.FATAL)
func (c *Config) SetLevelEnum(level Level) {
	c.level = level
}

// Setting the file path size to be logged
// FilePathSizeShort or filePathSize
func (c *Config) SetFilePathSize(filePathSize string) {
	switch filePathSize {
	case FilePathSizeShort:
		c.filePathSize = SHORT
	case FilePathSizeFull:
		c.filePathSize = FULL
	default:
		c.filePathSize = SHORT
	}
}

// Setting the file path size to be logged
// SHORT or FULL
func (c *Config) SetFilePathSizeEnum(filePathSize int) {
	c.filePathSize = filePathSize
}

// Setting the log reference to Config
func (c *Config) SetReference(ref string) {
	c.reference = ref
}
