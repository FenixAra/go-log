package log

import "strings"

type Config struct {
	// The supported log levels are as follows
	// DEBUG < INFO < WARN < ERROR < FATAL
	// If a log level is specified all logs with level below the specified level are ignored
	// Eg. If INFO is selected, All DEBUG logs are ignored
	// If ERROR is selected all logs except ERROR and FATAL are ignored
	Level Level

	// Log levels in string format.
	// The supported log level strings are Debug, Info, Warn, Error, Fatal
	// You can specify log level using Level Enum or string
	// The Enum value is given first preference
	LevelStr string

	// Size of the file to be printed, there are two possible values FULL, SHORT
	// SHORT - Only the file name is displayed
	// FULL - File name along with full file path is specified
	// SHORT is used by default
	FilePathSize int

	// Log Reference (context) ID to be added to each log
	// This can be used to search relevent logs for the context
	Reference string

	// Name of the App sending the log
	AppName string
}

// Creates Log Config using Reference, Level in string ( DEBUG (OR) INFO (OR) WARN (OR) ERROR (OR) FATAL),
// File Path Size to be added to log (SHORT, FULL), Application name
func NewConfig(ref, levelStr, filePathSizeStr, appName string) *Config {
	var level Level
	var filePathSize int
	levelStr = strings.ToUpper(levelStr)
	filePathSizeStr = strings.ToUpper(filePathSizeStr)
	switch levelStr {
	case Debug:
		level = DEBUG
	case Info:
		level = INFO
	case Warn:
		level = WARN
	case Error:
		level = ERROR
	case Fatal:
		level = FATAL
	default:
		level = INFO
	}

	switch filePathSizeStr {
	case FilePathShort:
		filePathSize = SHORT
	case FilePathFull:
		filePathSize = FULL
	default:
		filePathSize = SHORT
	}

	return &Config{
		Reference:    ref,
		Level:        level,
		FilePathSize: filePathSize,
		AppName:      appName,
	}
}

// Setting the log level to Log Config. Use the log level enum
// ( log.DEBUG (or) log.INFO (or) log.WARN (or) log.ERROR (or) log.FATAL)
func (c *Config) SetLevel(level Level) {
	c.Level = level
}

// Setting the file path size to be logged
// SHORT or FULL
func (c *Config) SetFilePathSize(filePathSize int) {
	c.FilePathSize = filePathSize
}

// Setting the log reference to Config
func (c *Config) SetReference(ref string) {
	c.Reference = ref
}
