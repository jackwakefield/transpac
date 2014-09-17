// Copyright 2014 Jack Wakefield
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/kdar/factorlog"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	location       = kingpin.Arg("location", "Path or URL of the proxy auto-config (PAC) file.").Required().String()
	verbose        = kingpin.Flag("verbose", "Enable verbose logging.").Bool()
	cacheDirectory = kingpin.Flag("cache-dir", "The directory where files are cached to.").Default("/var/cache/transpac").String()
	cacheLength    = kingpin.Flag("cache-length", "The length of time in seconds to cache downloaded files.").Default("86400").Int64()
	serverPort     = kingpin.Flag("server-port", "The port the proxy server will listen on.").Default("8080").Int()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	if *verbose {
		logger.SetMinMaxSeverity(factorlog.DEBUG, factorlog.PANIC)
	} else {
		logger.SetMinMaxSeverity(factorlog.INFO, factorlog.PANIC)
	}

	// ensure the cache directory exists
	if err := os.MkdirAll(*cacheDirectory, 0755); err != nil {
		logger.Fatalf("Failed to create cache directory (%s)", err)
	}

	logger.Debugf("Creating proxy auto-config parser for '%s'", *location)
	p, err := newParser(*location)

	if err != nil {
		logger.Fatalf("Failed to load proxy file (%s)", err)
	}

	logger.Debugf("Creating proxy client")
	c := newProxyClient(p)

	logger.Debugf("Creating proxy server")
	s := newServer(p, c)

	if err := s.listen(); err != nil {
		logger.Fatalf("The proxy server failed to listen (%s)", err)
	}
}
