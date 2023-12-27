package outboundgroup

import (
	"time"

	C "github.com/moshaoli688/clash/constant"
	"github.com/moshaoli688/clash/constant/provider"
)

const (
	defaultGetProxiesDuration = time.Second * 5
)

func touchProviders(providers []provider.ProxyProvider) {
	for _, provider := range providers {
		provider.Touch()
	}
}

func getProvidersProxies(providers []provider.ProxyProvider, touch bool) []C.Proxy {
	proxies := []C.Proxy{}
	for _, provider := range providers {
		if touch {
			provider.Touch()
		}
		proxies = append(proxies, provider.Proxies()...)
	}
	return proxies
}
