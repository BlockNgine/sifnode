package types

// Ethbridge module event types
var (
	EventTypeCreateClaim              = "create_claim"
	EventTypeProphecyStatus           = "prophecy_status"
	EventTypeBurn                     = "burn"
	EventTypeLock                     = "lock"
	EventTypeUpdateWhiteListValidator = "update_whitelist_validator"
	EventTypeSetNativeToken           = "set_native_token"

	AttributeKeyEthereumSender             = "ethereum_sender"
	AttributeKeyEthereumSenderNonce        = "ethereum_sender_nonce"
	AttributeKeyCosmosReceiver             = "cosmos_receiver"
	AttributeKeyAmount                     = "amount"
	AttributeKeyNativeTokenAmount          = "native_token_amount"
	AttributeKeySymbol                     = "symbol"
	AttributeKeyCoins                      = "coins"
	AttributeKeyStatus                     = "status"
	AttributeKeyClaimType                  = "claim_type"
	AttributeKeyValidator                  = "validator"
	AttributeKeyPowerType                  = "power"
	AttributeKeyNativeTokenReceiverAccount = "native_token_receiver_account"

	AttributeKeyTokenContract        = "token_contract_address"
	AttributeKeyCosmosSender         = "cosmos_sender"
	AttributeKeyCosmosSenderSequence = "cosmos_sender_sequence"
	AttributeKeyEthereumReceiver     = "ethereum_receiver"
	AttributeKeyNetworkDescriptor    = "network_id"
	AttributeKeyNativeToken          = "native_token"
	AttributeKeyNativeTokenGas       = "native_token_gas"
	AttributeKeyMinimumLockCost      = "minimum_lock_cost"
	AttributeKeyMinimumBurnCost      = "minimum_burn_cost"

	AttributeValueCategory = ModuleName
)
