// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = BridgePalletName + ".add_relayer"
var SetResourceMethod Method = BridgePalletName + ".set_resource"
var SetThresholdMethod Method = BridgePalletName + ".set_threshold"
var WhitelistChainMethod Method = BridgePalletName + ".whitelist_chain"
var ChainSafeTransferToBridgeMethod Method = "ChainSafeTransfer.transfer_to_bridge"
var ChainSafeTransferNativeToBridgeMethod Method = "ChainSafeTransfer.transfer_native_to_bridge"
var ChainSafeRegisterResourceIdMethod Method = "ChainSafeTransfer.register_resource_id"
var ChainSafeTransferFromBridgeMethod Method = "ChainSafeTransfer.transfer_from_bridge"
var SudoMethod Method = "Sudo.sudo"
