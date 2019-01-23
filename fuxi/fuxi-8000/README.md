# How to Finish A Task In Fuxi Incentivized Testnet?

## What is IRISnet incentivized testnet ?

IRIS foundation will offer rewards of IRIS token to early testnet participants according to the points they earned. They could get their rewards after IRISHub is launched. IRISnet development team will publish several task for each iteration of Fuxi testnet. Everyone is welcomed, and you have to fill in this [form](http://cn.mikecrm.com/H9aoXak) to complete your sign-up process. 

You need to use keybase to generate your own [pgp fingerprint](https://github.com/irisnet/testnets/blob/master/fuxi/How%20to%20use%20keybase.md) first. 

## Value Propositio of Fuxi-8000
Fuxi-8000 will be the last two testnet before launch of IRISnet Betanet. In Fuxi-8000, it's designed to simulate some real world senario on on-chain software upgrade and transaction load.
The two main features to be tested are:
* On-chain software upgrade workflow of different types
* Simulation of massive trasaction load for a short period, 500 tps at peak

## Tasks for Fuxi-8000

Fuxi-8000 testnet wil be composed of several phases. We expect to execute **three** types of software upgrade. So, stay tuned!
In  **first phase**, we bootstrap the network with iris version 0.10.0 in a decentralized way. After that, we expect the network to be running for one day.  You are supposed to finish the following tasks in phase one :

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Participate in Genesis file generation process | Submit your gen-tx.json and use `name-pgp-fingerprint` as validator's name | Submit url of your PR                                        | 100    |
| 2    | Setup a Full node                              | set up a full node and use `name-pgp-fingerprint` as `monkier` | Submit your IP and the team could visit the port of your node | 100    |
| 3    | Delegate some **iris** to a validator          | Get some **iris** from faucet and delegate to a validator    | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |
| 4    | Withdraw some **iris** from reward pool        | After you become a validator, you could withdraw your rewards by `iriscli distribution withdraw-rewards` | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |

Then, we go to **phase two** of this testnet. In this phase, a smooth software upgrade  to version v0.10.2 is planed. This upgrade will be fully backward compatible. Read more here. You are supposed to finish the following tasks in phase two :

| No   | Name                                                         | Details                                                      | Criteria                                                     | Points |
| ---- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Vote for `software upgrade` proposal in Phase two            | Complete an `vote `transaction                               | Run `vote` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction. Refer to this [doc](https://www.irisnet.org/docs/cli-client/gov/vote.html) to learn how to vote | 50     |
| 2    | Participate in Software upgrade process. You need to Upgrade your sofware to version `v0.10.2` once the specific software upgrade proposal is passed in the future. | Participate in software-upgrade                              | Submit your IP and the team could visit the port of your node | 200    |
| 3    | When running `v0.10.2`, delegate some **iris** to a validator | Get some **iris** from faucet and delegate to a validator    | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |
| 4    | When running `v0.10.2`, withdraw some **iris** from reward pool | After you become a validator, you could withdraw your rewards by `iriscli distribution withdraw-rewards` | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |

If software upgrade is passed, then we will enter **phase three** of testnet. We will test the software upgrade workflow to recover from consensus failure:Patch softwre upgrade. In `v0.10.2`, there is a easter egg. The team will trigger a consensus failure deliberately. This bug will cause the network to halt and you need to switch to v0.10.2-patch to come back live. You are supposed to finish the following tasks in phase three:

| No   | Name                                                         | Details                                                      | Criteria                                                     | Points |
| ---- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Switch to `v0.10.2-patch`                                   | Install the new version of iris and start with flag `--replay-last-block` | Submit your IP and the team could visit the port of your node | 200    |
| 2    | When the network come back and your node use`v0.10.2-patch`, delegate some **iris** to a validator | Get some **iris** from faucet and delegate to a validator    | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |
| 3    | When the network come back and your node use `v0.10.2-patch`, withdraw some **iris** from reward pool | After you become a validator, you could withdraw your rewards by `iriscli distribution withdraw-rewards` | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |

If +2/3 validators have use the patch and do a restart in last phase , then we will enter **phase four** of testnet. In this phase we will try to export a snapshop of blockchain and use it as genesis file for a new chain. This is another type of  software upgarade: Restart Software Upgrade . In this phase, we will first submit a system-halt proposal. If this proposal is passed, Tendermint will stop at certain height. Then, you can export the snapshot or download the new genesis file provided by IRIS team and start with new version `v0.11.0`. You are supposed to finish the following tasks in phase four:

| No   | Name                                                         | Details                                                      | Criteria                                                     | Points |
| ---- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Vote for `system-halt` proposal in Phase four                 | Complete an `vote `transaction                               | Run `vote` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction. Refer to this [doc](https://www.irisnet.org/docs/cli-client/gov/vote.html) to learn how to vote | 50     |
| 2    | After `system-halt` proposal is passed, Tendermint will stop node at certain height.  Then, you can export the snapshot or download the new genesis file provided by IRIS team and start with new version `v0.11.0`. | Participate in software-upgrade                              | follow this [guide](https://www.irisnet.org/docs/features/upgrade.html#interaction-process) to upgrade your node | 200    |
| 3    | When running `v0.11.0`, delegate some **iris** to a validator | Get some **iris** from faucet and delegate to a validator    | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |
| 4    | When running `v0.11.0`, withdraw some **iris** from reward pool | After you become a validator, you could withdraw your rewards by `iriscli distribution withdraw-rewards` | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |

In **phase four**, we expect to deploy a transaction bot which will push system tps to certain level. We expect TPS to increase from 100 to 500 in following days. 

### How to submit evidence

You could submit the transaction Hash to prove you have finished the tasks above by replying this [issue](https://github.com/irisnet/testnets/issues/252):

```
GitHub ID: XXXX
pgp ID: XXX
Phase 1:
Task 1: Link to your PR
Task 2: Node URL
Task 3: Tx Hash
Task 4: Tx Hash

Phase 2:
Task 1: 
Task 2: 
Task 3: 
Task 4: 

Phase 3:
Task 1: 
Task 2: 
Task 3: 
Task 4:

Phase 4:
Task 1: 
Task 2: 
Task 3: 
Task 4: 
```
Only top 40 participants will receive the rewards. 
