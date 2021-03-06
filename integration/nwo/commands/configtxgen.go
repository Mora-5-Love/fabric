/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package commands

type OutputBlock struct {
	ChannelID   string
	Profile     string
	ConfigPath  string
	OutputBlock string
}

func (o OutputBlock) SessionName() string {
	return "configtxgen-output-block"
}

func (o OutputBlock) Args() []string {
	return []string{
		"-channelID", o.ChannelID,
		"-profile", o.Profile,
		"-configPath", o.ConfigPath,
		"-outputBlock", o.OutputBlock,
	}
}

type CreateChannelTx struct {
	ChannelID             string
	Profile               string
	ConfigPath            string
	OutputCreateChannelTx string
}

func (c CreateChannelTx) SessionName() string {
	return "configtxgen-create-channel-tx"
}

func (c CreateChannelTx) Args() []string {
	return []string{
		"-channelID", c.ChannelID,
		"-profile", c.Profile,
		"-configPath", c.ConfigPath,
		"-outputCreateChannelTx", c.OutputCreateChannelTx,
	}
}
