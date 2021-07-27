package txs

import (
	"errors"
	"math/big"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	abci "github.com/tendermint/tendermint/abci/types"
	"go.uber.org/zap"

	"github.com/Sifchain/sifnode/cmd/ebrelayer/types"
	oracletypes "github.com/Sifchain/sifnode/x/oracle/types"
)

// ProphecyCompletedEventToProphecyInfo parses data from a prophecy completed event witnessed on Cosmos
func ProphecyCompletedEventToProphecyInfo(claimType types.Event, attributes []abci.EventAttribute, sugaredLogger *zap.SugaredLogger) (types.ProphecyInfo, error) {
	var prophecyID []byte
	var cosmosSender []byte
	var cosmosSenderSequence *big.Int
	var ethereumReceiver common.Address
	var symbol string
	var amount big.Int
	var networkDescriptor uint32

	attributeNumber := 0

	for _, attribute := range attributes {
		key := string(attribute.GetKey())
		val := string(attribute.GetValue())

		// Set variable based on the attribute's key
		switch key {
		case types.CosmosSender.String():
			cosmosSender = []byte(val)
			attributeNumber++
		case types.CosmosSenderSequence.String():
			attributeNumber++
			tempSequence := new(big.Int)
			tempSequence, ok := tempSequence.SetString(val, 10)
			if !ok {
				// log.Println("Invalid account sequence:", val)
				sugaredLogger.Errorw("Invalid account sequence", "account sequence", val)
				return types.ProphecyInfo{}, errors.New("invalid account sequence: " + val)
			}
			cosmosSenderSequence = tempSequence
		case types.EthereumReceiver.String():
			attributeNumber++
			if !common.IsHexAddress(val) {
				// log.Printf("Invalid recipient address: %v", val)
				sugaredLogger.Errorw("Invalid recipient address", "recipient address", val)

				return types.ProphecyInfo{}, errors.New("invalid recipient address: " + val)
			}
			ethereumReceiver = common.HexToAddress(val)
		case types.Symbol.String():
			attributeNumber++
			if claimType == types.MsgBurn {
				if !strings.Contains(val, defaultSifchainPrefix) {
					// log.Printf("Can only relay burns of '%v' prefixed coins", defaultSifchainPrefix)
					sugaredLogger.Errorw("only relay burns prefixed coins", "coin symbol", val)
					return types.ProphecyInfo{}, errors.New("can only relay burns of '%v' prefixed coins" + defaultSifchainPrefix)
				}
				res := strings.SplitAfter(val, defaultSifchainPrefix)
				symbol = strings.Join(res[1:], "")
			} else {
				symbol = val
			}
		case types.Amount.String():
			attributeNumber++
			tempAmount, ok := sdk.NewIntFromString(val)
			if !ok {
				// log.Println("Invalid amount:", val)
				sugaredLogger.Errorw("Invalid amount", "amount", val)

				return types.ProphecyInfo{}, errors.New("invalid amount:" + val)
			}
			amount = *big.NewInt(tempAmount.Int64())
		case types.NetworkDescriptor.String():
			attributeNumber++
			tempNetworkDescriptor, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				sugaredLogger.Errorw("network id can't parse", "networkDescriptor", val)
				return types.ProphecyInfo{}, errors.New("network id can't parse")
			}
			networkDescriptor = uint32(tempNetworkDescriptor)

			// check if the networkDescriptor is valid
			if !oracletypes.NetworkDescriptor(networkDescriptor).IsValid() {
				return types.ProphecyInfo{}, errors.New("network id is invalid")
			}
		}
	}

	if attributeNumber < 6 {
		sugaredLogger.Errorw("message not complete", "attributeNumber", attributeNumber)
		return types.ProphecyInfo{}, errors.New("message not complete")
	}
	return types.ProphecyInfo{
		ProphecyID:           prophecyID,
		NetworkDescriptor:    oracletypes.NetworkDescriptor(networkDescriptor),
		CosmosSender:         string(cosmosSender),
		CosmosSenderSequence: cosmosSenderSequence.Uint64(),
		EthereumReceiver:     ethereumReceiver.String(),
		TokenSymbol:          symbol,
		TokenAmount:          amount,
		DoublePeg:            false,
		GlobalNonce:          0,
		EthereumAddress:      []string{},
		Signatures:           []string{},
	}, nil
}
