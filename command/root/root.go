package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lerax-chain/lerax-chain/command/backup"
	"github.com/lerax-chain/lerax-chain/command/bridge"
	"github.com/lerax-chain/lerax-chain/command/genesis"
	"github.com/lerax-chain/lerax-chain/command/helper"
	"github.com/lerax-chain/lerax-chain/command/ibft"
	"github.com/lerax-chain/lerax-chain/command/license"
	"github.com/lerax-chain/lerax-chain/command/monitor"
	"github.com/lerax-chain/lerax-chain/command/peers"
	"github.com/lerax-chain/lerax-chain/command/polybft"
	"github.com/lerax-chain/lerax-chain/command/polybftsecrets"
	"github.com/lerax-chain/lerax-chain/command/regenesis"
	"github.com/lerax-chain/lerax-chain/command/rootchain"
	"github.com/lerax-chain/lerax-chain/command/secrets"
	"github.com/lerax-chain/lerax-chain/command/server"
	"github.com/lerax-chain/lerax-chain/command/status"
	"github.com/lerax-chain/lerax-chain/command/txpool"
	"github.com/lerax-chain/lerax-chain/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
		polybftsecrets.GetCommand(),
		polybft.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
