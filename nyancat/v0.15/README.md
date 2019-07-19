# Nyancat Testnet V0.15

## Latest version: [v0.15.0-rc1](https://github.com/irisnet/irishub/releases/tag/v0.15.0-rc1)

**Note**:

- The point-to-`iris` conversion rate will be announced at the end of this incentivized testnet campaign.

- For every task, only the top 40 validators who completes can get rewards (the upgrade task is based on the voting list of the upgrade proposal, if a valid voter doesn't complete upgrade in time, the reward will be extended to the next completed validator)

## v0.15.0-rc1

In the `v0.15` version, we introduced a new module: [Asset Management](https://github.com/irisnet/irishub/blob/develop/docs/features/asset.md). IRISHub allows individuals and companies to create and distribute their own tokens for any scenario they can imagine.

Since `v0.15.0-rc1` is not a stable version, we will launch a new testnet `nyancat-3` to test and perform tasks, during which we may upgrade one or more times until the final `v0.15.0` is released.

The upgrade of mainnet from `0.14` to `v0.15` is a major event, because this will be the first attempt among BPoS blockchains to perform on-chain software upgrade without stopping the network!  So after `v0.15.0` is released, we will propose a new testnet upgrade task, restart `v0.14.3` and upgrade to `v0.15.0` through on-chain governance, and rehearse the entire process of upgrading the mainnet.

### Keys New Features Tasks

| No   | Name             | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 6 | Multisig account test| [Operation steps](#Multisig-account-test) | Received by public account, signed with multisig account, successful transfer transaction | 10 |
| 7 | Keystore test | [Operation Steps](#Keystore-test) | Determine the task completion by importing the Keystore submitted to GitHub | 10 |

Steps:

#### Multisig account test

[documentation](https://stage.irisnet.org/docs/cli-client/keys/add.html#create-multisig-account)

- Create a multisig account
- Transfer some `iris` to the multisig account.
- Use the multisig account to transfer any amount of `iris` to the public account and memo your PGP-ID: faa1dhwf97xsdjy8xdrxqac5f6zvl2f664dsujmrl5

#### Keystore test

[documentation](https://stage.irisnet.org/docs/cli-client/keys/export.html#example)

- Create a new local account
- Download [Public Test Keystore](task-7/public-keystore.json) and import(password: 1234567890)
- Transfer 10 'iris` to your local account(created by the first step) using the public test account imported in the previous step and memo your PGP-ID(If the balance of public test account is insufficient, please transfer some yourself :)
- Export your local account(created by the first step) to the file `<Your PGP-ID>.json` with password: 1234567890
- Create a Pull Request and submit the `<Your PGP-ID>.json` to GitHub: <https://github.com/irisnet/testnets/tree/master/nyancat/v0.15/task-7/>

### ~~Asset Management Tasks (Completed)~~

| No   | Name             | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1 | Create a Native Token | Create a Native token using your own validator's `Operator` account and transfer any amount to the public account 1: faa1jpuzrxtu346pae9szdk9thumaxxew740w8w3an | Match the owner of the tokens recieved by the public account | 20 |
2 | Create a Gateway Token | Create a gateway using your own validator's `Operator` account, then refer to Task 1, create a gateway token, and transfer any amount to the public account 2: faa1jmm8fa63ncsphk29e5wj6jy56tea9zz2y202mn | Match the owner of the tokens recieved by the public account | 30 |
| 3 | Mint Tokens | Use your own validator's `Operator` account to mint any of your own issued tokens, and specify the account to accept the minted tokens is the public account 3: faa1cnm7zenktxuhksxx3vj3e78cux98rrh28y4tcd | Match the owner of the tokens recieved by the public account | 10 |
| 4 | Burn Tokens | Use your own validator's `Operator` account to burn any amount of tokens that are issued by yourself | Match the sender address of the valid Burn transaction | 10 |
| 5 | Transfer Ownership of a token | Transfer ownership of any tokens issued by yourself to the public account 4: faa14sk203d2akh96gdlhveglw5emfpf0h8xqls3l0 | Match the original owner of the tokens recieved by the public account | 10 |
| 6 | Transfer Ownership of a token | Transfer ownership of any tokens issued by yourself to the public account 4: faa14sk203d2akh96gdlhveglw5emfpf0h8xqls3l0 | Match the original owner of the tokens recieved by the public account | 10 |

### Upgrade Tasks (pre-announcement)

#### Steps [TODO]

| No   | Name             | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| * | Upgrade to v0.15.0 | Rehearse the mainnet upgrade from v0.14.* to v0.15.0 | Validators with a valid Pre Commit in the first 300 blocks after the version switch | 200 |
