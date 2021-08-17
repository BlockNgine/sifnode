// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.0;

import "./Oracle.sol";
import "./BridgeBank/BridgeBank.sol";
import "./CosmosBridgeStorage.sol";
import "hardhat/console.sol";

contract CosmosBridge is CosmosBridgeStorage, Oracle {
    bool private _initialized;
    uint256[100] private ___gap;

    /*
     * @dev: Event declarations
     */
    event LogBridgeBankSet(address bridgeBank);

    event LogNewProphecyClaim(
        uint256 indexed prophecyID,
        address indexed ethereumReceiver,
        uint256 indexed amount
    );

    event LogNewBridgeTokenCreated(
        uint8 decimals,
        uint256 indexed sourcechainId,
        string name,
        string symbol,
        address indexed sourceContractAddress,
        address indexed bridgeTokenAddress
    );

    event LogProphecyCompleted(uint256 prophecyID, bool success);

    /*
     * @dev: Modifier to restrict access to current ValSet validators
     */
    modifier onlyValidator {
        require(
            isActiveValidator(msg.sender),
            "Must be an active validator"
        );
        _;
    }

    /*
     * @dev: Constructor
     */
    function initialize(
        address _operator,
        uint256 _consensusThreshold,
        address[] calldata _initValidators,
        uint256[] calldata _initPowers,
        uint256 _networkDescriptor
    ) external {
        require(!_initialized, "Initialized");

        operator = _operator;
        networkDescriptor = _networkDescriptor;
        hasBridgeBank = false;
        _initialized = true;
        Oracle._initialize(
            _operator,
            _consensusThreshold,
            _initValidators,
            _initPowers
        );
    }

    function changeOperator(address _newOperator) external onlyOperator {
        require(_newOperator != address(0), "invalid address");
        operator = _newOperator;
    }

    /*
     * @dev: setBridgeBank
     */
    function setBridgeBank(address payable _bridgeBank) external onlyOperator {
        require(
            !hasBridgeBank,
            "The Bridge Bank cannot be updated once it has been set"
        );

        hasBridgeBank = true;
        bridgeBank = _bridgeBank;

        emit LogBridgeBankSet(bridgeBank);
    }

    function getProphecyID(
        bytes calldata cosmosSender,
        uint256 cosmosSenderSequence,
        address payable ethereumReceiver,
        address tokenAddress,
        uint256 amount,
        bool doublePeg,
        uint128 nonce,
        uint256 _networkDescriptor
    ) public pure returns (uint256) {
        return uint256(
            keccak256(
                abi.encode(
                    cosmosSender,
                    cosmosSenderSequence,
                    ethereumReceiver,
                    tokenAddress,
                    amount,
                    doublePeg,
                    nonce,
                    _networkDescriptor
                )
            )
        );
    }

    function verifySignature(
        address signer,
        bytes32 hashDigest,
        uint8 _v,
		    bytes32 _r,
		    bytes32 _s
    ) private pure returns (bool) {
		    bytes32 messageDigest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hashDigest));
		    return signer == ecrecover(messageDigest, _v, _r, _s);
	}
    
    function getSignedPowerAndFindDup(
        SignatureData[] calldata _validators,
        bytes32 hashDigest
    ) private view returns(bool, uint256 pow) {
        for (uint256 i = 0; i < _validators.length; i++) {
            SignatureData memory validator = _validators[i];

            require(isActiveValidator(validator.signer), "INV_SIGNER");

            require(verifySignature(
                    validator.signer,
                    hashDigest,
                    validator._v,
                    validator._r,
                    validator._s
                ),
                "INV_SIG"
            );

            pow += getValidatorPower(validator.signer);

            for (uint256 j = i + 1; j < _validators.length; j++) {
                if (validator.signer == _validators[j].signer) {
                    return (true, pow);
                }
            }
        }
    }

    struct SignatureData {
        address signer;
        uint8 _v;
		    bytes32 _r;
		    bytes32 _s;
    }

    struct ClaimData {
        bytes cosmosSender;
        uint256 cosmosSenderSequence;
        address payable ethereumReceiver;
        address tokenAddress;
        uint256 amount;
        uint256 networkDescriptor;
        string tokenName;
        string tokenSymbol;
        uint8 tokenDecimals;
        bool doublePeg;
        uint128 nonce;
    }

    function batchSubmitProphecyClaimAggregatedSigs(
        bytes32[] calldata sigs,
        ClaimData[] calldata claims,
        SignatureData[][] calldata signatureData
    ) external {
        require(sigs.length == claims.length, "INV_CLM_LEN");
        require(sigs.length == signatureData.length, "INV_SIG_LEN");

        uint256 intermediateNonce = lastNonceSubmitted;
        require(intermediateNonce + 1 == claims[0].nonce, "INV_ORD");

        for (uint256 i = 0; i < sigs.length; i++) {
            intermediateNonce++;
            _submitProphecyClaimAggregatedSigs(sigs[i], claims[i], signatureData[i]);
        }

        lastNonceSubmitted = intermediateNonce;
    }

    function submitProphecyClaimAggregatedSigs(
        bytes32 hashDigest,
        ClaimData calldata claimData,
        SignatureData[] calldata signatureData
    ) external {
        uint256 previousNonce = lastNonceSubmitted;
        require(previousNonce + 1 == claimData.nonce, "INV_ORD");

        _submitProphecyClaimAggregatedSigs(hashDigest, claimData, signatureData);

        // update the nonce
        lastNonceSubmitted = claimData.nonce;
    }

    function _submitProphecyClaimAggregatedSigs(
        bytes32 hashDigest,
        ClaimData calldata claimData,
        SignatureData[] calldata signatureData
    ) private {
        uint256 prophecyID = getProphecyID(
            claimData.cosmosSender,
            claimData.cosmosSenderSequence,
            claimData.ethereumReceiver,
            claimData.tokenAddress,
            claimData.amount,
            claimData.doublePeg,
            claimData.nonce,
            claimData.networkDescriptor
        );
        
        require(
            uint256(hashDigest) == prophecyID,
            "INV_DATA"
        );

        // ensure signature lengths are correct
        require(
            signatureData.length > 0 && signatureData.length <= validatorCount,
            "INV_SIG_LEN"
        );

        // ensure the networkDescriptor matches
        if (!claimData.doublePeg) {
            require(_verifyNetworkDescriptor(claimData.networkDescriptor), "INV_NET_DESC");
        }
        
        (bool dup, uint256 pow) = getSignedPowerAndFindDup(signatureData, hashDigest);
        require(!dup, "DUP_SIGNER");

        require(getProphecyStatus(pow), "INV_POW");

        address tokenAddress;
        if(claimData.doublePeg) {
            if(!_isManagedToken(claimData.tokenAddress)) {
              // if we are double pegging AND we don't control the token, we deploy a new smart contract
              tokenAddress = _createNewBridgeToken(
                  claimData.tokenSymbol,
                  claimData.tokenName,
                  claimData.tokenAddress,
                  claimData.tokenDecimals,
                  claimData.networkDescriptor
              );
            } else {
              // if we are double pegging and already control the token, then we are going to need to get the address on this chain
              tokenAddress = sourceAddressToDestinationAddress[claimData.tokenAddress];
            }
        } else {
            tokenAddress = claimData.tokenAddress;
        }

        completeProphecyClaim(
            prophecyID,
            claimData.ethereumReceiver,
            tokenAddress,
            claimData.amount
        );

        emit LogNewProphecyClaim(
            prophecyID,
            claimData.ethereumReceiver,
            claimData.amount
        );
    }

    function _isManagedToken(address tokenAddress) private returns(bool) {
        return sourceAddressToDestinationAddress[tokenAddress] != address(0);
    }

    function _verifyNetworkDescriptor(uint256 _networkDescriptor) private returns(bool) {
        return _networkDescriptor == networkDescriptor;
    }

    function _createNewBridgeToken(
        string calldata symbol,
        string calldata name,
        address sourceChainTokenAddress,
        uint8 decimals,
        uint256 _networkDescriptor
    ) internal returns(address tokenAddress) {
        require(
            sourceAddressToDestinationAddress[sourceChainTokenAddress] == address(0),
            "INV_SRC_ADDR"
        );
        // need to make a business decision on what this symbol should be
        // First lock of this asset, deploy new contract and get new symbol/token address
        address tokenAddress = BridgeBank(bridgeBank)
            .createNewBridgeToken(
                name,
                symbol,
                decimals
            );

        sourceAddressToDestinationAddress[sourceChainTokenAddress] = tokenAddress;

        emit LogNewBridgeTokenCreated(
            decimals,
            _networkDescriptor,
            name,
            symbol,
            sourceChainTokenAddress,
            tokenAddress
        );
    }

    /*
     * @dev: completeProphecyClaim
     *       Allows for the completion of ProphecyClaims once processed by the Oracle.
     *       Burn claims unlock tokens stored by BridgeBank.
     *       Lock claims mint BridgeTokens on BridgeBank's token whitelist.
     */
    function completeProphecyClaim(
        uint256 prophecyID,
        address payable ethereumReceiver,
        address tokenAddress,
        uint256 amount
    ) internal {
        (bool success, ) = bridgeBank.call{gas: 120000}(
            abi.encodeWithSignature(
                "handleUnpeg(address,address,uint256)",
                ethereumReceiver,
                tokenAddress,
                amount
            )
        );

        // prophecy completed and whether or not the call to bridgebank was successful
        emit LogProphecyCompleted(prophecyID, success);
    }
}
