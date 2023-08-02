package utils

import (
	"github.com/urfave/cli/v2"
	"littleriver.cc/go-nano/internal/flags"
)

var (
	// TestnetFlags is the flag group of all built-in supported testnets.
	TestnetFlags = []cli.Flag{
		//GoerliFlag,
		//SepoliaFlag,
	}
	// NetworkFlags is the flag group of all built-in supported networks.
	NetworkFlags = append([]cli.Flag{MainnetFlag}, TestnetFlags...)

	MainnetFlag = &cli.BoolFlag{
		Name:     "mainnet",
		Usage:    "Ethereum mainnet",
		Category: flags.EthCategory,
	}

	// DatabasePathFlags is the flag group of all database path flags.
	DatabasePathFlags = []cli.Flag{
		//DataDirFlag,
		//AncientFlag,
		//RemoteDBFlag,
		//HttpHeaderFlag,
	}

	CachePreimagesFlag = &cli.BoolFlag{
		Name:     "cache.preimages",
		Usage:    "Enable recording the SHA3/keccak preimages of trie keys",
		Category: flags.PerfCategory,
	}
)
