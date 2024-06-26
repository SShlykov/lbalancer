package urls

import (
	"net/http/httputil"
	"net/url"
)

type Urls struct {
	Req  int
	Urls []*httputil.ReverseProxy
}

func New(links []string) *Urls {
	return &Urls{
		Req:  0,
		Urls: strtoproxies(links),
	}
}

// Get - proxy Round Robbin
func (ur *Urls) Get() *httputil.ReverseProxy {
	ur.Req++
	return ur.Urls[ur.Req%len(ur.Urls)]
}

func strtoproxies(links []string) []*httputil.ReverseProxy {
	proxies := make([]*httputil.ReverseProxy, len(links))
	for i, link := range links {
		uri, _ := url.Parse(link)
		proxies[i] = httputil.NewSingleHostReverseProxy(uri)
	}

	return proxies
}
