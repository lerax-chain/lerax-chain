package server

import (
	"github.com/lerax-chain/lerax-chain/chain"
	"github.com/lerax-chain/lerax-chain/consensus"
	consensusDev "github.com/lerax-chain/lerax-chain/consensus/dev"
	consensusDummy "github.com/lerax-chain/lerax-chain/consensus/dummy"
	consensusIBFT "github.com/lerax-chain/lerax-chain/consensus/ibft"
	consensusPolyBFT "github.com/lerax-chain/lerax-chain/consensus/polybft"
	"github.com/lerax-chain/lerax-chain/forkmanager"
	"github.com/lerax-chain/lerax-chain/secrets"
	"github.com/lerax-chain/lerax-chain/secrets/awsssm"
	"github.com/lerax-chain/lerax-chain/secrets/gcpssm"
	"github.com/lerax-chain/lerax-chain/secrets/hashicorpvault"
	"github.com/lerax-chain/lerax-chain/secrets/local"
	"github.com/lerax-chain/lerax-chain/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

var forkManagerFactory = map[ConsensusType]ForkManagerFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerFactory,
}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
