package polybft

import (
	"github.com/lerax-chain/lerax-chain/command/rootchain/registration"
	"github.com/lerax-chain/lerax-chain/command/rootchain/staking"
	"github.com/lerax-chain/lerax-chain/command/rootchain/supernet"
	"github.com/lerax-chain/lerax-chain/command/rootchain/supernet/stakemanager"
	"github.com/lerax-chain/lerax-chain/command/rootchain/validators"
	"github.com/lerax-chain/lerax-chain/command/rootchain/whitelist"
	"github.com/lerax-chain/lerax-chain/command/rootchain/withdraw"
	"github.com/lerax-chain/lerax-chain/command/sidechain/rewards"
	"github.com/lerax-chain/lerax-chain/command/sidechain/unstaking"
	sidechainWithdraw "github.com/lerax-chain/lerax-chain/command/sidechain/withdraw"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	polybftCmd := &cobra.Command{
		Use:   "polybft",
		Short: "Polybft command",
	}

	polybftCmd.AddCommand(
		// sidechain (validator set) command to unstake on child chain
		unstaking.GetCommand(),
		// sidechain (validator set) command to withdraw stake on child chain
		sidechainWithdraw.GetCommand(),
		// sidechain (reward pool) command to withdraw pending rewards
		rewards.GetCommand(),
		// rootchain (stake manager) command to withdraw stake
		withdraw.GetCommand(),
		// rootchain (supernet manager) command that queries validator info
		validators.GetCommand(),
		// rootchain (supernet manager) whitelist validator
		whitelist.GetCommand(),
		// rootchain (supernet manager) register validator
		registration.GetCommand(),
		// rootchain (stake manager) stake command
		staking.GetCommand(),
		// rootchain (supernet manager) command for finalizing genesis
		// validator set and enabling staking
		supernet.GetCommand(),
		// rootchain command for deploying stake manager
		stakemanager.GetCommand(),
	)

	return polybftCmd
}
