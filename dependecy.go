package logger

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	IsDebugMode bool
}

type Logger struct {
	cache map[time.Time]string
	cMu   sync.Mutex

	// config
	c *Config
}

func New(c *Config) *Logger {
	return &Logger{
		cache: make(map[time.Time]string),
		c:     c,
	}
}

func (l *Logger) Error(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	currentTime := getCurrentTime()

	msg := fmt.Sprint(args...)
	contentRaw := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		" ERROR",
		date,
		currentTime,
		fName,
		strconv.Itoa(line),
	)
	l.addCache(time.Now(), contentRaw)
	Error(msg)
}

func (l *Logger) Info(args ...interface{}) {
	date := getCurrentDate()
	time := getCurrentTime()
	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s %s`,
		formatTextExt(bold, brightGreen, " INFO "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(bold, msg),
	)
	fmt.Println(content)

}

func (l *Logger) InfoC(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, brightGreen, " INFO "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, brightYellow, msg))
	}
}
func (l *Logger) Warn(args ...interface{}) {

	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s %s`,
		formatTextExt(bold, orange, " WARN "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(bold, msg),
	)
	fmt.Println(content)
}

func (l *Logger) WarnC(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, orange, " WARN "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, brightYellow, msg))
	}

}
func (l *Logger) Debug(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	date := getCurrentDate()
	time := getCurrentTime()

	msg := fmt.Sprint(args...)
	content := fmt.Sprintf(`[%s] %s %s (%s:%s)`,
		formatTextExt(bold, magenta, " DBUG "),
		formatTextExt(dim, italic, date),
		formatText(underline, time),
		formatText(brightBlue, fName),
		formatText(bold, strconv.Itoa(line)),
	)
	fmt.Println(content)
	if len(args) > 0 {
		fmt.Println("↳", formatTextExt(bold, brightYellow, msg))
	}
}
