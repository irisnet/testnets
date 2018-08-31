# IRISHub

## Introduction

IRIShub is a blockchain based on Cosmos-SDK. Cosmos-SDK is a blockchain development kit which  provides many useful modules for building a Proof-of-stake blockchain. The MVP modules already implemented in Cosmos-SDK are the following:

* Bank:
* Auth
* Staking
* Slashing
* Governance
* IBC



# Validators


The IRISHub is based on a consensus engine called Tendermint. It relies on a set of validators to secure the network. The role of validators is to run a full-node and participate in consensus by broadcasting votes which contain cryptographic signatures signed by their private key. Validators commit new blocks in the blockchain and receive revenue in exchange for their work. They must also participate in governance by voting on proposals. Validators are weighted according to their total stake.  

The reward for a validator is 3-fold:

* Block provision
* Transaction fee
* Commission

## Potential Risks for A Validator

* DDoS attack


## Running a Validator Node

Validators are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable, double sign a transaction, or don't cast their votes. 

## How to protect against DDoS attack


Please read about [Sentry Node Architecture](https://github.com/irisnet/testnets/blob/master/fuxi/docs/Setup%20A%20Sentry%20Node.md) to protect your node from DDOS attacks and to ensure high-availability.