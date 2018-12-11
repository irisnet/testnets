# How to Finish A Task In Fuxi Incentivized Testnet?

## What is IRISnet incentivized testnet ?

IRIS foundation will offer rewards of IRIS token to early testnet participants according to the points they earned. They could get their rewards after IRISHub is launched. IRISnet development team will publish several task for each iteration of Fuxi testnet. Everyone is welcomed, and you have to fill in this [form](http://cn.mikecrm.com/H9aoXak) to complete your sign-up process. 

You need to use keybase to generate your own [pgp fingerprint](https://github.com/irisnet/testnets/blob/master/fuxi/How%20to%20use%20keybase.md) first. 

The instructions are the following: 

## Tasks for Fuxi-4000

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Participate in Genesis file generation process | Submit your gen-tx.json and use `name-pgp-fingerprint` as validator's name | Submit url of your PR                                        | 100    |
| 2    | Setup a Full node                              | set up a full node and use `name-pgp-fingerprint` as `monkier` | Submit your IP and the team could visit the port of your node | 100    |
| 3    | Delegate some **iris** to a validator          | Get some **iris** from faucet and delegate to a validator    | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |
| 4    | Redelegate some **iris** from a validator      | Complete an redelegate transaction                           | Run `redelegate begin` and `redelegat complete` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50     |
| 5    | Participate in Software upgrade process (More details will be announced in Riot)       | Participate in software-upgraden vote&Upgrade your sofware to version v0.6.2 | follow this guide to upgrade your node                       | 200    |
| 6    | Vote a proposal                                | Complete an `vote `transaction                               | Run `vote` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 20     |
| 7    | Deposit to a governance proposal               | Complete an `deposit `transaction                            | Run `deposit` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 20     |


### How to submit evidence

You could submit the transaction Hash to prove you have finished the tasks above by replying this [issue](https://github.com/irisnet/testnets/issues/129):

```
GitHub ID: XXXX
pgp ID: XXX
Task 1: Link to your PR
Task 2: Node URL
Task 3: Tx Hash
Task 4: Tx Hash
Task 5: Hash of your submit switch transaction
Task 6: Tx Hash
Task 7: Tx Hash
...

```
Only top 40 participants will receive the rewards. 
