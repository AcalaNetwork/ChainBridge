// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/centrifuge/go-substrate-rpc-client/v3/scale"
	substrate_utils "github.com/ChainSafe/ChainBridge/shared/substrate"
)

type TokenSymbol = uint8
type EvmAddress = [20]byte

type DEXShare struct {
	IsToken bool
	AsToken TokenSymbol
	IsERC20 bool
	AsERC20 EvmAddress
}

func (p *DEXShare) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsToken = true
		err = decoder.Decode(&p.AsToken)
	case 1:
		p.IsERC20 = true
		err = decoder.Decode(&p.AsERC20)
	}

	if err != nil {
		return err
	}

	return nil
}

func (p DEXShare) Encode(encoder scale.Encoder) error {
	var err1, err2 error

	switch {
	case p.IsToken:
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(p.AsToken)
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

type CurrencyId struct {
	IsToken 	bool
	AsToken 	TokenSymbol
	IsDEXShare	bool
	AsDEXShare	struct {
		Share_0	DEXShare
		Share_1	DEXShare
	}
	IsERC20		bool
	AsERC20		EvmAddress
	IsChainSafeResource bool
	AsChainSafeResource types.Bytes32
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
	case 3:
		p.IsChainSafeResource = true
		err = decoder.Decode(&p.AsChainSafeResource)
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
	case p.IsChainSafeResource:
		err1 = encoder.PushByte(2)
		err2 = encoder.Encode(p.AsChainSafeResource)
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
	CurrencyId      CurrencyId
	From			types.AccountID
	To				types.AccountID
	Amount          types.U128
	Topics          []types.Hash
}

type EventCurrenciesBalanceUpdated struct {
	Phase           types.Phase
	CurrencyId      CurrencyId
	Who				types.AccountID
	Amount			types.I128
	Topics          []types.Hash
}

type EventCurrenciesDeposited struct {
	Phase           types.Phase
	CurrencyId      CurrencyId
	Who				types.AccountID
	Amount			types.U128
	Topics          []types.Hash
}

type EventCurrenciesWithdrawn struct {
	Phase           types.Phase
	CurrencyId      CurrencyId
	Who				types.AccountID
	Amount			types.U128
	Topics          []types.Hash
}

type EventTokensTransferred struct {
	Phase           types.Phase
	CurrencyId      CurrencyId
	From			types.AccountID
	To				types.AccountID
	Amount			types.U128
	Topics          []types.Hash
}

type EventTokensDustLost struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId      CurrencyId
	Amount			types.U128
	Topics          []types.Hash
}

type EventChainSafeTransferRegisteredResourceId struct {
	Phase           types.Phase
	ResourceId		types.U128
	CurrencyId      CurrencyId		
	Topics          []types.Hash
}

type EventRenVmBridgeMinted struct {
	Phase           types.Phase
	Who				types.AccountID
	Amount      	types.U128		
	Topics          []types.Hash
}

type EventRenVmBridgeBurnt struct {
	Phase           types.Phase
	Who				types.AccountID
	Dest			types.Bytes
	Amount      	types.U128		
	Topics          []types.Hash
}

type EventRenVmBridgeRotatedKey struct {
	Phase           types.Phase
	PublicKey		types.H160	
	Topics          []types.Hash
}

type EventAuctionManagerNewCollateralAuction struct {
	Phase           types.Phase
	AuctionId		types.U32
	CurrencyId		CurrencyId
	Size			types.U128
	Target			types.U128
	Topics          []types.Hash
}

type EventAuctionManagerCancelAuction struct {
	Phase           types.Phase
	AuctionId		types.U32
	Topics          []types.Hash
}

type EventAuctionManagerCollateralAuctionDealt struct {
	Phase           types.Phase
	AuctionId		types.U32
	CurrencyId		CurrencyId
	Size			types.U128
	Winner			types.AccountID
	Payment			types.U128
	Topics          []types.Hash
}

type EventAuctionManagerDEXTakeCollateralAuction struct {
	Phase           types.Phase
	AuctionId		types.U32
	CurrencyId		CurrencyId
	Size			types.U128
	Amount			types.U128
	Topics          []types.Hash
}

type EventCdpEngineLiquidateUnsafeCDP struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Who				types.AccountID
	Amount			types.U128
	Debit			types.U128
	Strategy		types.U8
	Topics          []types.Hash
}

type EventCdpEngineSettleCDPInDebit struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Who				types.AccountID
	Topics          []types.Hash
}

type EventCdpEngineCloseCDPInDebitByDEX struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Who				types.AccountID
	Amount			types.U128
	Refund			types.U128
	Debit			types.U128
	Topics          []types.Hash
}

type option struct {
	hasValue bool
}

type OptionU128 struct {
	option
	value types.U128
}

func (o OptionU128) Encode(encoder scale.Encoder) error {
	return encoder.EncodeOption(o.hasValue, o.value)
}

func (o *OptionU128) Decode(decoder scale.Decoder) error {
	return decoder.DecodeOption(&o.hasValue, &o.value)
}

type EventCdpEngineInterestRatePerSec struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Change			OptionU128
	Topics          []types.Hash
}

type EventCdpEngineLiquidationRatioUpdated struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Change			OptionU128
	Topics          []types.Hash
}

type EventCdpEngineLiquidationPenaltyUpdated struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Change			OptionU128
	Topics          []types.Hash
}

type EventCdpEngineRequiredCollateralRatioUpdated struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Change			OptionU128
	Topics          []types.Hash
}

type EventCdpEngineMaximumTotalDebitValueUpdated struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Cap				types.U128
	Topics          []types.Hash
}

type EventCdpEngineGlobalInterestRatePerSecUpdated struct {
	Phase           types.Phase
	Rate			types.U128
	Topics          []types.Hash
}

type EventCdpTreasuryExpectedCollateralAuctionSizeUpdated struct {
	Phase           types.Phase
	CurrencyId		CurrencyId
	Size			types.U128
	Topics          []types.Hash
}

type EventDexAddProvision struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId_0	CurrencyId
	Contribution_0	types.U128
	CurrencyId_1	CurrencyId
	Contribution_1	types.U128
	Topics          []types.Hash
}

type EventDexAddLiquidity struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId_0	CurrencyId
	Amount_0		types.U128
	CurrencyId_1	CurrencyId
	Amount_1		types.U128
	Share			types.U128
	Topics          []types.Hash
}

type EventDexRemoveLiquidity struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId_0	CurrencyId
	Amount_0		types.U128
	CurrencyId_1	CurrencyId
	Amount_1		types.U128
	Share			types.U128
	Topics          []types.Hash
}

type EventDexSwap struct {
	Phase           types.Phase
	Who				types.AccountID
	Path			[]CurrencyId
	Supply			types.U128
	Target			types.U128
	Topics          []types.Hash
}

type EventDexEnableTradingPair struct {
	Phase           types.Phase
	TradingPair		struct{
						CurrencyId_0	CurrencyId
						CurrencyId_1	CurrencyId
					}
	Topics          []types.Hash
}

type EventDexListTradingPair struct {
	Phase           types.Phase
	TradingPair		struct{
						CurrencyId_0	CurrencyId
						CurrencyId_1	CurrencyId
					}
	Topics          []types.Hash
}

type EventDexDisableTradingPair struct {
	Phase           types.Phase
	TradingPair		struct{
						CurrencyId_0	CurrencyId
						CurrencyId_1	CurrencyId
					}
	Topics          []types.Hash
}

type EventDexProvisioningToEnabled struct {
	Phase           types.Phase
	TradingPair		struct{
						CurrencyId_0	CurrencyId
						CurrencyId_1	CurrencyId
					}
	Pool_0			types.U128
	Pool_1			types.U128
	Share			types.U128
	Topics          []types.Hash
}

type EventEmergencyShutdownShutdown struct {
	Phase           types.Phase
	BlockNumber		types.U32
	Topics          []types.Hash
}

type EventEmergencyShutdownOpenRefund struct {
	Phase           types.Phase
	BlockNumber		types.U32
	Topics          []types.Hash
}

type EventEmergencyShutdownRefund struct {
	Phase           types.Phase
	Who				types.AccountID
	Amount			types.U128
	Refund			[]struct {
						CurrencyId	CurrencyId
						Amount		types.U128
					}
	Topics          []types.Hash
}

type EventEvmAccountsClaimAccount struct {
	Phase           types.Phase
	AccountId		types.AccountID
	EvmAddress		types.H160
	Topics          []types.Hash
}

type EventHomaValidatorListModuleFreezeValidator struct {
	Phase           types.Phase
	Validator		types.AccountID
	Topics          []types.Hash
}

type EventHomaValidatorListModuleThawValidator struct {
	Phase           types.Phase
	Validator		types.AccountID
	Topics          []types.Hash
}

type EventHomaValidatorListModuleBondGuarantee struct {
	Phase           types.Phase
	Who				types.AccountID
	Validator		types.AccountID
	Amount			types.U128
	Topics          []types.Hash
}

type EventHomaValidatorListModuleUnbondGuarantee struct {
	Phase           types.Phase
	Who				types.AccountID
	Validator		types.AccountID
	Amount			types.U128
	Topics          []types.Hash
}

type EventHomaValidatorListModuleWithdrawnGuarantee struct {
	Phase           types.Phase
	Who				types.AccountID
	Validator		types.AccountID
	Amount			types.U128
	Topics          []types.Hash
}

type EventHomaValidatorListModuleSlashGuarantee struct {
	Phase           types.Phase
	Who				types.AccountID
	Validator		types.AccountID
	Amount			types.U128
	Topics          []types.Hash
}

type EventHonzonAuthorization struct {
	Phase           types.Phase
	Authorizer		types.AccountID
	Authorizee		types.AccountID
	CurrencyId		CurrencyId
	Topics          []types.Hash
}

type EventHonzonUnAuthorization struct {
	Phase           types.Phase
	Authorizer		types.AccountID
	Authorizee		types.AccountID
	CurrencyId		CurrencyId
	Topics          []types.Hash
}

type EventHonzonUnAuthorizationAll struct {
	Phase           types.Phase
	Authorizer		types.AccountID
	Topics          []types.Hash
}

type PoolId struct {
	IsLoansIncentive	bool
	AsLoansIncentive 	CurrencyId
	IsDexIncentive		bool
	AsDexIncentive		CurrencyId
	IsHomaIncentive		bool
	IsDexSaving 		bool
	AsDexSaving			CurrencyId
	IsHomaValidatorAllowance	bool
	AsHomaValidatorAllowance	types.AccountID
}

func (p *PoolId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		p.IsLoansIncentive = true
		err = decoder.Decode(&p.AsLoansIncentive)
	case 1:
		p.IsDexIncentive = true
		err = decoder.Decode(&p.AsDexIncentive)
	case 2:
		p.IsHomaIncentive = true
	case 3:
		p.IsDexSaving = true
		err = decoder.Decode(&p.AsDexSaving)
	case 4:
		p.IsHomaValidatorAllowance = true
		err = decoder.Decode(&p.AsHomaValidatorAllowance)
	}

	if err != nil {
		return err
	}

	return nil
}

func (p PoolId) Encode(encoder scale.Encoder) error {
	var err1, err2 error

	switch {
	case p.IsLoansIncentive:
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(p.AsLoansIncentive)
	case p.IsDexIncentive:
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(p.AsDexIncentive)
	case p.IsHomaIncentive:
		err1 = encoder.PushByte(2)
	case p.IsDexSaving:
		err1 = encoder.PushByte(3)
		err2 = encoder.Encode(p.AsDexSaving)
	case p.IsHomaValidatorAllowance:
		err1 = encoder.PushByte(4)
		err2 = encoder.Encode(p.AsHomaValidatorAllowance)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type EventIncentivesDepositDexShare struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId		CurrencyId
	Amount			types.U128
	Topics          []types.Hash
}

type EventIncentivesWithdrawDexShare struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId		CurrencyId
	Amount			types.U128
	Topics          []types.Hash
}

type EventIncentivesClaimRewards struct {
	Phase           types.Phase
	Who				types.AccountID
	PoolId			PoolId
	Topics          []types.Hash
}

type EventLoansPositionUpdated struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId		CurrencyId
	CollateralAdjustment	types.I128
	DebitAdjustment	types.I128
	Topics          []types.Hash
}

type EventLoansConfiscateCollateralAndDebit struct {
	Phase           types.Phase
	Who				types.AccountID
	CurrencyId		CurrencyId
	CollateralAmount	types.I128
	Debit			types.I128
	Topics          []types.Hash
}

type EventLoansTransferLoan struct {
	Phase           types.Phase
	Sender			types.AccountID
	Receiver		types.AccountID
	CurrencyId		CurrencyId
	Topics          []types.Hash
}

type EventNFTCreatedClass struct {
	Phase           types.Phase
	Who				types.AccountID
	ClassId			types.U32
	Topics          []types.Hash
}

type EventNFTMintedToken struct {
	Phase           types.Phase
	Minter			types.AccountID
	Receiver		types.AccountID
	ClassId			types.U32
	Quantity		types.U32
	Topics          []types.Hash
}

type EventNFTTransferredToken struct {
	Phase           types.Phase
	From			types.AccountID
	To				types.AccountID
	ClassId			types.U32
	TokenId			types.U64
	Topics          []types.Hash
}

type EventNFTBurnedToken struct {
	Phase           types.Phase
	Who				types.AccountID
	ClassId			types.U32
	TokenId			types.U64
	Topics          []types.Hash
}

type EventNFTDestroyedClass struct {
	Phase           types.Phase
	Owner			types.AccountID
	ClassId			types.U32
	Topics          []types.Hash
}


type Events struct {
	substrate_utils.Events
	
	AcalaTreasury_Deposit					[]types.EventTreasuryDeposit	//nolint:stylecheck,golint

	Currencies_Transferred					[]EventCurrenciesTransferred	//nolint:stylecheck,golint
	Currencies_BalanceUpdated 				[]EventCurrenciesBalanceUpdated	//nolint:stylecheck,golint
	Currencies_Deposited 					[]EventCurrenciesDeposited		//nolint:stylecheck,golint
	Currencies_Withdrawn 					[]EventCurrenciesWithdrawn		//nolint:stylecheck,golint

	Tokens_Transferred						[]EventTokensTransferred		//nolint:stylecheck,golint
	Tokens_DustLost							[]EventTokensDustLost			//nolint:stylecheck,golint

	ChainSafeTransfer_RegisteredResourceId	[]EventChainSafeTransferRegisteredResourceId	//nolint:stylecheck,golint

	RenVmBridge_Minted						[]EventRenVmBridgeMinted		//nolint:stylecheck,golint
	RenVmBridge_Burnt						[]EventRenVmBridgeBurnt			//nolint:stylecheck,golint
	RenVmBridge_RotatedKey					[]EventRenVmBridgeRotatedKey	//nolint:stylecheck,golint

	AuctionManager_NewCollateralAuction		[]EventAuctionManagerNewCollateralAuction		//nolint:stylecheck,golint
	AuctionManager_CancelAuction			[]EventAuctionManagerCancelAuction
	AuctionManager_CollateralAuctionDealt	[]EventAuctionManagerCollateralAuctionDealt		//nolint:stylecheck,golint
	AuctionManager_DEXTakeCollateralAuction	[]EventAuctionManagerDEXTakeCollateralAuction	//nolint:stylecheck,golint

	CdpEngine_LiquidateUnsafeCDP			[]EventCdpEngineLiquidateUnsafeCDP				//nolint:stylecheck,golint
	CdpEngine_SettleCDPInDebit				[]EventCdpEngineSettleCDPInDebit				//nolint:stylecheck,golint
	CdpEngine_CloseCDPInDebitByDEX			[]EventCdpEngineCloseCDPInDebitByDEX			//nolint:stylecheck,golint
	CdpEngine_InterestRatePerSec			[]EventCdpEngineInterestRatePerSec				//nolint:stylecheck,golint
	CdpEngine_LiquidationRatioUpdated		[]EventCdpEngineLiquidationRatioUpdated			//nolint:stylecheck,golint
	CdpEngine_LiquidationPenaltyUpdated		[]EventCdpEngineLiquidationPenaltyUpdated		//nolint:stylecheck,golint
	CdpEngine_RequiredCollateralRatioUpdated	[]EventCdpEngineRequiredCollateralRatioUpdated	//nolint:stylecheck,golint
	CdpEngine_MaximumTotalDebitValueUpdated	[]EventCdpEngineMaximumTotalDebitValueUpdated	//nolint:stylecheck,golint
	CdpEngine_GlobalInterestRatePerSecUpdated	[]EventCdpEngineGlobalInterestRatePerSecUpdated	//nolint:stylecheck,golint

	CdpTreasury_ExpectedCollateralAuctionSizeUpdated	[]EventCdpTreasuryExpectedCollateralAuctionSizeUpdated	//nolint:stylecheck,golint

	Dex_AddProvision						[]EventDexAddProvision				//nolint:stylecheck,golint
	Dex_AddLiquidity						[]EventDexAddLiquidity				//nolint:stylecheck,golint
	Dex_RemoveLiquidity						[]EventDexRemoveLiquidity			//nolint:stylecheck,golint
	Dex_Swap								[]EventDexSwap						//nolint:stylecheck,golint
	Dex_EnableTradingPair					[]EventDexEnableTradingPair			//nolint:stylecheck,golint
	Dex_ListTradingPair						[]EventDexListTradingPair			//nolint:stylecheck,golint
	Dex_DisableTradingPair					[]EventDexDisableTradingPair		//nolint:stylecheck,golint
	Dex_ProvisioningToEnabled				[]EventDexProvisioningToEnabled		//nolint:stylecheck,golint

	EmergencyShutdown_Shutdown				[]EventEmergencyShutdownShutdown	//nolint:stylecheck,golint
	EmergencyShutdown_OpenRefund			[]EventEmergencyShutdownOpenRefund	//nolint:stylecheck,golint
	EmergencyShutdown_Refund				[]EventEmergencyShutdownRefund		//nolint:stylecheck,golint

	EvmAccounts_ClaimAccount				[]EventEvmAccountsClaimAccount		//nolint:stylecheck,golint

	HomaValidatorListModule_FreezeValidator		[]EventHomaValidatorListModuleFreezeValidator		//nolint:stylecheck,golint
	HomaValidatorListModule_ThawValidator		[]EventHomaValidatorListModuleThawValidator			//nolint:stylecheck,golint
	HomaValidatorListModule_BondGuarantee		[]EventHomaValidatorListModuleBondGuarantee			//nolint:stylecheck,golint
	HomaValidatorListModule_UnbondGuarantee		[]EventHomaValidatorListModuleUnbondGuarantee		//nolint:stylecheck,golint
	HomaValidatorListModule_WithdrawnGuarantee	[]EventHomaValidatorListModuleWithdrawnGuarantee	//nolint:stylecheck,golint
	HomaValidatorListModule_SlashGuarantee		[]EventHomaValidatorListModuleSlashGuarantee		//nolint:stylecheck,golint

	Honzon_Authorization					[]EventHonzonAuthorization			//nolint:stylecheck,golint
	Honzon_UnAuthorization					[]EventHonzonUnAuthorization		//nolint:stylecheck,golint
	Honzon_UnAuthorizationAll				[]EventHonzonUnAuthorizationAll		//nolint:stylecheck,golint

	Incentives_DepositDexShare				[]EventIncentivesDepositDexShare	//nolint:stylecheck,golint
	Incentives_WithdrawDexShare				[]EventIncentivesWithdrawDexShare	//nolint:stylecheck,golint
	Incentives_ClaimRewards					[]EventIncentivesClaimRewards		//nolint:stylecheck,golint

	Loans_PositionUpdated					[]EventLoansPositionUpdated					//nolint:stylecheck,golint
	Loans_ConfiscateCollateralAndDebit		[]EventLoansConfiscateCollateralAndDebit	//nolint:stylecheck,golint
	Loans_TransferLoan						[]EventLoansTransferLoan					//nolint:stylecheck,golint

	NFT_CreatedClass						[]EventNFTCreatedClass				//nolint:stylecheck,golint
	NFT_MintedToken							[]EventNFTMintedToken				//nolint:stylecheck,golint
	NFT_TransferredToken					[]EventNFTTransferredToken			//nolint:stylecheck,golint
	NFT_BurnedToken							[]EventNFTBurnedToken				//nolint:stylecheck,golint
	NFT_DestroyedClass						[]EventNFTDestroyedClass			//nolint:stylecheck,golint
}

