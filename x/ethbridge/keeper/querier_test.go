package keeper_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	ethbridgekeeper "github.com/Sifchain/sifnode/x/ethbridge/keeper"
	"github.com/Sifchain/sifnode/x/ethbridge/test"
	"github.com/Sifchain/sifnode/x/ethbridge/types"
	oracletypes "github.com/Sifchain/sifnode/x/oracle/types"
)

//nolint:lll
const (
	// TestResponseJSON = "{\"prophecy_id\":\"\\ufffd\\ufffd\\ufffd\\ufffdE|q\\ufffdrt\\ufffdS\\u0012D\\ufffdUj\\ufffd\\ufffd\\ufffd\\ufffdI\\ufffd\\u0018\\ufffdA9\\n \\ufffdJz\",\"status\":1,\"claim_validators\":[\"cosmosvaloper1mnfm9c7cdgqnkk66sganp78m0ydmcr4pn7fqfk\"]}"
	TestResponseJSON              = "{\"prophecy_id\":\"xy7EH/x26sNeLf62aTOQAqo3H9Fnrl19yDlub31XG5o=\",\"status\":1,\"claim_validators\":[\"cosmosvaloper1mnfm9c7cdgqnkk66sganp78m0ydmcr4pn7fqfk\"]}"
	TestCrossChainFeeResponseJSON = "{\"fee_currency\":\"ceth\",\"fee_currency_gas\":\"1\",\"minimum_lock_cost\":\"1\",\"minimum_burn_cost\":\"1\"}"
)

func TestNewQuerier(t *testing.T) {
	ctx, keeper, _, _, _, encCfg, _, _ := test.CreateTestKeepers(t, 0.7, []int64{3, 3}, "")

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := ethbridgekeeper.NewLegacyQuerier(keeper, encCfg.Amino)

	//Test wrong paths
	bz, err := querier(ctx, []string{"other"}, query)
	require.Error(t, err)
	require.Nil(t, bz)
}

func TestParseEthProphecy(t *testing.T) {
	ctx, _, _, _, oracleKeeper, encCfg, _, validatorAddresses := test.CreateTestKeepers(t, 0.7, []int64{3, 3}, "")
	valAddress := validatorAddresses[0]
	NewTestResponseJSON := strings.Replace(TestResponseJSON, "cosmosvaloper1353a4uac03etdylz86tyq9ssm3x2704j3a9n7n", valAddress.String(), -1)
	testEthereumAddress := types.NewEthereumAddress(types.TestEthereumAddress)
	testBridgeContractAddress := types.NewEthereumAddress(types.TestBridgeContractAddress)
	testTokenContractAddress := types.NewEthereumAddress(types.TestTokenContractAddress)
	networkID := oracletypes.NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM

	initialEthBridgeClaim := types.CreateTestEthClaim(
		t, testBridgeContractAddress, testTokenContractAddress, valAddress,
		testEthereumAddress, types.TestCoinsAmount, types.TestCoinsSymbol, types.ClaimType_CLAIM_TYPE_LOCK)

	_, err := oracleKeeper.ProcessClaim(ctx, networkID, initialEthBridgeClaim.GetProphecyID(), initialEthBridgeClaim.ValidatorAddress)
	require.NoError(t, err)
	testResponse := types.CreateTestQueryEthProphecyResponse(t, valAddress, types.ClaimType_CLAIM_TYPE_LOCK)
	testJSON, err := encCfg.Amino.MarshalJSON(testResponse)
	require.NoError(t, err)
	require.Equal(t, NewTestResponseJSON, string(testJSON))

	req := types.NewQueryEthProphecyRequest(initialEthBridgeClaim.GetProphecyID())
	oldID := req.ProphecyId

	bz, err2 := encCfg.Amino.MarshalJSON(req)
	require.Nil(t, err2)

	var decodedReq types.QueryEthProphecyRequest

	encCfg.Amino.MustUnmarshalJSON(bz, &decodedReq)
	newID := decodedReq.ProphecyId
	require.Equal(t, oldID, newID)
}

func TestQueryEthProphecy(t *testing.T) {
	ctx, keeper, _, _, oracleKeeper, encCfg, _, validatorAddresses := test.CreateTestKeepers(t, 0.7, []int64{3, 3}, "")
	valAddress := validatorAddresses[0]
	NewTestResponseJSON := strings.Replace(TestResponseJSON, "cosmosvaloper1353a4uac03etdylz86tyq9ssm3x2704j3a9n7n", valAddress.String(), -1)
	testEthereumAddress := types.NewEthereumAddress(types.TestEthereumAddress)
	testBridgeContractAddress := types.NewEthereumAddress(types.TestBridgeContractAddress)
	testTokenContractAddress := types.NewEthereumAddress(types.TestTokenContractAddress)
	networkID := oracletypes.NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM

	initialEthBridgeClaim := types.CreateTestEthClaim(
		t, testBridgeContractAddress, testTokenContractAddress, valAddress,
		testEthereumAddress, types.TestCoinsAmount, types.TestCoinsSymbol, types.ClaimType_CLAIM_TYPE_LOCK)

	_, err := oracleKeeper.ProcessClaim(ctx, networkID, initialEthBridgeClaim.GetProphecyID(), initialEthBridgeClaim.ValidatorAddress)
	require.NoError(t, err)
	testResponse := types.CreateTestQueryEthProphecyResponse(t, valAddress, types.ClaimType_CLAIM_TYPE_LOCK)

	//Test query String()
	testJSON, err := encCfg.Amino.MarshalJSON(testResponse)
	require.NoError(t, err)
	require.Equal(t, NewTestResponseJSON, string(testJSON))

	req := types.NewQueryEthProphecyRequest(initialEthBridgeClaim.GetProphecyID())

	bz, err2 := encCfg.Amino.MarshalJSON(req)
	require.Nil(t, err2)

	query := abci.RequestQuery{
		Path: "/custom/ethbridge/prophecies",
		Data: bz,
	}

	//Test query
	querier := ethbridgekeeper.NewLegacyQuerier(keeper, encCfg.Amino)
	res, err3 := querier(ctx, []string{types.QueryEthProphecy}, query)
	require.Nil(t, err3)

	var ethProphecyResp types.QueryEthProphecyResponse
	err4 := encCfg.Amino.UnmarshalJSON(res, &ethProphecyResp)
	require.Nil(t, err4)
	require.True(t, reflect.DeepEqual(ethProphecyResp, testResponse))

	// Test error with bad request
	query.Data = bz[:len(bz)-1]

	_, err5 := querier(ctx, []string{types.QueryEthProphecy}, query)
	require.NotNil(t, err5)

	// Test error with nonexistent request
	// badEthereumAddress := types.NewEthereumAddress("badEthereumAddress")

	bz2, err6 := encCfg.Amino.MarshalJSON(types.NewQueryEthProphecyRequest([]byte(types.TestProphecyID)))
	require.Nil(t, err6)

	query2 := abci.RequestQuery{
		Path: "/custom/oracle/prophecies",
		Data: bz2,
	}

	_, err7 := querier(ctx, []string{types.QueryEthProphecy}, query2)
	require.NotNil(t, err7)

	// Test error with empty address
	// emptyEthereumAddress := types.NewEthereumAddress("")

	bz3, err8 := encCfg.Amino.MarshalJSON(
		types.NewQueryEthProphecyRequest([]byte(types.TestProphecyID)))

	require.Nil(t, err8)

	query3 := abci.RequestQuery{
		Path: "/custom/oracle/prophecies",
		Data: bz3,
	}

	_, err9 := querier(ctx, []string{types.QueryEthProphecy}, query3)
	require.NotNil(t, err9)
}

func TestQueryCrosschainFeeConfig(t *testing.T) {
	ctx, keeper, _, _, _, encCfg, _, _ := test.CreateTestKeepers(t, 0.7, []int64{3, 3}, "")
	networkID := oracletypes.NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM

	//Test query String()
	_, err := encCfg.Amino.MarshalJSON(TestCrossChainFeeResponseJSON)
	require.NoError(t, err)
	req := types.NewQueryCrosschainFeeConfigRequest(networkID)

	bz, err2 := encCfg.Amino.MarshalJSON(req)
	require.Nil(t, err2)
	path := fmt.Sprintf("%s/%s", "/custom/ethbridge/crosschainFeeConfig", networkID.String())

	query := abci.RequestQuery{
		Path: path,
		Data: bz,
	}

	//Test query
	querier := ethbridgekeeper.NewLegacyQuerier(keeper, encCfg.Amino)
	res, err3 := querier(ctx, []string{types.QueryCrosschainFeeConfig}, query)
	require.Nil(t, err3)

	var crosschainFeeConfigResponse types.QueryCrosschainFeeConfigResponse
	err4 := encCfg.Amino.UnmarshalJSON(res, &crosschainFeeConfigResponse)
	require.Nil(t, err4)
	config, err4 := json.Marshal(crosschainFeeConfigResponse.GetCrosschainFeeConfig())
	require.Nil(t, err4)
	require.Equal(t, string(config), TestCrossChainFeeResponseJSON)
}
