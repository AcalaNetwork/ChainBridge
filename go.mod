module github.com/ChainSafe/ChainBridge

go 1.15

require (
	github.com/ChainSafe/chainbridge-substrate-events v0.0.0-20200715141113-87198532025e
	github.com/ChainSafe/chainbridge-utils v1.0.6
	github.com/ChainSafe/log15 v1.0.0
	github.com/centrifuge/go-substrate-rpc-client/v3 v3.0.0
	github.com/ethereum/go-ethereum v1.10.2
	github.com/prometheus/client_golang v1.4.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/ChainSafe/chainbridge-substrate-events v0.0.0-20200715141113-87198532025e => github.com/wangjj9219/chainbridge-substrate-events v0.0.0-20210421142230-2efb6d1066fe

replace github.com/ChainSafe/chainbridge-utils v1.0.6 => github.com/AcalaNetwork/chainbridge-utils v1.0.7-0.20210422040608-a57d12461efd
