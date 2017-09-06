package mlog

import (
	"flag"
	"log"
	"os"
)

type Level uint

var DefaultLevel Level

var levelFlag = flag.String("loglevel", "info", "")

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func init() {
	switch *levelFlag {
	case "info":
		DefaultLevel = InfoLevel
	case "debug":
		DefaultLevel = DebugLevel
	case "warn":
		DefaultLevel = WarnLevel
	case "error":
		DefaultLevel = ErrorLevel
	default:
		DefaultLevel = InfoLevel
	}
}

//#tpl ["Info","Warn","Error"].map(function(elm){ return tpl.replace(new RegExp("Debug",'g'),elm).replace(new RegExp("debug",'g'),elm.toLowerCase())}).join("\n")

var debugLogger = log.New(os.Stderr, "[debug]", log.Lshortfile|log.LstdFlags)

func Debug(arg ...interface{}) {
	if DefaultLevel < DebugLevel {
		return
	}
	debugLogger.Println(arg...)
}

//#tpl

//#code

var infoLogger = log.New(os.Stderr, "[info]", log.Lshortfile|log.LstdFlags)

func Info(arg ...interface{}) {
	if DefaultLevel < InfoLevel {
		return
	}
	infoLogger.Println(arg...)
}

var warnLogger = log.New(os.Stderr, "[warn]", log.Lshortfile|log.LstdFlags)

func Warn(arg ...interface{}) {
	if DefaultLevel < WarnLevel {
		return
	}
	warnLogger.Println(arg...)
}

var errorLogger = log.New(os.Stderr, "[error]", log.Lshortfile|log.LstdFlags)

func Error(arg ...interface{}) {
	if DefaultLevel < ErrorLevel {
		return
	}
	errorLogger.Println(arg...)
}

//#code
