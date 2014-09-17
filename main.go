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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	locationKey       = "location"
	verboseKey        = "verbose"
	cacheDirectoryKey = "cache-dir"
	cacheLengthKey    = "cache-length"
	serverPortKey     = "server-port"
)

func main() {
	transpacCommand := &cobra.Command{
		Use:   "transpac",
		Short: "A transparent proxy which uses proxy auto-config (PAC) files for forwarding",
		Run:   run,
	}

	flags := transpacCommand.Flags()

	flags.String(locationKey, "", "Path or URL of the proxy auto-config (PAC) file")
	flags.Bool(verboseKey, false, "Enable verbose logging")
	flags.String(cacheDirectoryKey, "/var/cache/transpac", "The directory where files are cached to")
	flags.Int64(cacheLengthKey, 86400, "The length of time in seconds to cache downloaded files")
	flags.Int(serverPortKey, 8080, "The port the proxy server will listen on")

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/transpac")
	viper.AddConfigPath("$HOME/.config/transpac")
	viper.BindPFlag(locationKey, flags.Lookup(locationKey))
	viper.BindPFlag(verboseKey, flags.Lookup(verboseKey))
	viper.BindPFlag(cacheDirectoryKey, flags.Lookup(cacheDirectoryKey))
	viper.BindPFlag(cacheLengthKey, flags.Lookup(cacheLengthKey))
	viper.BindPFlag(serverPortKey, flags.Lookup(serverPortKey))
	viper.ReadInConfig()

	transpacCommand.Execute()
}

func run(cmd *cobra.Command, args []string) {
	if viper.GetBool("verbose") {
		logger.SetMinMaxSeverity(factorlog.DEBUG, factorlog.PANIC)
	} else {
		logger.SetMinMaxSeverity(factorlog.INFO, factorlog.PANIC)
	}

	// ensure the cache directory exists
	if err := os.MkdirAll(viper.GetString(cacheDirectoryKey), 0755); err != nil {
		logger.Fatalf("Failed to create cache directory (%s)", err)
	}

	location := viper.GetString(locationKey)

	if len(location) == 0 {
		logger.Fatal("You must provide a location for the proxy auto-config file")
	}

	logger.Debugf("Creating proxy auto-config parser for '%s'", location)
	p, err := newParser(location)

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
