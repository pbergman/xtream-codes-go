package xtream_codes_go

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pbergman/logger"
)

const (
	playerApi string = "player_api.php"
)

func NewApiClient(config ApiClientConfig, logger *logger.Logger, client *http.Client, dumper io.Writer) (*ApiClient, error) {

	if nil == client {
		client = &http.Client{}
	}

	var transport = client.Transport

	if nil == transport {
		transport = http.DefaultTransport
	}

	client.Transport = &ApiTransport{
		inner:  transport,
		logger: logger,
		dumper: dumper,
		config: config,
	}

	var api = &ApiClient{client: client}

	if err := authenticate(api, logger); err != nil {
		return nil, err
	}

	api.client.Transport.(*ApiTransport).loginInfo = api.loginInfo

	return api, nil
}

type ApiClient struct {
	client    *http.Client
	loginInfo *LoginInfo
}

func (a *ApiClient) context(action string, params map[string]string) context.Context {
	var values = make(url.Values)

	values.Set("action", action)

	for k, v := range params {
		values.Set(k, v)
	}

	return context.WithValue(context.Background(), "values", values)
}

func (a *ApiClient) fetch(ctx context.Context, path string, data any) error {

	request, err := http.NewRequestWithContext(ctx, "GET", path, nil)

	if err != nil {
		return err
	}

	response, err := a.client.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return fmt.Errorf("unexpected status code (%d) returned for '%s'", response.StatusCode, response.Request.URL.RequestURI())
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

func (a *ApiClient) streamUrl(stream string, id int, extension string) string {
	return fmt.Sprintf(
		"%s://%s/%s/%s/%s/%d.%s",
		a.loginInfo.ServerInfo.ServerProtocol,
		a.loginInfo.ServerInfo.Url,
		stream,
		a.loginInfo.UserInfo.Username,
		a.loginInfo.UserInfo.Password,
		id,
		extension,
	)
}
