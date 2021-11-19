const { setup } = require("./helpers/testFixture");
const { upgrades } = require("hardhat");
const { use, expect } = require("chai");

const web3 = require("web3");
const BigNumber = web3.BigNumber;

require("chai").use(require("chai-as-promised")).use(require("chai-bignumber")(BigNumber)).should();

describe("CosmosBridge Upgrade", function () {
  const consensusThreshold = 70;
  let userOne;
  let userTwo;
  let userThree;
  let userFour;
  let accounts;
  let signerAccounts;
  let operator;
  let owner;
  let initialPowers;
  let initialValidators;
  let networkDescriptor;
  let state;
  let pauser;
  let MockCosmosBridgeUpgrade;

  before(async function () {
    accounts = await ethers.getSigners();

    signerAccounts = accounts.map((e) => {
      return e.address;
    });

    operator = accounts[0];
    userOne = accounts[1];
    userTwo = accounts[2];
    userFour = accounts[3];
    userThree = accounts[7];

    owner = accounts[5];
    pauser = accounts[6];

    initialPowers = [25, 25, 25, 25];
    initialValidators = signerAccounts.slice(0, 4);
    networkDescriptor = 1;
    MockCosmosBridgeUpgrade = await ethers.getContractFactory("MockCosmosBridgeUpgrade");
    state = {};
  });

  describe("CosmosBridge smart contract deployment", function () {
    beforeEach(async function () {
      state = await setup({
        initialValidators: [userOne.address, userTwo.address, userThree.address, userFour.address],
        initialPowers: [30, 20, 21, 29],
        operator,
        consensusThreshold,
        owner,
        user: userOne,
        recipient: userThree,
        pauser,
        networkDescriptor,
      });

      state.cosmosBridge = await upgrades.upgradeProxy(
        state.cosmosBridge.address,
        MockCosmosBridgeUpgrade
      );
    });

    it("should be able to mint tokens for a user", async function () {
      const amount = 100000000000;
      state.cosmosBridge.should.exist;

      await state.cosmosBridge.connect(operator).tokenFaucet();
      const operatorBalance = await state.cosmosBridge.balanceOf(operator.address);
      Number(operatorBalance).should.be.bignumber.equal(amount);
    });

    it("should be able to transfer tokens from the operator", async function () {
      const startingOperatorBalance = await state.cosmosBridge.balanceOf(operator.address);
      Number(startingOperatorBalance).should.be.bignumber.equal(0);

      const amount = 100000000000;
      state.cosmosBridge.should.exist;

      await state.cosmosBridge.connect(operator).tokenFaucet();
      await state.cosmosBridge.connect(operator).transfer(userOne.address, amount);

      const operatorBalance = await state.cosmosBridge.balanceOf(operator.address);
      const userOneBalance = await state.cosmosBridge.balanceOf(userOne.address);

      Number(operatorBalance).should.be.bignumber.equal(0);
      Number(userOneBalance).should.be.bignumber.equal(amount);
    });

    it("should not be able to initialize cosmos bridge a second time", async function () {
      state.cosmosBridge.should.exist;

      await expect(
        state.cosmosBridge.initialize(
          userFour.address,
          50,
          state.initialValidators,
          state.initialPowers,
          state.networkDescriptor
        )
      ).to.be.revertedWith("Initialized");
    });

    describe("Storage Remains Intact", function () {
      it("should not allow the operator to update the Bridge Bank once it has been set", async function () {
        await expect(
          state.cosmosBridge.connect(operator).setBridgeBank(state.bridgeBank.address)
        ).to.be.revertedWith("The Bridge Bank cannot be updated once it has been set");
      });
    });
  });
});
