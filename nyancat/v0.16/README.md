# Nyancat Testnet V0.16

To support `DeFi` functionalities, we introduced the [HTLC](https://stage.irisnet.org/docs/features/htlc.html) and [Coin Swap](https://stage.irisnet.org/docs/features/coinswap.html) modules in v0.16, so that investors can transfer their tokens across chains and trade them decentralized.

## Latest version: [v0.16.0-rc0](https://github.com/irisnet/irishub/releases/tag/v0.16.0-rc0)

**Note**:

- The point-to-`iris` conversion rate will be announced at the end of this incentivized testnet campaign.

- For every task, only the top 40 completed validators can get rewards (the upgrade task is based on the voting list of the upgrade proposal, if a valid voter doesn't complete upgrade in time, the reward will be extended to the next completed validator)

## v0.16.0-rc0

As what we did on the testnets for v0.15, we will first test the new introduced functions of v0.16 on a new testnet `nyancat-5`, if everything works fine, then test the upgrade on `nyancat-6`.

### Coin Swap Tasks

Coinswap relies on dynamic data to calculate exchange rates in real time, making it difficult to initiate transactions using the command line. So the coinswap module only has the [Restful API](#TODO). We have a demo web tool to help you understand and complete the tasks: [TODO](#TODO).

If you have a Ledger hardware, it will be easy to do the tasks with the web tool. Otherwise you can only use the [Restful API](#TODO) to generate unsigned transactions and [sign them offline](https://stage.irisnet.org/docs/cli-client/tx.html#flags), then [broadcast](https://stage.irisnet.org/docs/cli-client/tx.html#iriscli-tx-broadcast).

Read more about [IRISHub Coin Swap](https://stage.irisnet.org/docs/features/coinswap.html)

**Note:** Please use your testnet validator's operator account (which account you used to create the validator) to sign the transactions, so that we can verify the task completion.

| NO. | Name             | Details                                  | Points |
| --- | ---------------- | ---------------------------------------- | ------ |
| 1   | Add Liquidity    | Deposit liquidity to the reserve pool    | 10     |
| 2   | Remove Liquidity | Withdraw liquidity from the reserve pool | 10     |
| 3   | IRIS-Token Swap  | Buy or sell some BTC by IRIS             | 10     |
| 4   | Token-Token Swap | Buy or sell some ETH by BTC              | 10     |

### Snapshot Tasks

Since the block data gets bigger and bigger, making it more and more difficult to backup or to sync from scratch.

From v0.16, some of the IRISHub nodes (such as validators, sentries, seeds) will no longer need to retain all historical data.

- You can choose to drop the data of a node which before a certain block height to free up disk space.
- You can quickly start up a new node without syncing all historical data.

Read more about [IRISHub Snapshot](https://stage.irisnet.org/docs/daemon/snapshot.html)

### HTLC Tasks

[HTLC](https://stage.irisnet.org/docs/features/htlc.html) is mostly used in cross-chain atomic swaps. This task is gonna be very interesting:

| NO. | Name                | Details                                 | Points |
| --- | ------------------- | --------------------------------------- | ------ |
| *   | Get testnet rewards | Get testnet rewards on Mainnet via HTLC | 20     |

This is a special task, it can only be done **AFTER** the Mainnet upgrade. You will need to do as the following steps:

**Note:** Please use your testnet validator's operator account (which account you used to create the testnet validator) to sign the testnet transactions, so that we can verify the task completion.

- Contact @Yelong in the community, make sure he is online (His time zone is UTC+8), send him your `irishub` address for receiving the rewards
- You [Create an HTLC](https://stage.irisnet.org/docs/cli-client/htlc.html#iriscli-htlc-create) on `nyancat-6`, specifing: `--to="TBD"`
- Send @Yelong your HTLC tx hash, wait for him to create another HTLC on `irishub` with your hash-lock
- Confirm your rewards and [Claim](https://stage.irisnet.org/docs/cli-client/htlc.html#iriscli-htlc-claim) on `irishub`
- [Check status](https://stage.irisnet.org/docs/cli-client/htlc.html#iriscli-htlc-query-htlc)
