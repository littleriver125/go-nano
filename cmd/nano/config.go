package main

import (
	"littleriver.cc/go-nano/nano/node/proto"
)

type Config struct {
	Addr      string        `json:"addr"`
	AddrRPC   string        `json:"addr_rpc"`
	AddrPprof string        `json:"addr_pprof"`
	Peers     []string      `json:"peers"`
	Network   proto.Network `json:"network"`
}

//func makeFullNode(ctx *cli.Context) (*node.Node, ethapi.Backend) {
//
//}
