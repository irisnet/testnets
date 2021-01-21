# Bifrost-2

## How to Join

It'll be the same as joining [IRIS Hub Mainnet](https://bifrost.irisnet.org/docs/get-started/mainnet.html), except that:
- Go 1.15+ is required for building and installing the IRIShub software.
- [version](https://github.com/irisnet/irishub/releases/tag/v1.0.0-rc0)
- [genesis file](./genesis.json)

If you haven't pariticipate in Phase-1 of Bifrost testnet, please fill in this [form](http://nyancat-irisnet.mikecrm.com/JsLCZ4H) to complete your sign-up process, and remember to set your validator moniker to `<your-name>-<pgp-id>`, identity to `<pgp-id>`. You need to complete the tasks use your validator account.

### Peers

Persistent Peers:

```bash
a246f5ea3055f7507c63b9ea50dbf5d401537b3a@34.80.22.255:26656
```

### Resources

**Explorer**

https://bifrost.irisplorer.io

**Documentation**

https://bifrost.irisnet.org/docs/

**Faucet**

Unfortunately we do not have a public faucet for bifrost, however you are welcome to ask for the test tokens in our validator communities:

- [Riot English community](https://matrix.to/#/!bmimZgJrUWSmxqQEmG:matrix.org?via=matrix.org&via=t2bot.io)
- [QQ Chinese community](https://jq.qq.com/?_wv=1027&k=5BeP3tJ)

## Task list and reward

### Network upgrade task list

| Task                                        | Details                                                              | Completion standard                         | Badge     |
| ------------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------- | --------- |
| Testnet parameter-change voting   | Vote for bifrost-2 parameter-change proposal (activating IBC bidirection transfer)               | Successfully voted for bifrost-2 parameter-change proposal   | Bronze *1 |
| Testnet upgrade voting   | Vote for bifrost-2 upgrade proposal (adding SDK modules)               | Successfully voted for bifrost-2 upgrade proposal   | Bronze *1 |
| Testnet upgrade | Upgrade the validator node according to the upgrade proposal | The upgraded validator node signs at least 1 block in the first 200 blocks after the testnet upgraded  | Silver *1 |


### IBC task list

| Task                                        | Details                                                              | Completion standard                         | Badge     |
| ------------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------- | --------- |
| Send IBC transfer tx to other testnets      | Send IBC transfer tx from bifrost-2 to other testnets                        | Successfully send at least one IBC tx to other testnets  | Gold *1 |
| Receive IBC transfer tx from other testnets      | Receive IBC transfer tx from other testnets to bifrost-2       | Successfully receive at least one IBC tx to the your registered address  | Gold *1 |

**NOTE:**

IBC transfer can be done very easily through the clients, for example, you can send tokens from `Bifrost-2` to other testnets:

```bash
iris tx ibc-transfer transfer transfer <src-channel> <receiver> <amount> \
    --absolute-timeouts \
    --packet-timeout-timestamp 0 \
    --from <sender> \
    --chain-id bifrost-2
```

 For more details, please refer to the client docs:

```bash
iris tx ibc-transfer transfer --help
```

