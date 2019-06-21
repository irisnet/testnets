# Nyancat Testnet V0.14

## Latest version: [v0.14.2](https://github.com/irisnet/irishub/releases/tag/v0.14.2)

## v0.14.2

We found a bug in the Nyancat testnet `v0.14.1` version: for Proposals, we need to deposit a certain amount to enter the voting phase, the unit of `min_deposit` of `Gov` configuration in the `genesis.json` file is not converted in the system, which causes the Nyancat testnet to initiate the proposal with the required amount of 100 `iris`, and the system does not actually have the `iris` token, only `iris-atto`. At the UI layer, the `iris` we saw is actually converted by the formula: 1 iris = 10^18 iris-atto. In other words, we can only deposit `iris-atto`, but the system only recognizes the `iris` that does not exist in the system, so that the amount of the quality deposit can never meet the requirements, and final the proposal fails.

### Bug Details

[Testnet Genesis Parameter Verification](https://github.com/irisnet/irishub/blob/v0.14.1/modules/gov/params.go#L362)

[Nyancat Genesis Configuration](../config/genesis.json#L90)

From the above code we can see that for the `testnet` version of the software, for the convenience of testing, the verification was skipped when loading the Genesis configuration, and the `min_deposit` unit associated with the `Gov` configuration was not converted. , which triggered this bug. However, this problem does not exist in the mainnet, because the software of the mainnet version strictly checks the genesis parameters at startup.

So we made the following optimizations:

- Turn on the testnet version software to verify the genesis parameters
- Genesis only accepts `iris-atto` as `denom` (except account balance)

### Upgrade task

Thanks to the bug, we can do some homework before v0.15 is released!

As we all know that IRIS Hub has a complete on-chain [upgrade](https://www.irisnet.org/docs/features/upgrade.html) and [governance](https://www.irisnet.org/docs/features/governance.html), but for security reasons, the governance can not manage itself, so the modification of the `Gov` parameter can only be done through the `Class 4` upgrade. And because the bug happens to be in the governance module, we can't even let our Nyancat gracefully stop with `SystemHalt Proposal`, so for this upgrade, we can only upgrade in the traditional way.

First, we need to determine a specific block height to derive the final state of the current blockchain. When the block reaches the target height, the validators can follow these steps:

1. Stop the running node and export the blockchain status of the specified height.

    ```bash
    iris export --for-zero-height --height <pending> --output-file nyancat_export.json
    ```

2. Modify the parameters and start time of the new network (to be determined)

    You need to install [jq](https://stedolan.github.io/jq/) to execute the following command, [[online test](https://jqplay.org/s/9QSR4xq_TX)]

    ```bash
    jq -S -c -M '.genesis_time="pending" |.app_state.gov.params = (.app_state.gov.params | .critical_min_deposit[0] = {"denom": "iris-atto", "amount" : "100000000000000000000"}|.important_min_deposit[0] = {"denom": "iris-atto", "amount": "100000000000000000000"}|.normal_min_deposit[0] = {"denom": "iris-atto", " Amount": "50000000000000000000"})' nyancat_export.json > new_genesis.json
    ```

    Verify SHA256 of `new_genesis.json`

    ```bash
    sha256sum new_genesis.json
    xxx (to be determined)
    ```

3. Overwrite the original `genesis.json` with `new_genesis.json`

    ```bash
    cp new_genesis.json <your_home>/config/genesis.json
    ```

4. Install the new version

    ```bash
    git clone https://github.com/irisnet/irishub
    cd irishub && git checkout v0.14.2
    source scripts/setTestEnv.sh
    make get_tools && make install
    ```

    Verify version information

    ```bash
    iris version
    0.14.2-1241d9d-0
    ```

5. Reset the original Nyancat testnet

    ```bash
    iris unsafe-reset-all --home=<your_home>
    ```

6. Start the node

    ```bash
    iris start --home=<your_home>
    ```

| No | Name | Details | Criteria | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
|  2   | Upgrade Node Version to v0.14.2 | Complete the upgrade as above | In the first 50 block heights of the new Nyancat network, the validator nodes with a valid Pre Commit will receive the reward | 200 |

## v0.14.1

As the initial version for Nyancat Testnet, v0.14.1 is already running smoothly on mainnet, so we do not have many tasks todo here, we only need a testnet for new validators and get ready for upcoming tasks in v0.15.

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Setup a Validator Node                         | Set up a full node and use `name-pgp-fingerprint` as `monkier`, then upgrade the full node to a validator | Submit your IP and the team could visit the port of your node | N/A    |
