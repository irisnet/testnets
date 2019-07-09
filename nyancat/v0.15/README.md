# Nyancat Testnet V0.15

## Latest version: v0.15.0-rc0

**Note**:

- The point-to-`iris` conversion rate will be announced at the end of this incentivized testnet campaign.

- For every task, only the top 40 validators who completes can get rewards (the upgrade task is based on the voting list of the upgrade proposal, if a valid voter doesn't complete upgrade in time, the reward will be extended to the next completed validator)

## v0.15.0-rc0

In the `v0.15` version, we introduced a new module: [Asset Management](https://github.com/irisnet/irishub/blob/develop/docs/features/asset.md). IRISHub allows individuals and companies to create and distribute their own tokens for any scenario they can imagine.
Since `v0.15.0-rc0` is not a stable version, we will launch a new testnet `nyancat-3` to test and perform tasks, during which we may upgrade one or more times until the final `v0.15.0` is released. The upgrade of mainnet from `0.14` to `v0.15` is a major event, because this will become the first time in the history of POS network without hard forks to achieve non-compatibility upgrade! So after `v0.15.0` is released, we will propose a new testnet upgrade task, restart `v0.14.3` and upgrade to `v0.15.0` through On-Chain Governance, simulate the entire process of upgrading on the mainnet.

### Tasks

| No   | Name             | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1 | Create a Native Token | Create a Native token using your own validator's `Operator` account and transfer any amount to the public account 1: faa1jpuzrxtu346pae9szdk9thumaxxew740w8w3an | Match the owner of the tokens recieved by the public account | 20 |
2 | Create a Gateway Token | Create a gateway using your own validator's `Operator` account, then refer to Task 1, create a gateway token, and transfer any amount to the public account 2: faa1jmm8fa63ncsphk29e5wj6jy56tea9zz2y202mn | Match the owner of the tokens recieved by the public account | 30 |
| 3 | Mint Tokens | Use your own validator's `Operator` account to mint any of your own issued tokens, and specify the account to accept the minted tokens is the public account 3: faa1cnm7zenktxuhksxx3vj3e78cux98rrh28y4tcd | Match the owner of the tokens recieved by the public account | 10 |
| 4 | Burn Tokens | Use your own validator's `Operator` account to burn any amount of tokens that are issued by yourself | Match the sender address of the valid Burn transaction | 10 |
| 5 | Transfer Ownership of a token | Transfer ownership of any tokens issued by yourself to the public account 4: faa14sk203d2akh96gdlhveglw5emfpf0h8xqls3l0 | Match the original owner of the tokens recieved by the public account | 10 |

### Upgrade Tasks (pre-announcement)

#### Steps [TODO]

| No   | Name             | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 6 | Upgrade to v0.15.0 | Rehearse the mainnet upgrade from v0.14.1 to v0.15.0 | Validators with a valid Pre Commit in the first 300 blocks after the version switch | 200 |
