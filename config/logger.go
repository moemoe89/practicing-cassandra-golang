//
//  Practicing Cassandra
//
//  Copyright © 2020. All rights reserved.
//

package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// InitLog will create a variable that represent the logrus.Entry
func InitLog() *logrus.Entry {

	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
		DisableColors:   true,
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	host, err := os.Hostname()
	if err != nil {
		log.Error("", err.Error())
		host = "unknown"
	}

	return log.WithField("host", host)
}
