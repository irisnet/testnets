# Nyancat Testnet

## What is IRISnet Nyancat Testnet

Nyancat Testnet is an IRISHub Testnet operate as a stable application testnet which has the same version as the mainnet, so that the service providers of IRISnet can develop their apps on or interact with IRIShub without running a node or lcd instance. 

## Public Endpoints

- GRPC: 34.80.202.172:9090
- RPC: http://34.80.202.172:26657
- REST: http://34.80.202.172:1317/swagger/

## Running a Node

If you'd like to setup a test node yourself instead of using the public endpoints, it'll be the same as joining [IRIS Hub Mainnet](https://stage.irisnet.org/docs/get-started/mainnet.html), except that:

### Genesis File

[Download](https://github.com/irisnet/testnets/raw/master/nyancat/config/genesis.json)

### Peers

Add the following `seeds` and `persistent_peers` in the `config.toml`:

Seeds:

```bash
07e58f179b2b7101b72f04248f542f67af8993bd@35.234.10.84:26656
```

Persistent Peers:

```bash
bc77e49df0de4d70ab6f97f1e3a17bfb51a1ea7a@34.80.202.172:26656
```

## Faucet

Welcome to ask for the test tokens in our [testnet channels](https://discord.gg/9cSt7MX2fn)

## Explorer

<https://nyancat.irisplorer.io>

