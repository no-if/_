package http

import (
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"time"
)

type transport struct {
	*http.Transport
}

func NewTransport() *transport {
	return &transport{&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
	}}
}

func (self *transport) SetTimeout(millisecond int64) *transport {
	self.Dial = (&net.Dialer{
		Timeout:   time.Duration(millisecond) * time.Millisecond,
		KeepAlive: time.Duration(millisecond) * time.Millisecond,
	}).Dial
	return self
}

func (self *transport) SetProxy(proxy_url string) *transport {
	if proxy_url == "" {
		self.Proxy = http.ProxyFromEnvironment
		self.Dial = (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).Dial
		return self
	}

	if uri, err := url.Parse(proxy_url); err == nil {
		if dialer, err := proxy.FromURL(uri, proxy.Direct); err == nil {
			//socks5
			self.Dial = dialer.Dial
		} else {
			self.Proxy = func(*http.Request) (*url.URL, error) {
				return uri, nil
			}
		}
	}
	return self
}
