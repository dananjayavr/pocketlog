package pocketlog

// Level respresents an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes
	LevelDebug = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace errors
	LevelError
)

/*
// New returns you a logger, ready to log at the required threshold.
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

// Infof formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	_, _ = fmt.Printf(format+"n", args...)
}
*/
