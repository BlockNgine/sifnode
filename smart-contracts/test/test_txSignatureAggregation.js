const {
  multiTokenSetup,
  getValidClaim
} = require('./helpers/testFixture');

const web3 = require("web3");
const { expect } = require('chai');
const BigNumber = web3.BigNumber;

require("chai")
  .use(require("chai-as-promised"))
  .use(require("chai-bignumber")(BigNumber))
  .should();

describe("Gas Cost Tests", function () {
  let userOne;
  let userTwo;
  let userThree;
  let userFour;
  let accounts;
  let operator;
  let owner;
  let pauser;

  // Consensus threshold of 70%
  const consensusThreshold = 70;
  let initialPowers;
  let initialValidators;
  let networkDescriptor;
  let state;

  before(async function() {
    accounts = await ethers.getSigners();
    
    operator = accounts[0];
    userOne = accounts[1];
    userTwo = accounts[2];
    userThree = accounts[3];
    userFour = accounts[4];

    owner = accounts[5];
    pauser = accounts[6].address;

    initialPowers = [25, 25, 25, 25];
    initialValidators = [
      userOne.address,
      userTwo.address,
      userThree.address,
      userFour.address
    ];

    networkDescriptor = 1;
  });

  beforeEach(async function () {
    // Deploy Valset contract
    state = await multiTokenSetup(
      initialValidators,
      initialPowers,
      operator,
      consensusThreshold,
      owner,
      userOne,
      userThree,
      pauser,
      networkDescriptor
    );

    // Add the token into white list
    await state.bridgeBank.connect(operator)
      .updateEthWhiteList(state.token1.address, true)
      .should.be.fulfilled;

    // Lock tokens on contract
    await state.bridgeBank.connect(userOne).lock(
      state.sender,
      state.token1.address,
      state.amount
    ).should.be.fulfilled;
  });

  describe("Unlock Gas Cost With 4 Validators", function () {
    it("should allow us to check the cost of submitting a prophecy claim lock", async function () {
      let balance = Number(await state.token1.balanceOf(state.recipient.address));
      expect(balance).to.be.equal(0);

      // Last nonce should now be 0
      let lastNonceSubmitted = Number(await state.cosmosBridge.lastNonceSubmitted());
      expect(lastNonceSubmitted).to.be.equal(0);

      state.nonce = 1;

      const { digest, claimData, signatures } = await getValidClaim({
        state,
        sender: state.sender,
        senderSequence: state.senderSequence,
        recipientAddress: state.recipient.address,
        tokenAddress: state.token1.address,
        amount: state.amount,
        isDoublePeg: false,
        nonce: state.nonce,
        networkDescriptor: state.networkDescriptor,
        tokenName: state.name,
        tokenSymbol: state.symbol,
        tokenDecimals: state.decimals,
        validators: accounts.slice(1, 5),
      });

      const tx = await state.cosmosBridge
        .connect(userOne)
        .submitProphecyClaimAggregatedSigs(
            digest,
            claimData,
            signatures
        );
      const receipt = await tx.wait();

      const sum = Number(receipt.gasUsed);
      console.log("~~~~~~~~~~~~\nTotal: ", sum);

      // Bridge claim should be completed
      // Last nonce should now be 1
      lastNonceSubmitted = Number(await state.cosmosBridge.lastNonceSubmitted());
      expect(lastNonceSubmitted).to.be.equal(1);

      balance = Number(await state.token1.balanceOf(state.recipient.address));
      expect(balance).to.be.equal(state.amount);
    });

    it("should allow us to check the cost of submitting a prophecy claim mint", async function () {
      let balance = Number(await state.rowan.balanceOf(state.recipient.address));
      expect(balance).to.be.equal(0);

      // Last nonce should now be 0
      let lastNonceSubmitted = Number(await state.cosmosBridge.lastNonceSubmitted());
      expect(lastNonceSubmitted).to.be.equal(0);

      state.nonce = 1;

      const { digest, claimData, signatures } = await getValidClaim({
        state,
        sender: state.sender,
        senderSequence: state.senderSequence,
        recipientAddress: state.recipient.address,
        tokenAddress: state.rowan.address,
        amount: state.amount,
        isDoublePeg: false,
        nonce: state.nonce,
        networkDescriptor: state.networkDescriptor,
        tokenName: state.name,
        tokenSymbol: state.symbol,
        tokenDecimals: state.decimals,
        validators: accounts.slice(1, 5),
      });

      const tx = await state.cosmosBridge
        .connect(userOne)
        .submitProphecyClaimAggregatedSigs(digest, claimData, signatures);
      const receipt = await tx.wait();

      const sum = Number(receipt.gasUsed);
      console.log("~~~~~~~~~~~~\nTotal: ", sum);

      // Last nonce should now be 1
      lastNonceSubmitted = Number(await state.cosmosBridge.lastNonceSubmitted());
      expect(lastNonceSubmitted).to.be.equal(1);

      // balance should have increased
      balance = Number(await state.rowan.balanceOf(state.recipient.address));
      expect(balance).to.be.equal(state.amount);
    });

    it("should allow us to check the cost of creating a new BridgeToken", async function () {
      state.nonce = 1;

      const { digest, claimData, signatures } = await getValidClaim({
        state,
        sender: state.sender,
        senderSequence: state.senderSequence,
        recipientAddress: state.recipient.address,
        tokenAddress: state.token1.address,
        amount: state.amount,
        isDoublePeg: true,
        nonce: state.nonce,
        networkDescriptor: state.networkDescriptor,
        tokenName: state.name,
        tokenSymbol: state.symbol,
        tokenDecimals: state.decimals,
        validators: accounts.slice(1, 5),
      });

      const expectedAddress = ethers.utils.getContractAddress({ from: state.bridgeBank.address, nonce: 1 });

      const tx = await state.cosmosBridge
        .connect(userOne)
        .submitProphecyClaimAggregatedSigs(
            digest,
            claimData,
            signatures
        );

      const receipt = await tx.wait();
      console.log("~~~~~~~~~~~~\nTotal: ", Number(receipt.gasUsed));

      const newlyCreatedTokenAddress = await state.cosmosBridge.sourceAddressToDestinationAddress(state.token1.address);
      expect(newlyCreatedTokenAddress).to.be.equal(expectedAddress);
    });
  });
});

/**
 * 
 * 
Unlock Gas Cost With 4 Validators
tx0  173990
~~~~~~~~~~~~
Total:  173990

Mint Gas Cost With 4 Validators
tx0  179737
~~~~~~~~~~~~
Total:  179737

Create new BridgeToken Gas Cost With 4 Validators
tx0  1162781
~~~~~~~~~~~~
Total:  1162781
 * 
 */