# Testnet Status
## Feb. 4 Fuxi testnet comes online
The goal for Fuxi testnet is to provide a stable and free enviroment for testing applications, services and client tools development.  So valdiators no need to join the Fuxi testnet

## Feb. 11  New Release
Irishub v0.12.0 is released 

## Jan. 29 Load test

THe network tps was about 200 at peak.

## Jan. 25 Restrt Upgrade was tested 
A system-halt proposal was passed and the chain stopped. All the states snapshot were exported as genesis file and the chain restart with new version v0.11.0.

## Jan. 24  New Release
Irishub v0.11.0 is released 

## Jan. 21,Patch Upgrade was tested 
A consensus failure was triggered in v0.10.2 and the nodes need to switch to v0.10.2-patch to recover from consensus failure

## Jan. 20 2019 Sucessfulw Smooth Software Upgarade 
**iris** version for tsetnet switch from v0.10.0 to v0.10.2

## Jan. 19 2019 Software Upgarade Proposal is passed
A smooth software upgrade proposal is submitted by dev team and passed: https://testnet.irisplorer.io/#/ProposalsDetail/1

## Jan 17 2019 - Fuxi-8000 Go Live

We bootstraped the latest testnet with 21 genesis validators. Now, fuxi-8000 has 21 validators and the average block time is about 5s.
## Jan. 16  New Release
Irishub v0.10.0 is released 

## Jan. 9 Fuxi-7000 Testnet Halted
Fuxi-7000 on-chain upgrade process is halted due to Tendermint bug.The team is doing internal QA process and is about to release a new version: v0.10.0. 

## Jan. 5  New Release
Irishub v0.9.1 is released and this leads to consensus failure, team quickly released a patch: v0.9.1-patch0

## Jan. 4, 2019 Software Upgarade Proposal is passed
A software upgrade proposal is submitted by dev team and passed.

## Dec. 29, 2018 - Fuxi-7000 Go Live

We bootstraped the latest testnet with 12 genesis validators. Now, fuxi-7000 has 25 validators and the average block time is about 5s.

## Nov. 27, 2018 - New Release

Released v0.9.0
https://github.com/irisnet/irishub/releases/tag/v0.9.0
## Dec 17, 2018 - Fuxi-6000 Go Live

We bootstraped the latest testnet with 10 genesis validators. Now, fuxi-6000 has 20 validators and the average block time is about 5s.

## Dec. 7 2018 - New Release

Released v0.8.0

https://github.com/irisnet/irishub/releases/tag/v0.8.0

## Nov 30, 2018 - Fuxi-5000 Go Live

We bootstraped the latest testnet in a centralized manner. Now, fuxi-5000 has 25 validators and the average block time is about 6s.
## Nov 27, 2018 - New Release

Released v0.7.0
https://github.com/irisnet/irishub/releases/tag/v0.7.0


## Nov 12, 2018 - Fuxi-4000 Go Live

We bootstraped the latest testnet in a centralized manner. Now, fuxi-4000 has 28 validators and the average block time is about 5s.

## Nov 1, 2018 - New Release

Released v0.6.0

BREAKING CHANGES:

[monitor] Use new executable binary in monitor
FEATURES:

[record] Add the record module of the data certification on blockchain
[iservice] Add the feature of iService definition
[iservice] Use --def-chain-id flag to reference the blockchain defined of the iService
[cli] Add the example description in the cli help
[test] Add Cli/LCD/Sim auto-test
[test] Add cli and lcd test for record module
[docs] Update the user doc of iservice definition and record

BUG FIXES:
Fix software upgrade issue caused by tx fee
Report Panic when building the lcd proof
Fix bugs in converting validator power to byte array
Fix panic bug in wrong account number
Fix some bugs about iservice definition and record

## Oct 18, 2018 - New Release

Released v0.5.0


## Sep 26, 2018 - Fuxi-3001 Go Live

We started the third iteration of public testnet: fuxi-3001.

This testnet now has 28 validators and the average block time is about 5s. 

## Sep 22, 2018 - New Release

Released v0.4.2

BUG FIX: Fix consensus failure due to the double sign evidence be broadcasted before the genesis block

## Sep 6, 2018 - New Release

Released v0.4.0 with update for Tendermint v0.22.6 and cosmos-sdk v0.23.0

## Aug 10, 2018 - Fuxi-2000 Go live

We started the second iteration of public testnet: fuxi-2000.

This testnet now has 20 validators and the average block time is about 5s. 

## Aug 8, 2018 - New Release

Released v0.3.3 with update for Tendermint v0.22.2 and Cosmos-sdk v0.23.0

## Jul 19, 2018 - New Release

Released v0.2.0 with update for Tendermint v0.21.0 and Cosmos-sdk v0.19.1
Finished codebase refactor 
Please update to this to sync with the testnet

## Jul 25, 2018 - Fuxi-1002 Go live

We started the first public testnet: fuxi-1002.

This testnet has 11 validators and the average block time is about 5s. 
