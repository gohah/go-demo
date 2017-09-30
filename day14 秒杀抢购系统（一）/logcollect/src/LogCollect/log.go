package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"oldboy/logcollect/src/kafka"
)

/*
* log format: [2016-12-05 19:00:00] [LogCollect] [hostname] [NOTICE] [0] {}
 */

const (
	LOG_TOPIC    = "LogCollect"
	LOG_TOPIC_WF = "LogCollectWf"
)

type LogMap map[string]interface{}

type Writer struct {
}

type WriterWf struct {
}

func makePrefix(host string, level string) string {
	curTime := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] [LogCollect] [%s] [%s] [0] ", curTime, host, level)
}

func (l *Writer) Write(p []byte) (n int, err error) {
	return n, kafka.SendMessage(LOG_TOPIC, fmt.Sprintf("%s", strings.TrimRight(string(p), "\n")))
}

func (l *WriterWf) Write(p []byte) (n int, err error) {
	return n, kafka.SendMessage(LOG_TOPIC_WF, fmt.Sprintf("%s", strings.TrimRight(string(p), "\n")))
}

type Logger struct {
	l     *log.Logger
	wf    *log.Logger
	host  string
	appId string
}

func NewLogger(host string) *Logger {
	return &Logger{
		//l:  log.New(&Writer{}, fmt.Sprintf("[%s] ", host), log.Ldate|log.Ltime),
		//wf: log.New(&WriterWf{}, fmt.Sprintf("[%s] ", host), log.Ldate|log.Ltime),
		l:     log.New(&Writer{}, "", 0),
		wf:    log.New(&WriterWf{}, "", 0),
		host:  host,
		appId: "LogCollect",
	}
}
func (l *Logger) Notice(format string, v ...interface{}) {
	prefix := makePrefix(l.host, "NOTICE")
	msg := fmt.Sprintf(prefix+format, v...)
	l.l.Printf(msg)
}
func (l *Logger) Warn(format string, v ...interface{}) {
	prefix := makePrefix(l.host, "WARN")
	msg := fmt.Sprintf(prefix+format, v...)
	l.wf.Printf(msg)
}
func (l *Logger) Trace(format string, v ...interface{}) {
	prefix := makePrefix(l.host, "TRACE")
	msg := fmt.Sprintf(prefix+format, v...)
	l.l.Printf(msg)
}
func (l *Logger) Debug(format string, v ...interface{}) {
	prefix := makePrefix(l.host, "DEBUG")
	msg := fmt.Sprintf(prefix+format, v...)
	l.l.Printf(msg)
}
func (l *Logger) Fatal(format string, v ...interface{}) {
	prefix := makePrefix(l.host, "FATAL")
	msg := fmt.Sprintf(prefix+format, v...)
	fmt.Fprintln(os.Stderr, msg)
	l.wf.Printf(msg)
}

func (l *Logger) Monitor(dataMap map[string]interface{}) {
	prefix := makePrefix(l.host, "STAT")
	data, _ := json.Marshal(dataMap)
	msg := fmt.Sprintf("%s%v", prefix, string(data))
	l.l.Printf(msg)
	dataMap = nil
}
