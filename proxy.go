package main

import (
	"net/url"
	"strings"
)

func GetProxyURL(conf *attackOpts) (*url.URL, error) {
	if conf.proxy == "" {
		return nil, nil
	}
	proxyUrl, err := url.Parse(conf.proxy)
	if err != nil {
		return nil, err
	}
	if conf.proxyAuth != "" {
		proxyAuthSplit := strings.Split(conf.proxyAuth, ":")
		proxyUser := proxyAuthSplit[0]
		if len(proxyAuthSplit) == 1 {
			return nil, err
		}
		proxyPassword := proxyAuthSplit[1]
		proxyUrl.User = url.UserPassword(proxyUser, proxyPassword)
	}
	return proxyUrl, nil
}
