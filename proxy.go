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
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/jackwakefield/gopac"
)

var errorProxyRetrieval = errors.New("unable to retrieve proxy")

func newProxyClient(parser *gopac.Parser) *http.Client {
	return &http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			// retrieve the proxy entry for the requested URL and host
			entry, err := parser.FindProxy(req.URL.String(), req.URL.Host)

			if err == nil {
				// strip the PROXY prefix, the server delimiter, and then split
				// the entry to retrieve an array of proxy servers to be used for
				// the request
				entry = strings.Replace(entry, "PROXY ", "", -1)
				entry = strings.Replace(entry, "; ", "", -1)
				hosts := strings.Split(entry, " ")

				// ensure a proxy server was found
				if len(hosts) > 0 && hosts[0] != "" {
					// return the proxy server to be used to load the request
					return url.Parse("http://" + hosts[0])
				}
			}

			return nil, errorProxyRetrieval
		},
	}}
}
