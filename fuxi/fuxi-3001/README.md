# How to Finish A Task In Fuxi Incentivized Testnet?

## What is IRISnet incentivized testnet ?

IRIS foundation will offer rewards of IRIS token to early testnet participants. They could get their rewards after IRISHub is launched. IRISnet development team will publish several task for each iteration of Fuxi testnet. Everyone is welcomed, and you have to fill in this [form](http://cn.mikecrm.com/H9aoXak) to complete your sign-up process. 

You need to use keybase to generate your own [pgp fingerprint](https://github.com/irisnet/testnets/blob/master/fuxi/How%20to%20use%20keybase.md) first. 

The instructions are the following: 

## Tasks for Fuxi-3000

 | No   | Name                                           | Details                                                      | Criteria                                                     | Reward  |
  | ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------- |
  | 1    | Participate in Genesis file generation process | Submit your gen-tx.json and use `name-pgp-fingerprint` as validator's name | Submit your IP and thr team could check the configutration of your node | 50 iris |
  | 2   | Setup a Full node                              | set up a full node and use `name-pgp-fingerprint` as `monkier` | Submit your IP and thr team could check the configutration of your node | 50 iris |
  | 3   | Delegate some **iris** to a validator          | Get some **iris** from faucet and delegate to a validator    | Run this send transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50 iris |
  | 4   | Redelegate some **iris** from a validator      | Complete an redelegate transaction                           | Run `redelegate begin` and` redelegat complete` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 50 iris |
  | 5   | Submit a governance proposal                   | Complete an `submit-proposal`transaction                     | Run `submit-proposal` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 20 iris |
  | 6    | Vote a proposal                                | Complete an `vote`transaction                                | Run `vote` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 20 iris |
  | 7    | Deposit to a governance proposal               | Complete an `deposit`transaction                             | Run `deposit` transaction with your `name-pgpid` as memo. Submit your transaction hash and the team could verify the details of transaction | 20 iris |



### How to submit evidence

You could submit the transaction Hash to prove you have finished the tasks above by replying this [issue](https://github.com/irisnet/testnets/issues/83):

```
GitHub ID: XXXX
pgp ID: XXX
Task 1: Link to your PR
Task 2: Node URL
Task 3: Hash
Task 4: Hash
...

```
Only the Top 20 participants will receive the rewards. 
