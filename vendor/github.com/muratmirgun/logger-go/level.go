package log

import "strings"

// Type that defines the logging level
type Level int

// These are the supported log levels.
const (
	DEBUG Level = iota
	INFO  Level = iota
	WARN  Level = iota
	ERROR Level = iota
)

// String converts a log level to its string representation
func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	}
	panic("Unknown log level")
}

func Parse(s string) Level {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	}
	panic("Unknown log level " + s)
}
