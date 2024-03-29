package logging_test

import (
	"errors"
	"flag"
	"io"
	"os"
	"time"

	logger "github.com/apollo-command-and-service-module/apollo/pkg/logging"
	"github.com/rs/zerolog"
)

func setup(deBug bool) *logger.Logger {
	var w io.Writer
	w = os.Stdout
	zerolog.TimeFieldFormat = ""

	zerolog.TimestampFunc = func() time.Time {
		return time.Date(2008, 1, 8, 17, 5, 05, 0, time.UTC)
	}

	return logger.New(w, deBug)
}

func ExamplePrint() {
	l := setup(true)
	l.Print("hello world")

	// Output: {"level":"debug","time":1199811905,"message":"hello world"}
}

// Simple logging example using the Printf function in the log package
func ExamplePrintf() {
	l := setup(true)
	l.Printf("hello %s", "world")

	// Output: {"level":"debug","time":1199811905,"message":"hello world"}
}

// Example of a log with no particular "level"
func ExampleLog() {
	l := setup(true)
	l.Log().Msg("hello world")

	// Output: {"time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "debug")
func ExampleDebug() {
	l := setup(true)
	l.Debug().Msg("hello world")

	// Output: {"level":"debug","time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "info")
func ExampleInfo() {
	l := setup(true)
	l.Info().Msg("hello world")

	// Output: {"level":"info","time":1199811905,"message":"hello world"}
}

func ExamplePrintInfo() {
	l := setup(true)
	fakeError := errors.New("hello world error")
	l.PrintInfo(fakeError)

	// Output: {"level":"info","time":1199811905,"message":"hello world error"}
}

func ExamplePrintInfof() {
	l := setup(true)
	fakeError := errors.New("hello world error")
	l.PrintInfof("%s", fakeError)
	// Output: {"level":"info","time":1199811905,"message":"hello world error"}
}

// Example of a log at a particular "level" (in this case, "warn")
func ExampleWarn() {
	l := setup(true)
	l.Warn().Msg("hello world")

	// Output: {"level":"warn","time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "error")
func ExampleError() {
	l := setup(true)
	l.Error().Msg("hello world")

	// Output: {"level":"error","time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "fatal")
func ExampleFatal() {
	l := setup(true)
	err := errors.New("A repo man spends his life getting into tense situations")
	service := "myservice"

	l.Fatal().
		Err(err).
		Str("service", service).
		Msgf("Cannot start %s", service)

	// Outputs: {"level":"fatal","time":1199811905,"error":"A repo man spends his life getting into tense situations","service":"myservice","message":"Cannot start myservice"}
}

// This example uses command-line flags to demonstrate various outputs
// depending on the chosen log level.
func Example() {
	l := setup(true)
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	l.Debug().Msg("This message appears only when log level set to Debug")
	l.Info().Msg("This message appears when log level set to Debug or Info")

	if e := l.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}

	// Output: {"level":"info","time":1199811905,"message":"This message appears when log level set to Debug or Info"}
}
