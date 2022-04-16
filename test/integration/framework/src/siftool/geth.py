import json
import re
from typing import Mapping, Any, Iterable
from siftool import eth, command
from siftool.common import *


def js_fmt(str, *params):
    esc_params = tuple(json.dumps(x) for x in params)
    return str.format(*esc_params)


# Documentation: https://geth.ethereum.org/docs/
# - Dev mode ("--dev") https://geth.ethereum.org/docs/getting-started/dev-mode
# - Private Network Tutorial: https://geth.ethereum.org/docs/getting-started/private-net
# - Private Networks: https://geth.ethereum.org/docs/interface/private-network
# - Running a standalone private Geth node for non-production purpose: https://medium.com/coinmonks/running-a-standalone-private-geth-node-for-non-production-purpose-d6e0ff226150

class Geth:
    def __init__(self, cmd):
        self.cmd = cmd
        self.program = "geth"

    def geth_cmd(self, network_id=None, datadir=None, ipcpath=None, ws=False, ws_addr=None, ws_port=None, ws_api=None,
        http=False, http_addr=None, http_port=None, http_api=None, rpc_allow_unprotected_txs=False, dev=False,
        dev_period=None, rpcvhosts=None, mine=False, miner_threads=None
     ):
        args = [self.program] + \
            (["--networkid", str(network_id)] if network_id else []) + \
            (["--datadir", datadir] if datadir else []) + \
            (["--ipcpath", ipcpath] if ipcpath else []) + \
            (["--ws"] if ws else []) + \
            (["--ws.addr", ws_addr] if ws_addr else []) + \
            (["--ws.port", str(ws_port)] if ws_port is not None else []) + \
            (["--ws.api", ",".join(ws_api)] if ws_api else []) + \
            (["--http"] if http else []) + \
            (["--http.addr", http_addr] if http_addr is not None else []) + \
            (["--http.port", str(http_port)] if http_port is not None else []) + \
            (["--http.api", ",".join(http_api)] if http_api else []) + \
            (["--rpc.allow-unprotected-txs"] if rpc_allow_unprotected_txs else []) + \
            (["--dev"] if dev else []) + \
            (["--dev.period", str(dev_period)] if dev_period is not None else []) + \
            (["--rpcvhosts", rpcvhosts] if rpcvhosts else []) + \
            (["--mine"] if mine else []) + \
            (["--miner.threads", str(miner_threads)] if miner_threads is not None else [])
        return args

    def geth_exec(self, geth_cmd_string, ipcpath):
        args = [self.program, "--exec", geth_cmd_string, ipcpath]
        return self.cmd.execst(args)

    class AttachEvalFunction:
        def __init__(self, geth, ipcpath):
            self.geth = geth
            self.ipcpath = ipcpath

        def __call__(self, js_expr, raw=False):
            args = [self.geth.program, "attach", "--exec", js_expr, self.ipcpath]
            res = stdout(self.geth.cmd.execst(args))
            return res if raw else json.loads(res)

        @property
        def coinbase_addr(self):
            js_expr = f"eth.coinbase"
            return self(js_expr)

        def create_account(self, password):
            js_expr = js_fmt("personal.newAccount({})", password)
            return self(js_expr)

        def unlock_account(self, addr, password):
            js_expr = js_fmt("personal.unlockAccount({}, {})", addr, password)
            # TODO Exception if unlock fails
            # Returns true if acount was unlocked successfully
            # Prints an error if not successful
            return self(js_expr)

        def get_balance(self, addr):
            js_expr = js_fmt("eth.getBalance({})", addr)
            return self(js_expr)

        # Amount is in wei
        # Returns txhash
        def send(self, from_addr, to_addr, amount):
            js_expr = js_fmt("eth.sendTransaction({{from: {}, to: {}, value: {}}})", from_addr, to_addr, amount)
            return self(js_expr)


    def attach_eval_fn(self, ipcpath):
        return Geth.AttachEvalFunction(self, ipcpath)

    # Creates a password-protected account in geth keyring for a given private key. This works deterministically,
    # meaning the account address/pubkey is the same for the same private key, and also the same that you would get
    # when creating address/pubkey in Hardhat.
    #
    # This uses "geth account import", the keys are stored in datadir/keys. The alternative is to use "geth console"
    # personal.createAccount().
    #
    # Private key has is a hex string without "0x" prefix
    # Datadir cannot be the same datadir that a running geth uses
    # See "Creating an account by importing a private key": https://geth.ethereum.org/docs/interface/managing-your-accounts
    def create_account(self, password, private_key, datadir=None):
        assert (not private_key.startswith("0x")) and (len(private_key) == 64)
        passfile = self.cmd.mktempfile()
        keyfile = self.cmd.mktempfile()
        try:
            self.cmd.write_text_file(passfile, password)
            self.cmd.write_text_file(keyfile, private_key)
            args = [self.program, "account", "import", keyfile, "--password", passfile] + \
                (["--datadir", datadir] if datadir else [])
            res = self.cmd.execst(args)
            address = "0x" + re.compile("^Address: \\{(.*)\\}$").match(exactly_one(stdout_lines(res)))[1]
            return address
        finally:
            self.cmd.rm(keyfile)
            self.cmd.rm(passfile)

    def run_dev(self, network_id, datadir=None, http_port=None, ws_port=None, ipcpath=None):
        kwargs = {}
        if http_port is not None:
            kwargs["http"] = True
            kwargs["http_port"] = http_port
            kwargs["http_addr"] = ANY_ADDR
            kwargs["http_api"] = ("personal", "eth", "net", "web3", "debug")
        if ws_port is not None:
            kwargs["ws"] = True
            kwargs["ws_addr"] = ANY_ADDR
        cmd = self.geth_cmd(network_id=network_id, datadir=datadir, ipcpath=ipcpath, **kwargs)
        res = self.cmd.popen(cmd)
        return res

    # <editor-fold>

    # Examples of usage of geth from branch 'test-integration-geth'.
    # Dev mode creates one account with a near-infinite balance (console: eth.getBalance(eth.accounts[0])).
    # Not used at the moment
    def geth_cmd__test_integration_geth_branch(self, datadir=None):
        # def geth_cmd(args: env_ethereum.EthereumInput) -> str:
        #     apis = "personal,eth,net,web3,debug"
        #     cmd = " ".join([
        #         "geth",
        #         f"--networkid {args.network_id}",
        #         f"--ipcpath {ipcpath}",
        #         f"--ws --ws.addr 0.0.0.0 --ws.port {args.ws_port} --ws.api {apis}",
        #         f"--http --http.addr 0.0.0.0 --http.port {args.http_port} --http.api {apis}",
        #         "--rpc.allow-unprotected-txs",
        #         "--dev --dev.period 1",
        #         "--rpcvhosts=*",
        #         "--mine --miner.threads=1",
        #     ])
        #     return cmd
        #
        # geth --networkid 3 --ipcpath /tmp/geth.ipc \
        #     --ws --ws.addr 0.0.0.0 --ws.port 8646 --ws.api personal,eth,net,web3,debug \
        #     --http --http.addr 0.0.0.0 --http.port 7990 --http.api personal,eth,net,web3,debug \
        #     --rpc.allow-unprotected-txs \
        #     --dev --dev.period 1 --rpcvhosts=* --mine --miner.threads=1
        return self.geth_cmd(datadir=datadir, network_id=3, ipcpath="/tmp/geth.ipc", ws=True, ws_addr=ANY_ADDR,
            ws_port=8646, http=True, http_addr=ANY_ADDR, http_port=7990, http_api=("personal", "eth", "net", "web3", "debug"),
            rpc_allow_unprotected_txs=True, dev=True, dev_period=1, mine=True, miner_threads=1)

    # </editor-fold>

    def create_genesis_config_clique(self, chain_id: int, signer_addresses: Iterable[eth.Address],
        alloc: Mapping[eth.Address, int], gas_limit: int = 8000000, difficulty: int = 1
    ) -> Mapping[str, Any]:
        # See https://geth.ethereum.org/docs/interface/private-network
        # signer_address = "7df9a875a174b3bc565e6424a0050ebc1b2d1d82"
        # alloc = {
        #     signer_address.lower()[2:]: 300000,
        #     "f41c74c9ae680c1aa78f42e5647a62f353b7bdde": 400000,
        # }
        # chain_id = 15
        extradata = "0x" + "00"*32 + ''.join([addr.lower()[2:] for addr in signer_addresses]) + "00"*65
        return {
            "config": {
                "chainId": chain_id,
                "homesteadBlock": 0,
                "eip150Block": 0,
                "eip155Block": 0,
                "eip158Block": 0,
                "byzantiumBlock": 0,
                "constantinopleBlock": 0,
                "petersburgBlock": 0,
                "clique": {
                    "period": 5,
                    "epoch": 30000
                }
            },
            "difficulty": str(difficulty),
            "gasLimit": str(gas_limit),
            "extradata": extradata,
            "alloc": {k: {"balance": str(v)} for k, v in alloc.items()}
        }

    def run_env(self, path):
        signer_addr, signer_private_key = eth.web3_create_account()
        ethereum_chain_id = 9999
        if self.cmd.exists(path):
            self.cmd.rmdir(path)
        if not self.cmd.exists(path):
            datadir = path
            self.cmd.mkdir(datadir)
            tmp_genesis_file = self.cmd.mktempfile()
            try:
                genesis = self.create_genesis_config_clique(ethereum_chain_id, [signer_addr], {signer_addr: 1000000})
                self.cmd.write_text_file(tmp_genesis_file, json.dumps(genesis))
                args = [self.program, "init", tmp_genesis_file] + \
                    (["--datadir", datadir] if datadir else [])
                # cmd = command.buildcmd(args=args)
                return self.cmd.execst(args)
            finally:
                self.cmd.rm(tmp_genesis_file)

    def init(self, ethereum_chain_id: int, signers: Iterable[eth.Address], datadir: Optional[str] = None,
        funds_alloc: Optional[Mapping[eth.Address, int]] = None
    ):
        funds_alloc = funds_alloc or {}
        tmp_genesis_file = self.cmd.mktempfile()
        try:
            genesis = self.create_genesis_config_clique(ethereum_chain_id, signers, funds_alloc)
            self.cmd.write_text_file(tmp_genesis_file, json.dumps(genesis))
            args = [self.program, "init", tmp_genesis_file] + \
                (["--datadir", datadir] if datadir else [])
            # cmd = command.buildcmd(args=args)
            res = self.cmd.execst(args)
            print(repr(res))
        finally:
            self.cmd.rm(tmp_genesis_file)

    def buid_run_args(self, datadir, network_id):
        args = [self.program, "--networkid", str(network_id), "--nodiscover"] + \
            (["--datadir", datadir] if datadir else [])
        return command.buildcmd(args)


# How Wilson is running geth:
# https://github.com/Sifchain/sifnode/commit/3e4feff2d5f707109aa609b8941f06d3cd349c92

# TODO How to mint, create initial accounts, fund them
