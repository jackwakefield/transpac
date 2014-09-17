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
	"fmt"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/jackwakefield/gopac"
)

type server struct {
	parser *gopac.Parser
	client *http.Client
	http   *goproxy.ProxyHttpServer
}

func newServer(parser *gopac.Parser, client *http.Client) (s *server) {
	s = &server{
		parser: parser,
		client: client,
		http:   goproxy.NewProxyHttpServer(),
	}

	// determine whether a proxy entry is available for the URL/host for each
	// HTTP request and if so, forward it to the proxy client
	s.http.OnRequest(s.isForwardable()).DoFunc(s.forward)

	return
}

func (s *server) isForwardable() goproxy.ReqConditionFunc {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) bool {
		// find a proxy entry for the request URL/host
		entry, err := s.parser.FindProxy(req.URL.String(), req.URL.Host)

		// ensure no error occurred when finding an entry and that the entry
		// doesn't indicate the connection should be made directly
		if err == nil && entry != "DIRECT" {
			return true
		}

		return false
	}
}

func (s *server) forward(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// retrieve a response for the request using the proxy client
	if resp, err := s.client.Get(r.URL.String()); err == nil {
		return r, resp
	}

	return r, nil
}

func (s *server) listen() error {
	address := fmt.Sprintf(":%d", *serverPort)
	logger.Infof("Proxy server listening on address '%s'", address)

	return http.ListenAndServe(address, s.http)
}
