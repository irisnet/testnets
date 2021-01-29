# IRISHub Simulated Upgrade

## IRISHub v1.0 as a Significant Milestone
IRIS Hub v1.0 will be an **incompatible** upgrade including many important features: it will merge the IBC-integrated Cosmos Stargate's latest version (Cosmos SDK v0.4x and Tendermint v0.34.x) and will also integrate and upgrade many unique features of IRISHub e.g. Coinswap(AMM), NFT, iService v2, oracle, etc.

- New features brought by [Cosmos SDK v0.4x](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.40.1) and [Tendermint v0.34.x](https://github.com/tendermint/tendermint/blob/v0.34.3/UPGRADING.md)
  - [Inter-Blockchain Communication (IBC)– cross-chain transactions](https://figment.network/resources/cosmos-stargate-upgrade-overview/#ibc)
  - [Protobuf Migration – blockchain performance & dev acceleration](https://figment.network/resources/cosmos-stargate-upgrade-overview/#proto)
  - [State Sync – minutes to sync new nodes](https://figment.network/resources/cosmos-stargate-upgrade-overview/#sync)
  - [Full-Featured Light Clients](https://figment.network/resources/cosmos-stargate-upgrade-overview/#light)
  - [Chain Upgrade Module – upgrade automation](https://figment.network/resources/cosmos-stargate-upgrade-overview/#upgrade)

- IRISHub's unique features:
   - [Coinswap](https://bifrost.irisnet.org/docs/features/coinswap.html): is the AMM (automated market maker) module implementing the protocol similar to Uniswap on IRIS Hub
   - [NFT](https://bifrost.irisnet.org/docs/features/nft.html): Non-fungible Token module supporting assets to be tokenized
   - [iService v2](https://bifrost.irisnet.org/docs/features/service.html): Interchain Service module which is refactored and optimized
   - [oracle](https://bifrost.irisnet.org/docs/features/oracle.html): oracle module based on the upgraded iService module to implement the oracle feature that can provide decentralized off-chain data onto IRIS Hub

## The instruction of IRISHub v1.0's incompatibility

As one of the earliest mainnet in Cosmos ecosystem, IRISHub have used a very early version of Cosmos SDK which is v0.28.

Cosmos-SDK v0.4x adds lots of features which are incompatible with v0.2x. To integrate the revolutionary **IBC protocol**, IRISHub v1.0 must use Cosmos-SDK v0.4x which is incompatible with the current IRISHub mainnet.

Therefore, IRISHub will have the first hard fork upgrade since the mainnet launch. Because this upgrade is of great significance and is different from the previous smooth upgrade process, in order to ensure the success upgrade of the mainnet, the IRISnet team strongly recommends that all active mainnet validators take participate in this simulated upgrade. 

The IRIS Foundation is planning to reward each validator 1000 iris for successfully participating in this simulated upgrade.

## Simulated Upgrade Process

The simulated mainnet upgrade will use mainnet data to highly simulate the mainnet upgrade process:

1. The IRISnet team exports genesis.json based on the specified height of the current mainnet
2. The IRISnet team uniformly generates and replaces the validator public keys in `genesis.json` exported by the current mainnet, and publishes the `genesis.json`
3. The IRISnet team executes genesis migrate, start the seed node
4. The mainnet validators who intends to participate in the simulated upgrade will receive the test private key from IRISnet team
5. The mainnet validators execute genesis migrate, and start the validator nodes
6. Validators with enough voting power (67%) are expected to pariticpate the simulated upgrade before **2021-02-01 15:00:00 UTC**. If not, IRISnet will use the test private keys to start several validators who didn't participate.
7. After the simulated upgrade testnet is successfully launched, the validators who has valid pre-vote in the first block can use the Operator account(which the account to create the mainnet validator) to send `1iris` to our test address `iaa1eyyusfj0yv7hyr7z9e5z2u35c5hnf8lt2r7jfd` on this testnet to prove the identity of your mainnet validator. **This is required if you are willing to get the reward.**

### Steps

1. [Click here to register](http://nyancat-irisnet.mikecrm.com/I7ZzepJ), only mainnet active validators are accepted

2. [Download the genesis.json](TBD) of current mainnet, which is exported by the IRISnet team from height `8,800,000` and replaced the validator's public keys for testing

3. Install irishub v1.0.0-rc1

   ```bash
   git clone https://github.com/irisnet/irishub
   cd irishub && git checkout v1.0.0-rc1
   make install
   ```

4. Migrate genesis file

   Migrate the exported genesis.json to irishub v1.0

   ```bash
   iris migrate genesis.json --chain-id irishub-sim-upgrade > genesis_v1.0.0.json
   ```

5. Verify sha256sum

   ```bash
   sha256sum genesis_v1.0.0.json
   SHA256(genesis.json)= a28f5dff3295e4943d0eb63a1ca94de8bae1592ceceb5cb0644fe5ead3049776
   ```

6. Initialize new node

   Initialize the new node

   ```bash
   iris init [moniker] --home [v1.0.0_node_home]
   ```

7. Migrate privkey file

   Migrate private key file to irishub v1.0
   
   ```bash
   go run migrate/scripts/privValUpgrade.go <test_priv_validator.json> [v1.0.0_node_home]/config/priv_validator_key.json [v1.0.0_node_home]/data/priv_validator_state.json
   ```

8. Migrate node key file

   Migrate node key file to irishub v1.0

    ```bash
    cp [v0.16.x_node_home]/config/node_key.json [v1.0.0_node_home]/config/node_key.json
    ```

9. Copy migrated genesis file

   Copy genesis_v1.0.0.json to new node home

    ```bash
    cp genesis_v1.0.0.json [v1.0.0_node_home]/config/genesis.json
    ```

10. Configure new node

    Configure `minimum-gas-prices` in `[v1.0.0_node_home]/config/app.toml`

    ```toml
    # The minimum gas prices a validator is willing to accept for processing a
    # transaction. A transaction's fees must meet the minimum of any denomination
    # specified in this config (e.g. 0.25token1;0.0001token2).
    minimum-gas-prices = "0.2uiris"
    ```

    Add `seeds` to `[v1.0.0_node_home]/config/config.toml`

    ```toml
    # Comma separated list of seed nodes to connect to
    seeds = "d6393a68366377f4fc8152fe9e6ca3ccf635dcbe@35.221.220.135:26656"
    ```

11. Start new node

    Start new node with irishub v1.0.0-rc1

    ```bash
    iris unsafe-reset-all --home [v1.0.0_node_home]
    iris start --home [v1.0.0_node_home]
    ```


