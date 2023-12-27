<h1 align="center">
  <img src="https://github.com/moshaoli688/clash/raw/master/docs/logo.png" alt="Clash" width="200">
  <br>Clash<br>
</h1>

<h4 align="center">A rule-based tunnel in Go.</h4>

<p align="center">
  <a href="https://github.com/moshaoli688/clash/actions">
    <img src="https://img.shields.io/github/actions/workflow/status/moshaoli688/clash/release.yml?branch=master&style=flat-square" alt="Github Actions">
  </a>
  <a href="https://goreportcard.com/report/github.com/moshaoli688/clash">
    <img src="https://goreportcard.com/badge/github.com/moshaoli688/clash?style=flat-square">
  </a>
  <img src="https://img.shields.io/github/go-mod/go-version/moshaoli688/clash?style=flat-square">
  <a href="https://github.com/moshaoli688/clash/releases">
    <img src="https://img.shields.io/github/release/moshaoli688/clash/all.svg?style=flat-square">
  </a>
</p>

## Features

使用dreamacro的源码编译的，且用且珍惜吧，由于没有premium的源码，所以去premium镜像下载，地址（包含客户端镜像）：https://www.clash.la/releases/ 

This is a general overview of the features that comes with Clash.  

- Inbound: HTTP, HTTPS, SOCKS5 server, TUN device
- Outbound: Shadowsocks(R), VMess, Trojan, Snell, SOCKS5, HTTP(S), Wireguard
- Rule-based Routing: dynamic scripting, domain, IP addresses, process name and more
- Fake-IP DNS: minimises impact on DNS pollution and improves network performance
- Transparent Proxy: Redirect TCP and TProxy TCP/UDP with automatic route table/rule management
- Proxy Groups: automatic fallback, load balancing or latency testing
- Remote Providers: load remote proxy lists dynamically
- RESTful API: update configuration in-place via a comprehensive API

*Some of the features may only be available in the [Premium core](https://dreamacro.github.io/clash/premium/introduction.html).*

## Documentation

You can find the latest documentation at [https://dreamacro.github.io/clash/](https://dreamacro.github.io/clash/).
文档没弄，有空折腾

## Credits

- [riobard/go-shadowsocks2](https://github.com/riobard/go-shadowsocks2)
- [v2ray/v2ray-core](https://github.com/v2ray/v2ray-core)
- [WireGuard/wireguard-go](https://github.com/WireGuard/wireguard-go)

## License

This software is released under the GPL-3.0 license.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmoshaoli688%2Fclash.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmoshaoli688%2Fclash?ref=badge_large)
