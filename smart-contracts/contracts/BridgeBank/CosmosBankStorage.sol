// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.0;

/**
 * @title Cosmos Bank Storage
 * @dev Stores Cosmos deposits, nonces, networkDescriptor
 */
contract CosmosBankStorage {

    /**
    * @notice Cosmos deposit struct
    */
    struct CosmosDeposit {
        bytes cosmosSender;
        address payable ethereumRecipient;
        address bridgeTokenAddress;
        uint256 amount;
        bool locked;
    }

    /**
    * @notice number of bridge tokens
    */
    uint256 public bridgeTokenCount;

    /**
    * @notice cosmos deposit nonce
    */
    uint256 public cosmosDepositNonce;

    /*
    * @notice [DEPRECATED] mapping of symbols to token addresses
    */
    mapping(string => address) private controlledBridgeTokens;

    /*
    * @notice [DEPRECATED] mapping of lowercase symbols to properly capitalized tokens
    */
    mapping(string => string) private lowerToUpperTokens;

    /**
    * @notice network descriptor
    */
    uint256 public networkDescriptor;

    /*
    * @notice gap of storage for future upgrades
    */
    uint256[99] private ____gap;
}
