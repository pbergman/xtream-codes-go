package xtream_codes_go

import "net/url"

type ApiClientConfig interface {
	GetHost() *url.URL
	GetUsername() string
	GetPassword() string
}

func NewApiClientConfig(host, username, password string) (ApiClientConfig, error) {
	uri, err := url.Parse(host)

	if err != nil {
		return nil, err
	}

	return &apiClientConfig{host: uri, username: username, password: password}, nil
}

type apiClientConfig struct {
	host     *url.URL
	username string
	password string
}

func (a *apiClientConfig) GetHost() *url.URL {
	return a.host
}

func (a *apiClientConfig) GetUsername() string {
	return a.username
}

func (a *apiClientConfig) GetPassword() string {
	return a.password
}
