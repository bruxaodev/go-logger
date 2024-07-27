package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	Green   = "\033[97;42m"
	White   = "\033[90;47m"
	Yellow  = "\033[90;43m"
	Red     = "\033[97;41m"
	Blue    = "\033[97;44m"
	Magenta = "\033[97;45m"
	Cyan    = "\033[97;46m"
	Reset   = "\033[0m"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
	writer io.Writer
}

func New(prefix, color string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug: log.New(writer, fmt.Sprintf(
			"%s[%s]%s - %s[DEBUG]%s ",
			color,
			prefix,
			Reset,
			Cyan,
			Reset,
		), logger.Flags()),
		info: log.New(writer, fmt.Sprintf(
			"%s[%s]%s - %s[INFO]%s  ",
			color,
			prefix,
			Reset,
			Green,
			Reset,
		), logger.Flags()),
		warn: log.New(writer, fmt.Sprintf(
			"%s[%s]%s - %s[Warn]%s  ",
			color,
			prefix,
			Reset,
			Yellow,
			Reset,
		), logger.Flags()),
		error: log.New(writer, fmt.Sprintf(
			"%s[%s]%s - %s[Error]%s ",
			color,
			prefix,
			Reset,
			Red,
			Reset,
		), logger.Flags()),
		writer: writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	m := fmt.Sprintf("%10s[DEBUG]%s | %s", Cyan, Reset, fmt.Sprint(v...))
	l.debug.Print(m)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	m := fmt.Sprintf("%10s[DEBUG]%s | %s", Cyan, Reset, fmt.Sprintf(format, v...))
	l.debug.Print(m)
}
func (l *Logger) Info(v ...interface{}) {
	m := fmt.Sprintf("%10s[INFO]%s  | %s", Green, Reset, fmt.Sprint(v...))
	l.info.Print(m)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	m := fmt.Sprintf("%10s[INFO]%s  | %s", Green, Reset, fmt.Sprintf(format, v...))
	l.info.Print(m)
}
func (l *Logger) Warn(v ...interface{}) {
	m := fmt.Sprintf("%10s[WARN]%s  | %s", Yellow, Reset, fmt.Sprint(v...))
	l.warn.Print(m)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	m := fmt.Sprintf("%10s[WARN]%s  | %s", Yellow, Reset, fmt.Sprintf(format, v...))
	l.warn.Print(m)
}
func (l *Logger) Error(v ...interface{}) {
	m := fmt.Sprintf("%10s[ERROR]%s | %s", Red, Reset, fmt.Sprint(v...))
	l.error.Print(m)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	m := fmt.Sprintf("%10s[ERROR]%s | %s", Red, Reset, fmt.Sprintf(format, v...))
	l.error.Print(m)
}

// func main() {
// 	logger := new("main", Green)
// 	logger.Debug("Debug message")
// 	logger.Debugf("Debug message %s", "formatted")
// 	logger.Info("Info message")
// 	logger.Infof("Info message %s", "formatted")
// 	logger.Warn("Warn message")
// 	logger.Warnf("Warn message %s", "formatted")
// 	logger.Error("Error message")
// 	logger.Errorf("Error message %s", "formatted")
// }
