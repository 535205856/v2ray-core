package socks

import (
	"github.com/v2ray/v2ray-core/app"
	"github.com/v2ray/v2ray-core/proxy/common/connhandler"
	"github.com/v2ray/v2ray-core/proxy/socks/config/json"
)

type SocksServerFactory struct {
}

func (factory SocksServerFactory) Create(dispatcher app.PacketDispatcher, rawConfig interface{}) (connhandler.InboundConnectionHandler, error) {
	config := rawConfig.(*json.SocksConfig)
	config.Initialize()
	return NewSocksServer(dispatcher, config), nil
}

func init() {
	connhandler.RegisterInboundConnectionHandlerFactory("socks", SocksServerFactory{})
}
