# How to Finish A Task In Fuxi Incentivized Testnet?

## What is IRISnet incentivized testnet ?

IRIS foundation will offer rewards of IRIS token to early testnet participants according to the points they earned. They could get their rewards after IRISHub is launched. IRISnet development team will publish several task for each iteration of Fuxi testnet. Everyone is welcomed, and you have to fill in this [form](http://cn.mikecrm.com/H9aoXak) to complete your sign-up process. 

You need to use keybase to generate your own [pgp fingerprint](https://github.com/irisnet/testnets/blob/master/fuxi/How%20to%20use%20keybase.md) first. 

The instructions are the following: 

## Tasks for Fuxi-6000

| No   | Name                                            | Details                                                      | Criteria                                                     | Points |
| ---- | ----------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Participate in Genesis file generation process  | Submit your gen-tx.json and use `name-pgp-fingerprint` as validator's name | Submit url of your PR                                        | 100    |
| 2    | Setup a Full node                               | set up a full node and use `name-pgp-fingerprint` as `monkier` | Submit your IP and the team could visit the port of your node | 100    |
| 3    | Delegate some **iris** to a validator           | Get some **iris** from faucet and delegate to a validator    | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |
| 4    | Vote for `software upgrade` proposal                                | Complete an `vote `transaction                               | Run `vote` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction. Refer to this [doc](https://www.irisnet.org/docs/cli-client/gov/vote.html) to learn how to vote  | 50     |
| 5    | Participate in Software upgrade process. You need to Upgrade your sofware to version v0.9.1 once the specific software upgrade proposal is passed in the future. | Participate in software-upgrade process | follow this [guide](https://www.irisnet.org/docs/features/upgrade.html#interaction-process) to upgrade your node  | 200    |


### How to submit evidence

You could submit the transaction Hash to prove you have finished the tasks above by replying this [issue](https://github.com/irisnet/testnets/issues/211):

```
GitHub ID: XXXX
pgp ID: XXX
Task 1: Link to your PR
Task 2: Node URL
Task 3: Tx Hash
Task 4: Tx Hash
Task 5: Validator URL at explorer after the upgrade

```
Only top 40 participants will receive the rewards. 
