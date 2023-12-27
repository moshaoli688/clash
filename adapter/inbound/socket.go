package inbound

import (
	"net"
	"net/netip"

	C "github.com/moshaoli688/clash/constant"
	"github.com/moshaoli688/clash/context"
	"github.com/moshaoli688/clash/transport/socks5"
)

// NewSocket receive TCP inbound and return ConnContext
func NewSocket(target socks5.Addr, conn net.Conn, source C.Type) *context.ConnContext {
	metadata := parseSocksAddr(target)
	metadata.NetWork = C.TCP
	metadata.Type = source
	if ip, port, err := parseAddr(conn.RemoteAddr()); err == nil {
		metadata.SrcIP = ip
		metadata.SrcPort = C.Port(port)
	}
	if addrPort, err := netip.ParseAddrPort(conn.LocalAddr().String()); err == nil {
		metadata.OriginDst = addrPort
	}
	return context.NewConnContext(conn, metadata)
}
