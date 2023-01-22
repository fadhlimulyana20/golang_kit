package log

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	env := os.Getenv("ENV")
	if env == "development" {
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: false,
			TimestampFormat:  "2006-01-02 15:04:05",
			FullTimestamp:    true,
			PadLevelText:     true,
		})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)

	// logrus show line numberW
	log.SetReportCaller(true)
}

func Make(r *http.Request) *log.Entry {
	if r != nil {
		return log.WithFields(log.Fields{
			"at":     time.Now().UTC(),
			"url":    r.URL.Path,
			"method": r.Method,
			"ip":     r.RemoteAddr,
		})
	}

	return log.WithFields(log.Fields{
		"at": time.Now().UTC(),
	})
}
