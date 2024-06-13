package xtream_codes_go

import (
	"fmt"
	"github.com/pbergman/logger"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ApiTransport struct {
	inner     http.RoundTripper
	config    ApiClientConfig
	logger    *logger.Logger
	dumper    io.Writer
	loginInfo *LoginInfo
}

func (t *ApiTransport) update(request *http.Request, stopwatch *stopwatch) *http.Request {
	var query = request.URL.Query()

	if value, ok := request.Context().Value("values").(url.Values); ok {
		if len(query) > 0 {
			for key, values := range value {
				if false == query.Has(key) {
					query[key] = values
				}
			}
		} else {
			query = value
		}
	}

	if '/' != request.URL.Path[0] {
		request.URL.Path = "/" + request.URL.Path
	}

	if nil == t.loginInfo {
		query.Set("username", t.config.GetUsername())
		query.Set("password", t.config.GetPassword())
		request.URL.Scheme = t.config.GetHost().Scheme
		request.URL.Host = t.config.GetHost().Host
	} else {
		query.Set("username", t.loginInfo.UserInfo.Username)
		query.Set("password", t.loginInfo.UserInfo.Password)
		request.URL.Scheme = t.loginInfo.ServerInfo.ServerProtocol
		request.URL.Host = t.loginInfo.ServerInfo.Url
	}

	request.URL.RawQuery = query.Encode()

	if nil != stopwatch {
		var ctx = httptrace.WithClientTrace(request.Context(), &httptrace.ClientTrace{
			DNSStart: func(info httptrace.DNSStartInfo) {
				stopwatch.GetEvent("dns").Start()
			},
			DNSDone: func(info httptrace.DNSDoneInfo) {
				stopwatch.GetEvent("dns").Stop()
			},
			ConnectStart: func(_, _ string) {
				stopwatch.GetEvent("tcp").Start()
			},
			ConnectDone: func(network, addr string, err error) {
				stopwatch.GetEvent("tcp").Stop()
			},
			WroteRequest: func(info httptrace.WroteRequestInfo) {
				stopwatch.GetEvent("srv").Start()
			},
			GotFirstResponseByte: func() {
				stopwatch.GetEvent("srv").Stop()
			},
		})

		return request.WithContext(ctx)
	}

	return request
}

func (r *ApiTransport) getRequestUri(uri url.URL) string {
	var values = uri.Query()

	for _, field := range []string{"username", "password"} {
		if values.Has(field) {

			var value = values.Get(field)

			if len(value) > 5 {
				value = value[:3]
			} else {
				value = ""
			}

			values.Set(field, value+"******")
		}
	}

	uri.RawQuery = values.Encode()

	if ret, err := url.QueryUnescape(uri.RequestURI()); err != nil {
		return uri.RequestURI()
	} else {
		return ret
	}
}

func (t *ApiTransport) dump(out []byte, prefix string) {
	var lines = strings.Split(string(out), "\r\n")

	for i, c := 0, len(lines); i < c; i++ {
		_, _ = t.dumper.Write([]byte(prefix + lines[i] + "\n"))
	}
}

func (t *ApiTransport) RoundTrip(request *http.Request) (*http.Response, error) {

	var timer *stopwatch

	if nil != t.logger {
		timer = &stopwatch{events: make(map[string]*stopwatchEvent)}
	}

	request = t.update(request, timer)

	if nil != t.dumper {
		if out, err := httputil.DumpRequestOut(request, true); err == nil {
			t.dump(out, "> ")
		}
	}

	resp, err := t.inner.RoundTrip(request)

	if err != nil {
		return nil, err
	}

	if nil != t.logger {

		t.logger.Debug(logger.Message(fmt.Sprintf("%s %s %s %d", request.Method, t.getRequestUri(*request.URL), request.Proto, resp.StatusCode), timer.getContext()))
	}

	if nil != t.dumper {
		if out, err := httputil.DumpResponse(resp, true); err == nil {
			t.dump(out, "< ")
		}
	}

	return resp, nil
}
