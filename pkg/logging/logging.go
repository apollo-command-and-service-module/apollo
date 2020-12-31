package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/labstack/gommon/color"
)

var (
	infoLog    Logger
	warningLog Logger
	errorLog   Logger
)

type Logger interface {
	log() *log.Logger
}

type Infolog struct {
	Infolog *log.Logger `json:"Infolog"`
	colour  *color.Color
}

type Warninglog struct {
	Warninglog *log.Logger `json:"Warninglog"`
}

type Errorlog struct {
	Errorlog *log.Logger `json:"Errorlog"`
}

func (i *Infolog) log() *log.Logger {
	return &*i.Infolog
}

func (i *Warninglog) log() *log.Logger {
	return &*i.Warninglog
}

func (i *Errorlog) log() *log.Logger {
	return &*i.Errorlog
}

//Terminal Info Function (not wrapped in a logger)

func TerminalInfo(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func TerminalWarning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

//Infolog functions

//Info ...
func Info(format string, a ...interface{}) {
	infoLog.log().Printf(format, a...)
}

//Infoln ...
func Infoln(a ...interface{}) {
	color.Enable()
	//color.Println(color.Red("red"))
	infoLog.log().Println(a...)
}

//Infofatal ...
func Infofatal(a ...interface{}) {
	infoLog.log().Fatal(a...)
}

//Warning log functions

//Warning ...
func Warning(format string, a ...interface{}) {
	warningLog.log().Printf(format, a...)
}

//Warningln ...
func Warningln(a ...interface{}) {
	warningLog.log().Println(a...)
}

//Warningfatal ...
func Warningfatal(a ...interface{}) {
	warningLog.log().Fatal(a...)
}

//Error log functions

//Error ...
func Error(format string, a ...interface{}) {
	errorLog.log().Printf(format, a...)
}

//Errorln ...
func Errorln(a ...interface{}) {
	errorLog.log().Println(a...)
}

//Errorfatal ...
func Errorfatal(a ...interface{}) {
	errorLog.log().Fatal(a...)
}

func init() {
	infoLog = &Infolog{Infolog: log.New(ioutil.Discard, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)}
	warningLog = &Warninglog{Warninglog: log.New(ioutil.Discard, "WARNING\t", log.Ldate|log.Ltime|log.Lshortfile)}
	errorLog = &Errorlog{Errorlog: log.New(ioutil.Discard, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)}
}

func LogSetup(output io.Writer, logtype ...interface{}) {

	color.SetOutput(output)

	for _, i := range logtype {

		switch i {
		case "info":
			infoLog = &Infolog{Infolog: log.New(output, color.Cyan("INFO\t"), log.Ldate|log.Ltime|log.Lshortfile)}
		case "warning":
			warningLog = &Warninglog{Warninglog: log.New(output, color.Yellow("WARNING\t"), log.Ldate|log.Ltime|log.Lshortfile)}
		case "error":
			errorLog = &Errorlog{Errorlog: log.New(output, color.Red("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)}
		}
	}

}