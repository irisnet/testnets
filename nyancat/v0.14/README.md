# Nyancat Testnet V0.14

## Latest version: [v0.14.3](https://github.com/irisnet/irishub/releases/tag/v0.14.3)

## v0.14.3

This new version is to address a weakness in `v0.14.1` which, when coinciding with governance parameter setting errors, leads to governance procedures unusable in testnets.
We found this problem when trying to go through a governance procedure in Nyancat.

### Bug Details in `v0.14.1`

[Testnet Genesis Parameter Verification](https://github.com/irisnet/irishub/blob/v0.14.1/modules/gov/params.go#L362)

[Nyancat Genesis Configuration](../config/genesis.json#L90)

From the above code we can see that for the testnet version of the software, parameter validation was skipped when loading the genesis configuration, and the `min_deposit` denom associated with the `gov` configuration was not converted, which triggered this bug. However, this problem does not exist in the mainnet, because the mainnet version strictly validates all genesis parameters at startup.

To fix this problem, we made the following improvement in `v0.14.3`:

- Turn on the validation of genesis parameters for testnets

### Upgrade Task

Thanks to this bug, we can do some homework before v0.15 is released!

As we all know, IRIS Hub comes with a comprehensive on-chain [software upgrade](https://www.irisnet.org/docs/features/upgrade.html) and [governance](https://www.irisnet.org/docs/features/governance.html) mechanism; however, for security reasons, we have purposefully chosen not to allow the governance itself to be governed.  So, to update the incorrectly set governance parameters, 
we have to go through a hard fork.

Specifically for the Nyancat testnet, we will carry out the fork in the following steps:

1. We'll bump the voting power of `Bianjie` to over 2/3

2. We'll take `Bianjie` offline to stop the Nyancat network, announcement will be made before we do this

3. Export the blockchain state at the specified height (TBD).

    ```bash
    iris export --for-zero-height --height <TBD> --output-file nyancat_export.json
    ```

4. Modify the governance parameters and `genesis_time` (TBD)

    You need to install [jq](https://stedolan.github.io/jq/) to execute the following command, [[online test](https://jqplay.org/s/9QSR4xq_TX)]

    ```bash
    jq -S -c -M '.genesis_time="TBD" |.app_state.gov.params = (.app_state.gov.params | .critical_min_deposit[0] = {"denom": "iris-atto", "amount" : "100000000000000000000"}|.important_min_deposit[0] = {"denom": "iris-atto", "amount": "100000000000000000000"}|.normal_min_deposit[0] = {"denom": "iris-atto", " Amount": "50000000000000000000"})' nyancat_export.json > new_genesis.json
    ```

    Verify SHA256 of `new_genesis.json`

    ```bash
    sha256sum new_genesis.json
    xxx (TBD)
    ```

5. Overwrite the original `genesis.json` with `new_genesis.json`

    ```bash
    cp new_genesis.json <your_home>/config/genesis.json
    ```

6. Build and install `v0.14.3`

    ```bash
    git clone https://github.com/irisnet/irishub
    cd irishub && git checkout v0.14.3
    source scripts/setTestEnv.sh
    make get_tools && make install
    ```

    Verify version information

    ```bash
    iris version
    0.14.3-9084de2e-0
    ```

7. Reset the original Nyancat testnet

    ```bash
    iris unsafe-reset-all --home=<your_home>
    ```

8. Start the node

    ```bash
    iris start --home=<your_home>
    ```

| No | Name | Details | Criteria | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
|  2   | Upgrade to v0.14.3 | (see above) | In the first 100 blocks of the new Nyancat network, validators with a valid precommit will receive the reward | 200 |

**Note**: The point-to-`iris` conversion rate will be announced at the end of this incentivized testnet campaign.

## v0.14.1

As the initial version for Nyancat Testnet, v0.14.1 is already running smoothly on mainnet, so we do not have many tasks todo here, we only need a testnet for new validators and get ready for upcoming tasks in v0.15.

| No   | Name                                           | Details                                                      | Criteria                                                     | Points |
| ---- | ---------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------ |
| 1    | Setup a Validator Node                         | Set up a full node and use `name-pgp-fingerprint` as `monkier`, then upgrade the full node to a validator | Submit your IP and the team could visit the port of your node | N/A    |
