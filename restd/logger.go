// Copyright 2016 Vincent Landgraf. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"log"
	"log/syslog"
	"os"

	"github.com/codegangsta/negroni"
)

var (
	apiLogger   *negroni.Logger
	confdLogger *log.Logger
)

func initLogger() {
	apiLogger = negroni.NewLogger()
	flags := 0
	var writer io.Writer = os.Stdout

	if syslogEnabled {
		var err error
		writer, err = syslog.New(syslog.LOG_DAEMON|syslog.LOG_NOTICE, "restd")
		if err != nil {
			log.Fatalf("Can't connect to syslog: %s", err)
		}
	} else {
		flags = log.LstdFlags
	}

	apiLogger.SetPrefix("api ")
	apiLogger.SetFlags(flags)
	apiLogger.SetOutput(writer)

	log.SetPrefix("main ")
	log.SetFlags(flags)
	log.SetOutput(writer)

	confdLogger = log.New(writer, "confd ", flags)
}
