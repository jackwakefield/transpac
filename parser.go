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
	"io/ioutil"
	"net/http"

	"github.com/jackwakefield/gopac"
)

func newParser(location string) (*gopac.Parser, error) {
	parser := new(gopac.Parser)

	// determine whether the location exists indicating it's a file path,
	// otherwise, attempt to load the cache file for the location, or download
	// and then cache it in the configured directory
	if fileExists(location) {
		logger.Info("Parsing proxy auto-config file")

		if err := parser.Parse(location); err != nil {
			return nil, err
		}
	} else {
		logger.Debug("Checking proxy auto-config cache")
		cache := newCacheItem(location)

		// determine whether the cache file exists or is older than the
		// configured cache length
		if !cache.exists() || cache.outOfDate() {
			logger.Info("Downloading proxy auto-config file")

			resp, err := http.Get(location)

			if err != nil {
				return nil, err
			}

			logger.Debug("Reading response")

			data, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			if err != nil {
				return nil, err
			}

			logger.Debug("Caching proxy auto-config file")

			// save the proxy auto-config to the cache file
			err = cache.save(data)

			if err != nil {
				return nil, err
			}
		}

		// parse the cached proxy auto-config file
		return newParser(cache.path())
	}

	return parser, nil
}
