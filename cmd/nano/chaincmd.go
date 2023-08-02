package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"littleriver.cc/go-nano/cmd/utils"
	"littleriver.cc/go-nano/internal/flags"
)

var (
	initCommand = &cli.Command{
		Action:    initGenesis,
		Name:      "init",
		Usage:     "Bootstrap and initialize a new genesis block",
		ArgsUsage: "<genesisPath>",
		Flags:     flags.Merge([]cli.Flag{utils.CachePreimagesFlag}, utils.DatabasePathFlags),
		Description: `
The init command initializes a new genesis block and definition for the network.
This is a destructive action and changes the network in which you will be
participating.

It expects the genesis file as argument.`,
	}
)

func initGenesis(ctx *cli.Context) error {
	fmt.Println("initGenesis")
	return nil
}
