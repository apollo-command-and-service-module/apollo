package logging

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestSetLogTypeInfo(t *testing.T) {

	var buf = &bytes.Buffer{}
	customInfo := "this is a custom info message"

	//setup logging of type error
	LogSetup(buf, "info")

	Infoln(customInfo)

	got := buf.String()
	wanted := strings.Contains(got, "INFO")
	if !wanted {
		t.Logf("output should cointain `INFO` header, instead got: %s", got)
		t.Fail()
	}

}

func TestSetLogTypeWarning(t *testing.T) {

	var buf = &bytes.Buffer{}
	customWarning := "this is a custom warning message"

	//setup logging of type error
	LogSetup(buf, "warning")

	Warningln(customWarning)

	got := buf.String()
	wanted := strings.Contains(got, "WARNING")
	if !wanted {
		t.Logf("output should cointain `WARNING` header, instead got: %s", got)
		t.Fail()
	}

}

func TestSetLogTypeError(t *testing.T) {

	var buf = &bytes.Buffer{}
	customError := errors.New("this is a custom error")

	//setup logging of type error
	LogSetup(buf, "error")

	Errorln(customError)

	got := buf.String()
	wanted := strings.Contains(got, "ERROR")
	if !wanted {
		t.Logf("output should cointain `ERROR` header, instead got: %s", got)
		t.Fail()
	}

}
