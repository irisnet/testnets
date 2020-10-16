# Bifrost Testnet Rules

## Introduction

IRISnet’s DeFi and interchain Testnet Bifrost will run side by side with Stargate’s testnet in order to test the new DeFi features of IRIS Hub, the integration with the new version of Cosmos SDK and also the multi-Hub collaboration through the IBC protocol. [More details](https://medium.com/irisnet-blog/bifr%C3%B6st-irisnets-bridge-to-a-new-era-af32960dd88e)

## Objectives

In order to be compatible with IBC, IRISnet team refactored IRIS Hub based on the latest Cosmos SDK, which is incompatible with previous API and CLI versions. Doing so, we will do our very best to support service providers as wallets, explorers, exchanges to migrate to the new requirements, providing a stable development and a verified environment to them, so Bifrost testnet can continue to run until the mainnet upgrade, providing a stable development and verification environment for service providers.

At the same time, in order to achieve a smooth upgrade of IRIS Hub mainnet, we encourage validators and community members to actively participate in Bifrost testnet, verify the refactored IRIS Hub modules, test the new functions, and explore the potential commercial value and use cases of IRIS Hub's unique features and innovations.

On the other hand, Bifrost will interconnect with Stargate testnet to form a dual-Hub testnet. Teams in the Cosmos interchain ecosystem are welcomed to participate by starting their own testnet as well as service providers interested in providing Relay service in the future. Let’s BUIDL together the New Interchain "Cosmos"!

## Phase-0

As the current version is still rapidly developing, incentivized tasks will not be released in this phase.
This phase is intended for:

- Testing wallets, exchanges and block explorers against the legacy Amino REST interface
- Giving node operators and validators an opportunity to test their integrations against a work in progress version

[[Join phase 0](phase-0/README.md)]

## Phase-1

This phase will mainly test the upgraded modules and new modules and is intended for:

- Testing Service function, including the service creation, binding, request, response, service fees withdrawal, etc.
- Testing Oracle function, querying the response value of the running Oracle
- Testing Record function, including creating and querying records
- Testing NFT function, including issuing, minting, editing, transferring and burning NFT tokens
- Testing Random function, requesting random numbers with or without Oracle

[[Join phase 1](phase-1/README.md)]

## Rewards

**Rewards Levels**:
According to the percentage of achievement of tasks, there will be three reward levels:

- `Gold badge`: Reward coefficient = 1
- `Silver badge`: Reward coefficient = 0.67
- `Bronze badge`: Reward coefficient = 0.33

Each time you complete tasks in a phase, you can get a badge. Namely, when the Bifrost testnet finishes, participants may be eligible for multiple badges. For the specific task list and the conditions for obtaining badges, please refer to the phase details.

**Calculation formula**:

- `Share` = `Rewards Pool / (Number of Golden Badge * 1 + Number of Silver Badge * 0.67 + Number of Bronze Badge * 0.33)`
- `Rewards` = `Share * Reward Coefficient`

**Rewards distribution**:
Rewards will be distributed in the form of NFT badges within the Bifrost testnet. Bifrost testnet will maintain the same account address format as the IRIS Hub mainnet. When IRIS Hub mainnet will upgrade, IRIS foundation will export the list of NFT badges holders from Bifrost testnet at a set time, and import to the Genesis file to launch the mainnet. Therefore:

- The NFT badges can be exchanged into tokens, circulating and transferring in the Bifrost testnet and the upgraded IRIS Hub mainnet.
- When the mainnet completes the upgrade, you can choose to continue to hold NFT badges as an honor, or exchange them for mainnet IRIS. The exchange method will be announced before the mainnet upgrade.
- Please make sure to back up the testnet wallet with NFT badges. Once lost, you will not be able to redeem them.

## Qualification

Any community members of Cosmos and IRISnet can participate in Bifrost testnet, but the following conditions must be met to be eligible for the testnet rewards:

- Participants can't be employees of Bianjie/IRISnet
- The GitHub account has to be at least one year old
- The Keybase account needs to link and connect to the GitHub account

Participants who are found violating any of the rules, such as creating fake accounts for rewards or conducting any malicious behaviors, will be disqualified.

IRISnet reserves all rights for the final interpretation.
