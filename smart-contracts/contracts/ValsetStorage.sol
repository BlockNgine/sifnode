// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.0;

contract ValsetStorage {

    /*
     * @dev: Total power of all validators
     */
    uint256 public totalPower;

    /*
     * @dev: Current valset version
     */
    uint256 public currentValsetVersion;

    /*
     * @dev: validator count
     */
    uint256 public validatorCount;

    /*
     * @dev: Keep track of active validator
     */
    mapping(address => mapping(uint256 => bool)) public validators;

    /*
     * @dev: operator address that can:
     *   Set BridgeBank's address (if it's not already set)
     *   Add new Validators, remove Validators, and update Validators' powers
     *   Call the function `recoverGas(uint256,address)`
     *   Change the operator
     */
    address public operator;

    /*
     * @dev: validator address + uint then hashed equals key mapped to powers
     */
    mapping(address => mapping(uint256 => uint256)) public powers;

    /*
    * @notice gap of storage for future upgrades
    */
    uint256[100] private ____gap;
}