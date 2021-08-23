// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.0;

import "./CosmosBankStorage.sol";
import "./EthereumBankStorage.sol";
import "./CosmosWhiteListStorage.sol";

/**
 * @title Bank Storage
 * @dev Stores addresses for owner, operator, and CosmosBridge
 **/
contract BankStorage is 
    CosmosBankStorage,
    EthereumBankStorage,
    CosmosWhiteListStorage {

    /**
    * @notice operator address that can:
    *   Reinitialize BridgeBank
    *   Update Eth whitelist
    *   Change the operator
    */
    address public operator;

    /*
    * @notice [DEPRECATED] address of the Oracle smart contract
    */
    address private oracle;

    /**
    * @notice address of the Cosmos Bridge smart contract
    */
    address public cosmosBridge;

    /**
    * @notice owner address that can use the admin API
    */
    address public owner;

    /*
    * @notice [DEPRECATED] token limit
    */
    mapping (string => uint256) private maxTokenAmount;

    /*
    * @notice gap of storage for future upgrades
    */
    uint256[100] private ____gap;
}