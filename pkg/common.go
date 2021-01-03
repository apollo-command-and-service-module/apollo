package pkg

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func IdGenerator() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x%x%x", b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}

func FormatDate(currentTime time.Time) string {
	layOut := "2006-01-02 15:04:05 -0700 MST"
	ConvertDate := currentTime.Format(layOut)

	return ConvertDate
}

func ConvertUTC(strDate string) time.Time {
	layOut := "2006-01-02 15:04:05 -0700 MST"
	covDate, err := time.Parse(layOut, strDate)
	if err != nil {
		log.Fatal(err)
	}
	return covDate
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error, worker int, jobId string) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("worker%d: ID: %s error %s", worker, jobId, err))
	return
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
