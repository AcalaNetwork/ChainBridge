// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/centrifuge/go-substrate-rpc-client/v2/scale"
	substrate_utils "github.com/ChainSafe/ChainBridge/shared/substrate"
)

type Nmsl struct {
	IsHzy 	bool
	IsYxy 	bool
}

func (m *Nmsl) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	if b == 0 {
		m.IsHzy = true
	} else if b == 1 {
		m.IsYxy = true
	}

	return nil
}

type Hello struct {
	IsNothing bool
	IsNumber bool
	AsNumber types.U32
}

func (p *Hello) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsNothing = true
	case 1:
		p.IsNumber = true
		err = decoder.Decode(&p.AsNumber)
	}

	if err != nil {
		return err
	}

	return nil
}

type TokenSymbol = uint8

type CurrencyId struct {
	IsToken 	bool
	AsToken 	TokenSymbol
	IsDEXShare	bool
	AsDEXShare	struct {
		Symbol_0	TokenSymbol
		Symbol_1	TokenSymbol
	}
	IsERC20		bool
	AsERC20		[20]byte
}

func (p *CurrencyId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsToken = true
		err = decoder.Decode(&p.AsToken)
	case 1:
		p.IsDEXShare = true
		err = decoder.Decode(&p.AsDEXShare)
	case 2:
		p.IsERC20 = true
		err = decoder.Decode(&p.AsERC20)
	}

	if err != nil {
		return err
	}

	return nil
}

func (p CurrencyId) Encode(encoder scale.Encoder) error {
	var err1, err2 error

	switch {
	case p.IsToken:
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(p.AsToken)
	case p.IsDEXShare:
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(p.AsDEXShare)
	case p.IsERC20:
		err1 = encoder.PushByte(2)
		err2 = encoder.Encode(p.AsERC20)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type EventCurrenciesTransferred struct {
	Phase           types.Phase
	Currency      	CurrencyId
	From			types.AccountID
	To				types.AccountID
	Amount          types.U128
	Topics          []types.Hash
}

type EventCdpEngineGlobalStabilityFeeUpdated struct {
	Phase           types.Phase
	Rate      		types.U128
	Topics          []types.Hash
}

type EventCdpTreasuryCollateralAuctionMaximumSizeUpdated struct {
	Phase           types.Phase
	Currency      	CurrencyId
	Topics          []types.Hash
}

type EventCdpTreasuryTest1 struct {
	Phase           types.Phase
	Symbol      	uint8
	Topics          []types.Hash
}

type EventCdpTreasuryTest2 struct {
	Phase           types.Phase
	Who      		Nmsl
	Topics          []types.Hash
}

type EventCdpTreasuryTest3 struct {
	Phase           types.Phase
	NewHello      	Hello
	Topics          []types.Hash
}

type Events struct {
	substrate_utils.Events
	Currencies_Transferred								[]EventCurrenciesTransferred	//nolint:stylecheck,golint
	AcalaTreasury_Deposit								[]types.EventTreasuryDeposit	//nolint:stylecheck,golint
	CdpEngine_GlobalStabilityFeeUpdated					[]EventCdpEngineGlobalStabilityFeeUpdated	//nolint:stylecheck,golint
	CdpTreasury_CollateralAuctionMaximumSizeUpdated		[]EventCdpTreasuryCollateralAuctionMaximumSizeUpdated	//nolint:stylecheck,golint
	CdpTreasury_Test1 []EventCdpTreasuryTest1	//nolint:stylecheck,golint
	CdpTreasury_Test2 []EventCdpTreasuryTest2	//nolint:stylecheck,golint
	CdpTreasury_Test3 []EventCdpTreasuryTest3	//nolint:stylecheck,golint
}
