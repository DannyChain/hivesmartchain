package commands

import (
	"fmt"
	"net"
	"strings"

	cli "github.com/jawher/mow.cli"
	"github.com/klyed/hivesmartchain/config"
	"github.com/klyed/hivesmartchain/crypto"
)

type configOptions struct {
	configFileOpt      *string
	genesisFileOpt     *string
	validatorIndexOpt  *int
	accountIndexOpt    *int
	initAddressOpt     *string
	initPassphraseOpt  *string
	initMonikerOpt     *string
	externalAddressOpt *string
	grpcAddressOpt     *string
}

const configFileSpec = "[--config=<config file>]"

var configFileOption = cli.StringOpt{
	Name:   "c config",
	Desc:   "Use the specified Hive Smart Chain config file",
	EnvVar: "HSC_CONFIG_FILE",
}

const genesisFileSpec = "[--genesis=<genesis json file>]"

var genesisFileOption = cli.StringOpt{
	Name:   "g genesis",
	Desc:   "Use the specified genesis JSON file rather than a key in the main config, use - to read from STDIN",
	EnvVar: "HSC_GENESIS_FILE",
}

func addConfigOptions(cmd *cli.Cmd) *configOptions {
	spec := "[--moniker=<human readable moniker>] " +
		"[--index=<index of account in GenesisDoc> " +
		"|--validator=<index of validator in GenesisDoc> " +
		"|--address=<address of signing key>] " +
		"[--passphrase=<secret passphrase to unlock key>] " +
		"[--external-address=<hostname:port>] " +
		"[--grpc-address=<hostname:port>] " +
		configFileSpec + " " + genesisFileSpec

	cmd.Spec = strings.Join([]string{cmd.Spec, spec}, " ")
	return &configOptions{
		accountIndexOpt: cmd.Int(cli.IntOpt{
			Name:   "i index",
			Desc:   "Account index (in accounts list - GenesisSpec or GenesisDoc) from which to set Address",
			Value:  -1,
			EnvVar: "HSC_ACCOUNT_INDEX",
		}),

		validatorIndexOpt: cmd.Int(cli.IntOpt{
			Name:   "v validator",
			Desc:   "Validator index (in validators list - GenesisSpec or GenesisDoc) from which to set Address",
			Value:  -1,
			EnvVar: "HSC_VALIDATOR_INDEX",
		}),

		initAddressOpt: cmd.String(cli.StringOpt{
			Name:   "a address",
			Desc:   "The address of the signing key of this node",
			EnvVar: "HSC_ADDRESS",
		}),

		initPassphraseOpt: cmd.String(cli.StringOpt{
			Name:   "p passphrase",
			Desc:   "The passphrase of the signing key of this node (currently unimplemented but planned for future version of our KeyClient interface)",
			EnvVar: "HSC_PASSPHRASE",
		}),

		initMonikerOpt: cmd.String(cli.StringOpt{
			Name:   "m moniker",
			Desc:   "An optional human-readable moniker to identify this node amongst Tendermint peers in logs and status queries",
			EnvVar: "HSC_NODE_MONIKER",
		}),

		externalAddressOpt: cmd.String(cli.StringOpt{
			Name:   "x external-address",
			Desc:   "An external address or host name provided with the port that this node will broadcast over gossip in order for other nodes to connect",
			EnvVar: "HSC_EXTERNAL_ADDRESS",
		}),

		grpcAddressOpt: cmd.String(cli.StringOpt{
			Name:   "grpc-address",
			Desc:   "GRPC listen address",
			EnvVar: "HSC_GRPC_ADDRESS",
		}),

		configFileOpt: cmd.String(configFileOption),

		genesisFileOpt: cmd.String(genesisFileOption),
	}
}

func (opts *configOptions) obtainBurrowConfig() (*config.BurrowConfig, error) {
	conf, err := obtainDefaultConfig(*opts.configFileOpt, *opts.genesisFileOpt)
	if err != nil {
		return nil, err
	}
	// Which account am I?
	conf.ValidatorAddress, err = accountAddress(conf, *opts.initAddressOpt, *opts.accountIndexOpt, *opts.validatorIndexOpt)
	if err != nil {
		return nil, err
	}
	if *opts.initPassphraseOpt != "" {
		conf.Passphrase = opts.initPassphraseOpt
	}
	if *opts.initMonikerOpt == "" {

		chainIDHeader := ""
		if conf.GenesisDoc != nil && conf.GenesisDoc.ChainID != "" {
			chainIDHeader = conf.GenesisDoc.ChainID
		}
		if conf.ValidatorAddress != nil {
			// Set a default moniker... since we can at this stage of config completion and it is required for start
			conf.Tendermint.Moniker = fmt.Sprintf("%sNode_%s", chainIDHeader, conf.ValidatorAddress)
		}

		/*
		chainIDHeader := "69"
		if conf.GenesisDoc != nil && conf.GenesisDoc.ChainID() != "" {
			chainIDHeader = fmt.Sprintf("%s", chainIDHeader)
		}
		if conf.ValidatorAddress != nil {
			// Set a default moniker... since we can at this stage of config completion and it is required for start
			conf.Tendermint.Moniker = fmt.Sprintf("%s", conf.Tendermint.Moniker)
		}
		*/
	} else {
		conf.Tendermint.Moniker = *opts.initMonikerOpt
	}
	if *opts.externalAddressOpt != "" {
		conf.Tendermint.ExternalAddress = *opts.externalAddressOpt
	}
	if *opts.grpcAddressOpt != "" {
		host, port, err := net.SplitHostPort(*opts.grpcAddressOpt)
		if err != nil {
			return nil, fmt.Errorf("could not parse GRPC listen addres: %w", err)
		}
		conf.RPC.GRPC.ListenHost = host
		conf.RPC.GRPC.ListenPort = port
	}
	return conf, nil
}

// address is sourced in the following order:
// 	1. explicitly from cli
// 	2. genesis accounts (by index)
// 	3. genesis validators (by index)
// 	4. config
// 	5. genesis validator (if only one)

func accountAddress(conf *config.BurrowConfig, addressIn string, accIndex, valIndex int) (*crypto.Address, error) {
	if addressIn != "" {
		address, err := crypto.AddressFromHexString(addressIn)
		if err != nil {
			return nil, fmt.Errorf("could not read address for account in '%s': %v", addressIn, err)
		}
		return &address, nil
	} else if accIndex > -1 {
		if conf.GenesisDoc == nil {
			return nil, fmt.Errorf("unable to set Address from provided index since no " +
				"GenesisDoc/GenesisSpec provided")
		}
		if accIndex >= len(conf.GenesisDoc.Accounts) {
			return nil, fmt.Errorf("index of %v given but only %v accounts specified in GenesisDoc",
				accIndex, len(conf.GenesisDoc.Accounts))
		}
		return &conf.GenesisDoc.Accounts[accIndex].Address, nil
	} else if valIndex > -1 {
		if conf.GenesisDoc == nil {
			return nil, fmt.Errorf("unable to set Address from provided validator since no " +
				"GenesisDoc/GenesisSpec provided")
		}
		if valIndex >= len(conf.GenesisDoc.Validators) {
			return nil, fmt.Errorf("validator index of %v given but only %v validators specified in GenesisDoc",
				valIndex, len(conf.GenesisDoc.Validators))
		}
		return &conf.GenesisDoc.Validators[valIndex].Address, nil
	} else if conf.ValidatorAddress != nil {
		return conf.ValidatorAddress, nil
	} else if conf.GenesisDoc != nil && len(conf.GenesisDoc.Validators) == 1 {
		return &conf.GenesisDoc.Validators[0].Address, nil
	}
	return nil, nil
}
